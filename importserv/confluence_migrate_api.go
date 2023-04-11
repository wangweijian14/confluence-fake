package importserv

import (
	"encoding/json"
	"fmt"
	"sort"
)

type MigrationConfig struct {
	Domain       string `json:"domain"`
	UserName     string `json:"user_name"`
	PassWord     string `json:"pass_word"`
	OnesDomain   string `json:"ones_domain"`
	OperatorUUID string `json:"operator_uuid"`
	Name         string `json:"name"`
	Status       int    `json:"status"`    //1 确认任务保存 0 历史存储任务，可能修改
	TaskUUID     string `json:"task_uuid"` //1 TaskUUId
}

func CheckCFConfig(req *MigrationConfig) error {
	connector := NewConfluenceConnector(req.UserName, req.PassWord, req.Domain, 1)
	_, err := connector.checkConfluenceTong()
	if err != nil {
		return err
	}
	return nil
}

type ConfluenceSpaceListResp struct {
	SpaceName string `json:"space_name"`
	SpaceKey  string `json:"space_key"`
	SpaceId   int64  `json:"space_id"`
	Type      string `json:"type"`
	Status    int32  `json:"status"` //0 未迁移 1迁移过
}

func (c *ConfluenceSpaceListResp) ToJson() (string, error) {
	json_string, err := json.Marshal(c)
	if err != nil {
		return "", err
	} else {
		return string(json_string), nil
	}
}

func GetConfluenceSpaceList(config *MigrationConfig, keyList []string) ([]*ConfluenceSpaceListResp, error) {
	resp := make([]*ConfluenceSpaceListResp, 0)
	var err error
	keyMap := make(map[string]int)
	spaceList := make([]*SpaceKeyInfo, 0)
	confluenceClient := NewConfluenceConnector(config.UserName, config.PassWord, config.Domain, 1)
	if len(keyList) == 0 {
		//max = 0 去全部数据 ，limitSize 每次获取多少条数据，一次最多200
		spaceList, err = confluenceClient.getSpaceKeyList(0, 100)
		for _, v := range spaceList {
			keyList = append(keyList, v.Key)
			keyMap[v.Key] = NotMigrationSpaceStatus
		}
	} else {
		//取指定数据
		spaceList, err = confluenceClient.getSpaceKeyListByKeys(keyList)
		for _, v := range spaceList {
			keyMap[v.Key] = NotMigrationSpaceStatus
		}
	}

	if err != nil {
		fmt.Printf("GetCFSpaceByKeyList fail %v", err)
	}

	for _, v := range spaceList {
		temp := &ConfluenceSpaceListResp{
			SpaceName: v.Name,
			SpaceKey:  v.Key,
			Type:      v.Type,
			SpaceId:   v.Id,
			Status:    int32(keyMap[v.Key]),
		}
		resp = append(resp, temp)
	}

	sort.Slice(resp, func(i, j int) bool {
		if resp[i].Type < resp[j].Type {
			return true
		}
		return resp[i].Type == resp[j].Type && resp[i].SpaceKey < resp[j].SpaceKey
	})
	return resp, nil
}

// 获取所有用户
// confluence 没有专门的接口获取所有用户，目前有两种方案：
// 1. 获取 confluence-users 用户组的用户列表，该用户组为默认所有人都会加入的用户组，不在用户组中的用户没有使用 confluence 的权限
// 2. 所有所有用户组的用户信息，然后进行合并
// 下面的代码采取第二种方案
// func getAllUsers(connector *ConfluenceConnector) (map[string]*User, map[string][]string, error) {
// 	users := make(map[string]*User)
// 	groupMembers := make(map[string][]string)
// 	groups, err := connector.getGroups()
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	userMap := make(map[string]*User)
// 	// 获取所有用户组的成员列表，并进行合并
// 	for _, group := range groups {
// 		groupMember, err := connector.getGroupMembers(group.Name)
// 		if err != nil {
// 			return nil, nil, err
// 		}
// 		groupMemberIds := make([]string, 0, len(groupMember))
// 		for _, user := range groupMember {
// 			groupMemberIds = append(groupMemberIds, user.Id)
// 			if _, ok := userMap[user.Id]; !ok {
// 				userMap[user.Id] = user
// 			}
// 		}
// 		groupMembers[group.Name] = groupMemberIds
// 	}
// 	// 并发获取用户的 email
// 	helper := NewConcurrentHelper(connector.concurrentLimit)
// 	userCache := utility.NewSafeMap(len(userMap))
// 	for _, user := range userMap {
// 		if userCache.Load(user.Id) != nil {
// 			continue
// 		}
// 		u := user
// 		helper.Add(func() (interface{}, error) {
// 			email, err1 := connector.getEmailByName(u.Name)
// 			if email != "" {
// 				u.Email = email
// 				u.LowerEmail = strings.ToLower(u.Email)
// 				userCache.Store(u.Id, u)
// 			} else {
// 				userCache.Store(u.Id, u)
// 			}
// 			return email, err1
// 		})
// 	}
// 	helper.Run()
// 	userCache.ForEach(func(key string, value interface{}) {
// 		users[key] = value.(*User)
// 	})
// 	return users, groupMembers, nil
// }
