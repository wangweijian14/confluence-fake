package importserv

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/cenkalti/backoff"
)

type ConfluenceConnector struct {
	sync.RWMutex
	username        string
	password        string
	domain          string
	concurrentLimit int
}

var (
	confluenceAPIBackOff = &backoff.ExponentialBackOff{
		InitialInterval:     200 * time.Millisecond,
		RandomizationFactor: backoff.DefaultRandomizationFactor,
		Multiplier:          backoff.DefaultMultiplier,
		MaxInterval:         5 * time.Second,
		MaxElapsedTime:      20 * time.Second,
		Clock:               backoff.SystemClock,
	}
)

func NewConfluenceConnector(username, password, domain string, connectionLimit int) *ConfluenceConnector {
	if connectionLimit == 0 {
		connectionLimit = 10
	}

	if username == "" || password == "" || domain == "" {
		return nil
	}

	return &ConfluenceConnector{
		username:        username,
		password:        password,
		domain:          domain,
		concurrentLimit: connectionLimit,
	}
}

const (
	cf               = "Confluence"
	migrating        = "Migrating"
	timeOut          = "TimeOut"
	domainError      = "DomainError"
	authError        = "AuthError"
	exportExcelError = "ExportExcelError"
)

func (c *ConfluenceConnector) do(req *http.Request, result interface{}) error {
	client := &http.Client{}
	req.SetBasicAuth(c.username, c.password)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusUnauthorized {

	}
	if resp.StatusCode != http.StatusOK {
		reqUrl := fmt.Sprintf("%s?%s", req.URL.Path, req.URL.RawQuery)
		msg := fmt.Sprintf("confluence request fail, url: %s, status: %d, message: %s", reqUrl, resp.StatusCode, string(body))
		return fmt.Errorf(msg)
	}
	if result != nil {
		err = json.Unmarshal(body, result)
	}
	return err
}

func (c *ConfluenceConnector) doGetWithTimeOut(url string, result interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("doGetWithTimeOut new request  err %v", err)

	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	req.SetBasicAuth(c.username, c.password)
	req.Header.Add("Content-Type", "application/json")
	fmt.Printf("client sent : %v\n", req)
	resp, err := client.Do(req)
	if err != nil {

	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("doGetWithTimeOut err %v", err)
	}
	if resp.StatusCode == http.StatusUnauthorized {

	}
	if resp.StatusCode != http.StatusOK {
		reqUrl := fmt.Sprintf("%s?%s", req.URL.Path, req.URL.RawQuery)
		msg := fmt.Sprintf("confluence request fail, url: %s, status: %d, message: %s", reqUrl, resp.StatusCode, string(body))
		return fmt.Errorf(msg)

	}
	if result != nil {
		err = json.Unmarshal(body, result)
		if err != nil {
			return fmt.Errorf("doGetWithTimeOut unmarshal err %v", err)

		}
	}
	return nil
}

func (c *ConfluenceConnector) doGet(url string, resp interface{}) error {
	httpReq, err := http.NewRequest("GET", url, nil)
	fmt.Println(httpReq)
	if err != nil {
		return err
	}
	return c.do(httpReq, resp)
}
func (c *ConfluenceConnector) doPost(url string, resp interface{}, body []byte) error {

	httpReq, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return err
	}
	return c.do(httpReq, resp)
}

func (c *ConfluenceConnector) doGet2(url string, resp *[]byte) error {
	httpReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	return c.do2(httpReq, resp)
}

//直接返回resp的body，不用 unmarshal
func (c *ConfluenceConnector) do2(req *http.Request, result *[]byte) error {
	cacheKey := fmt.Sprintf("%s?%s", req.URL.Path, req.URL.RawQuery)
	client := &http.Client{}
	req.SetBasicAuth(c.username, c.password)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("do2  request fail, url: %s, status: %d", cacheKey, resp.StatusCode)
		fmt.Print(msg)
		return fmt.Errorf(msg)
	}
	*result = append(*result, body...)
	return err
}

type GroupInfoResp struct {
	Results []*UserInfoForCF `json:"results"`
	Start   int64            `json:"start"`
	Size    int64            `json:"size"`
	Limit   int64            `json:"limit"`
	Links   interface{}      `json:"_links"`
}

type UserInfoForCF struct {
	Type        string `json:"type"`
	Username    string `json:"username"`
	UserKey     string `json:"userKey"`
	UserEmail   string `json:"user_email"`
	DisplayName string `json:"displayName"`
	Status      string `json:"status"`
}

func (c *ConfluenceConnector) getGroupMembers(cfGroupName string) ([]*User, error) {
	users := make([]*User, 0)
	start := 0
	limit := 200
	apiTemplate := fmt.Sprintf("%s/rest/api/group/%s/member?expand=status&limit=%d&start=", c.domain, url.QueryEscape(cfGroupName), limit)
	//
	for {
		apiUrl := apiTemplate + fmt.Sprintf("%d", start)
		spaceResp := &GroupInfoResp{}

		err := backoff.Retry(func() error {
			if err := c.doGet(apiUrl, spaceResp); err != nil {
				return err
			}
			return nil
		}, confluenceAPIBackOff)
		if err != nil {
			return nil, fmt.Errorf("get confluence group member error: %v, groupName: %s, start: %d", err, cfGroupName, start)
		}
		for _, v1 := range spaceResp.Results {
			users = append(users, &User{
				Id:          v1.UserKey,
				Name:        v1.Username,
				DisplayName: v1.DisplayName,
				Status:      v1.Status,
			})
		}
		//
		if spaceResp.Size < int64(limit) {
			break
		}
		//
		start += limit
	}
	return users, nil
}

type UserEmailInfoForCFResp struct {
	Username  string      `json:"userName"`
	UserEmail string      `json:"email"`
	Start     int64       `json:"start"`
	Size      int64       `json:"size"`
	Limit     int64       `json:"limit"`
	Links     interface{} `json:"_links"`
}

func (c *ConfluenceConnector) getEmailByName(name string) (string, error) {
	apiUrl := fmt.Sprintf("%s/rest/mobile/1.0/profile/%s", c.domain, url.QueryEscape(name))
	userEmail := &UserEmailInfoForCFResp{}
	if err := c.doGet(apiUrl, userEmail); err != nil {
		return "", err
	}
	return userEmail.UserEmail, nil
}

//cf 导出逻辑
//1 获取空间所有列表，获取key数据列表，提供给选择
//2 给定列表key文件

//confluence rest api
const (
	spaceKeyListURI = "/rest/api/space?limit=%d&start="
	spaceKeyURI     = "/rest/api/space"
	exportUrlURI    = "/rpc/json-rpc/confluenceservice-v2/exportSpace"
	pagePDFURI      = "/spaces/flyingpdf/pdfpageexport.action?pageId="
)

type SpaceKeyInfo struct {
	Id   int64  `json:"id"`
	Key  string `json:"key"`
	Type string `json:"type"`
	Name string `json:"name"`
}
type SpaceResp struct {
	Results []SpaceInfo `json:"results"`
	Start   int64       `json:"start"`
	Size    int64       `json:"size"`
	Limit   int64       `json:"limit"`
	Links   interface{} `json:"_links"`
}
type SpaceInfo struct {
	Id         int64                  `json:"id"`
	Key        string                 `json:"key"`
	Type       string                 `json:"type"`
	Name       string                 `json:"name"`
	Links      map[string]interface{} `json:"_links"`
	Expandable interface{}            `json:"_expandable"`
}

func (c *ConfluenceConnector) checkConfluenceTong() ([]*SpaceKeyInfo, error) {

	apiUrl := fmt.Sprintf("%s"+spaceKeyListURI, c.domain, 2)
	index := 0

	spaceResp := SpaceResp{}
	spaceList := make([]*SpaceKeyInfo, 0)

	if err := c.doGetWithTimeOut(fmt.Sprintf("%s%d", apiUrl, index), &spaceResp); err != nil {
		return spaceList, err
	}
	for _, v := range spaceResp.Results {
		spaceList = append(spaceList, &SpaceKeyInfo{
			Id:   v.Id,
			Key:  v.Key,
			Type: v.Type,
			Name: v.Name,
		})
	}
	return spaceList, nil
}

//获取空间列表信息
func (c *ConfluenceConnector) getSpaceKeyList(max, limitSize int) ([]*SpaceKeyInfo, error) {
	if limitSize > 200 {
		limitSize = 200
	}
	apiUrl := fmt.Sprintf("%s"+spaceKeyListURI, c.domain, limitSize)
	index := 0

	spaceResp := SpaceResp{}
	spaceList := make([]*SpaceKeyInfo, 0)

	for {

		//限制获取列表的个数,max==0时，去全部数据
		if max != 0 && index >= max {
			break
		}

		if err := c.doGet(fmt.Sprintf("%s%d", apiUrl, index), &spaceResp); err != nil {
			return spaceList, err
		}
		for _, v := range spaceResp.Results {
			spaceList = append(spaceList, &SpaceKeyInfo{
				Id:   v.Id,
				Key:  v.Key,
				Type: v.Type,
				Name: v.Name,
			})
		}
		index += limitSize
		if spaceResp.Size < spaceResp.Limit {
			break
		}

	}

	return spaceList, nil
}

//http://127.0.0.1:8084/rest/api/space?spaceKey=CF743
func (c *ConfluenceConnector) getSpaceKeyListByKeys(keyList []string) ([]*SpaceKeyInfo, error) {
	apiUrl := fmt.Sprintf("%s"+spaceKeyURI, c.domain)
	index := 0

	spaceResp := SpaceResp{}
	spaceList := make([]*SpaceKeyInfo, 0)

	apiUrl += "?spaceKey=" + strings.Join(keyList, "&spaceKey=")

	if err := c.doGet(fmt.Sprintf("%s%d", apiUrl, index), &spaceResp); err != nil {
		return spaceList, err
	}
	for _, v := range spaceResp.Results {
		spaceList = append(spaceList, &SpaceKeyInfo{
			Id:   v.Id,
			Key:  v.Key,
			Type: v.Type,
			Name: v.Name,
		})
	}

	return spaceList, nil
}

//获取cf导出zip的url
func (c *ConfluenceConnector) getSpaceUrl(key string) (string, error) {
	result := ""

	apiUrl := fmt.Sprintf("%s"+exportUrlURI, c.domain)
	body := []interface{}{
		key,
		"TYPE_XML",
		true,
	}
	bytesBody, _ := json.Marshal(body)

	err := c.doPost(apiUrl, &result, bytesBody)
	return result, err
}

//下载cf的xml压缩包
func (c *ConfluenceConnector) exportSpaceXmlZip(url string) ([]byte, error) {
	result := make([]byte, 0)
	err := c.doGet2(url, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ConfluenceConnector) writerFile(path string, data []byte) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, bytes.NewReader(data))
	return err
}

//获取页面pdf
func (c *ConfluenceConnector) GetPagePDF(pageId string) ([]byte, error) {
	apiUrl := fmt.Sprint(c.domain + pagePDFURI + pageId)
	pagePDF := make([]byte, 0)
	fmt.Println("apiUrl:  ", apiUrl)
	if err := c.doGet(apiUrl, &pagePDF); err != nil {
		return pagePDF, err
	}
	return pagePDF, nil
}

// type GroupListResp struct {
// 	Groups []*Group `json:"results"`
// 	Start  int64    `json:"start"`
// 	Size   int      `json:"size"`
// 	Limit  int      `json:"limit"`
// }

// type Group struct {
// 	Type string `json:"type"`
// 	Name string `json:"name"`
// }

// func (c *ConfluenceConnector) getGroups() ([]*Group, error) {
// 	groups := make([]*Group, 0)
// 	start := 0
// 	limit := 200
// 	apiTemplate := fmt.Sprintf("%s/rest/api/group/?limit=%d&start=", c.domain, limit)
// 	// 没有总数，需要按照分页数据一直获取，直到取不到任何数据
// 	for {
// 		apiUrl := apiTemplate + fmt.Sprintf("%d", start)
// 		groupsResp := &GroupListResp{}
// 		// 重试
// 		err := backoff.Retry(func() error {
// 			if err := c.doGet(apiUrl, groupsResp); err != nil {
// 				return err
// 			}
// 			return nil
// 		}, confluenceAPIBackOff)
// 		if err != nil {
// 			return nil, fmt.Errorf("get group error: %v, start: %d", err, start)
// 		}
// 		groups = append(groups, groupsResp.Groups...)
// 		if groupsResp.Size == 0 {
// 			break
// 		}
// 		start += groupsResp.Size
// 	}
// 	return groups, nil
// }
