package importserv

const (
	NotMigrationSpaceStatus = 0
	MigratingSpaceStatus    = 1
	MigratedSpaceStatus     = 2
	MigrateFailSpaceStatus  = 3
)

type FakeConfluenceSpaceResult struct {
	Results []*ConfluenceSpaceData `json:"results"`
	Start   int                    `json:"start"`
	Limit   int                    `json:"limit"`
	Size    int                    `json:"size"`
	Links   LinksBase              `json:"_links"`
}

type Links struct {
	Webui string `json:"webui"`
	Self  string `json:"self"`
}
type Expandable struct {
	Metadata    string `json:"metadata"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
}
type ConfluenceSpaceData struct {
	ID         int        `json:"id"`
	Key        string     `json:"key"`
	Name       string     `json:"name"`
	Type       string     `json:"type"`
	Links      Links      `json:"_links"`
	Expandable Expandable `json:"_expandable"`
}
type LinksBase struct {
	Self    string `json:"self"`
	Base    string `json:"base"`
	Context string `json:"context"`
}

type FakeGroupResult struct {
	Results []*UserGroup `json:"results"`
	Start   int          `json:"start"`
	Limit   int          `json:"limit"`
	Size    int          `json:"size"`
	Links   LinksBase    `json:"_links"`
}

type FakeGroupUserListResp struct {
	Results []*UserConfluence `json:"results"`
	Start   int               `json:"start"`
	Size    int               `json:"size"`
	Limit   int               `json:"limit"`
	Links   LinksBase         `json:"_links"`
}

type UserGroup struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Links Links  `json:"_links"`
}

type UserConfluence struct {
	Type            string          `json:"type"`
	Status          string          `json:"status"`
	Username        string          `json:"username"`
	UserKey         string          `json:"userKey"`
	ProfilePicture  ProfilePicture  `json:"profilePicture"`
	DisplayName     string          `json:"displayName"`
	FullName        string          `json:"fullName"`
	Email           string          `json:"email"`
	UnknownUser     bool            `json:"unknownUser"`
	Anonymous       bool            `json:"anonymous"`
	Links           Links           `json:"_links"`
	UserPreferences UserPreferences `json:"userPreferences"`
	GpName          string          `json:"gpName"`
}

type ProfilePicture struct {
	Path      string `json:"path"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	IsDefault bool   `json:"isDefault"`
}

type UserPreferences struct {
	WatchOwnContent bool `json:"watchOwnContent"`
}

type GroupListResp struct {
	Groups []*UserGroup `json:"results"`
	Start  int64        `json:"start"`
	Size   int          `json:"size"`
	Limit  int          `json:"limit"`
}

// import (
// 	"archive/zip"
// 	"fmt"
// 	"strconv"
// 	"strings"

// 	"github.com/bangwork/ones-ai-api-common/utils/i18n"
// 	utilsModel "github.com/bangwork/wiki-api/app/models/utils"
// 	"github.com/bangwork/wiki-api/app/utils"
// 	"github.com/bangwork/wiki-api/app/utils/errors"
// )

// type ConfluenceDataModel interface {
// 	GetType() ConfluenceEntityType
// 	GetId() string
// }

// type WizEditorObject struct {
// 	Data map[string]interface{} `json:"-"`
// }

// func NewWizEditorObjectWithData(data map[string]interface{}) *WizEditorObject {
// 	editorObject := &WizEditorObject{Data: data}
// 	editorObject.Data["meta"] = make(map[string]interface{})
// 	editorObject.Data["comments"] = make(map[string]interface{})
// 	return editorObject
// }

// func NewWizEditorObject() *WizEditorObject {
// 	return NewWizEditorObjectWithData(make(map[string]interface{}))
// }

// func NewWizEditorObjectWithTitle(title string) *WizEditorObject {
// 	var blocks []WizDocElement
// 	head := NewWizDocTitle(title)
// 	blocks = append(blocks, head)
// 	data := map[string]interface{}{
// 		"blocks": blocks,
// 	}
// 	return NewWizEditorObjectWithData(data)
// }

// func NewWizEditorObjectWithTitleAndPlainText(title, content string) *WizEditorObject {
// 	var blocks []WizDocElement
// 	head := NewWizDocTitle(title)
// 	blocks = append(blocks, head)
// 	plainText := NewWizText([]*TextItem{{Insert: content}}, 0, "")
// 	blocks = append(blocks, plainText)
// 	data := map[string]interface{}{
// 		"blocks": blocks,
// 	}
// 	return NewWizEditorObjectWithData(data)
// }

type Base struct {
	Id                   string
	CreatorId            string
	CreationTime         int64 // 创建时间
	LastModificationTime int64 // 最后修改时间
}

// User 用户信息
type User struct {
	Id         string               `json:"id"`
	Type       ConfluenceEntityType `json:"-"`
	Name       string               `json:"name"`
	Email      string               `json:"email"`
	LowerEmail string               `json:"lower_email"`

	Status      string `json:"status"`      // xml 数据包中没有该属性，api 中可以获取
	DisplayName string `json:"displayName"` // xml 数据包中没有该属性，api 中可以获取

	OnesUUID   string `json:"ones_uuid"`
	OnesName   string `json:"ones_name"`
	OnesEmail  string `json:"ones_email"`
	OnesStatus int    `json:"ones_status"`
	Action     int    `json:"action"`
}

func (u *User) GetId() string {
	return u.Id
}

func (u *User) GetType() ConfluenceEntityType {
	return u.Type
}

func (u *User) ValidOnesUser() bool {
	return len(u.OnesUUID) != 0
}

// // NewUserByConfluenceData 使用解析出来的 xml 数据初始化 User 对象
// func NewUserByConfluenceData(data *ObjectData) (ConfluenceDataModel, error) {
// 	m := confluenceXMLPropertySliceToMap(data.Properties)
// 	u := &User{}
// 	u.Id = data.Id
// 	u.Type = ConfluenceEntityTypeUser
// 	if p, ok := m["name"]; ok {
// 		u.Name = p.Value
// 	} else {
// 		return nil, errors.Trace(errors.MalformedError(errors.XML))
// 	}
// 	if p, ok := m["email"]; ok {
// 		u.Email = p.Value
// 		u.LowerEmail = strings.ToLower(u.Email)
// 	}
// 	return u, nil
// }

// // Space 空间相关信息，忽略了权限相关的数据
// // todo 没有填充创建者、时间等
// type Space struct {
// 	Base
// 	Type        ConfluenceEntityType `json:"-"`
// 	Name        string               `json:"name"`
// 	Key         string               `json:"key"`
// 	SpaceType   string               `json:"space_type"`
// 	SpaceStatus string               `json:"space_status"`
// 	HomePageId  string               `json:"home_page_id"`
// 	Permissions []*SpacePermission   `json:"space_permissions"`
// }

// func (s *Space) GetId() string {
// 	return s.Id
// }

// func (s *Space) GetType() ConfluenceEntityType {
// 	return s.Type
// }

// func (s *Space) NoHomePage() bool {
// 	return len(s.HomePageId) == 0
// }

// func NewSpaceByConfluenceData(data *ObjectData) (ConfluenceDataModel, error) {
// 	m := confluenceXMLPropertySliceToMap(data.Properties)
// 	s := &Space{}
// 	s.Id = data.Id
// 	s.Type = ConfluenceEntityTypeSpace
// 	if p, ok := m["name"]; ok {
// 		s.Name = p.Value
// 	} else {
// 		return nil, errors.Trace(errors.MalformedError(errors.XML))
// 	}
// 	if p, ok := m["key"]; ok {
// 		s.Key = p.Value
// 	} else {
// 		return nil, errors.Trace(errors.MalformedError(errors.XML))
// 	}
// 	if p, ok := m["spaceType"]; ok {
// 		s.SpaceType = p.Value
// 	}
// 	if p, ok := m["spaceStatus"]; ok {
// 		s.SpaceStatus = p.Value
// 	}
// 	if p, ok := m["homePage"]; ok {
// 		s.HomePageId = p.Id
// 	}
// 	return s, nil
// }

// // Page 页面信息，目前有两种：普通页面、博文
// // 两种数据xml结构基本相同
// type Page struct {
// 	Base
// 	WikiUUID                string                                `json:"wiki_uuid"`        // 在 ones wiki 中的 UUID
// 	WikiParentUUID          string                                `json:"wiki_parent_uuid"` // 在 ones wiki 中的 parent UUID
// 	WizDocId                string                                `json:"wiz_doc_id"`       // 在 wiz 中的 id
// 	WizEditorObject         *WizEditorObject                      `json:"-"`                // 编辑器对象，用来生成文档 json
// 	Type                    ConfluenceEntityType                  `json:"-"`
// 	Title                   string                                `json:"title"`            // 标题
// 	TitleNumPrefix          *int                                  `json:"title_num_prefix"` // 标题字符串开头的数字，用来进行辅助排序页面
// 	BodyContentIds          []string                              `json:"body_content_ids"` // 正文对象 BodyContent 的 id
// 	AttachmentIds           []string                              `json:"attachment_ids,omitempty"`
// 	Attachments             map[string]map[string]*AttachmentInfo `json:"attachments" `                      // 页面内所有附件的所有历史版本信息，结构：文件名 - 版本号 - 附件信息
// 	ImagesNeedUpload        map[string]*FileUploadInfo            `json:"images_need_upload,omitempty"`      // 需要上传的图片 文件名->文件
// 	AttachmentsNeedUpload   map[string]*FileUploadInfo            `json:"attachments_need_upload,omitempty"` // 需要上传的附件 文件名->文件
// 	Position                int64                                 `json:"position"`
// 	ParentPageId            string                                `json:"parent_page_id"`
// 	Children                []*Page                               `json:"-"` // 子页面
// 	ChildrenPageIds         []string                              `json:"children_page_ids,omitempty"`
// 	IgnoreTransfer          bool                                  `json:"ignore_transfer"` // 已经包含内容，不需要再进行转化。忽略 transferConfluencePage
// 	ContentPermissionSetIds []string                              `json:"content_permission_set"`
// 	ContentPermissionSet    []*ContentPermissionSet               `json:"-"`
// 	EncryptStatus           int                                   `json:"encrypt_status"`
// 	WikiOwnerUUID           string                                `json:"owner_uuid"`
// 	SpaceId                 string                                `json:"-"`

// 	// 解析页面内容后才会初始化的数据
// 	CodeSnippets map[string][]byte        `json:"-"` // page 正文中的代码片段，片段的uuid -> 片段正文
// 	CardAttrList []map[string]interface{} `json:"-"` // 页面内所有附件列表 card 的 attrs

// 	// 页面导入是否成功，并发导入时使用
// 	ImportSuccess       bool   `json:"import_success,omitempty"` // 页面所有内容导入成功
// 	CardUUID            string `json:"card_uuid,omitempty"`
// 	NotSupportMacroName []string
// 	Notifications       []*Notification   // 关注者
// 	ReferSpace          []*ReferSpacePage // referSpace
// 	LayOut              bool              //样式
// }

// func NewHiddenRootPage() *Page {
// 	title := i18n.TranslateWithGls("confluence_import_hidden_title", nil)
// 	content := i18n.TranslateWithGls("confluence_import_hidden_content", nil)
// 	var hiddenRoot *Page = &Page{
// 		Base: Base{
// 			Id:                   "",
// 			CreatorId:            "",
// 			LastModificationTime: 0,
// 		},
// 		WikiUUID:        utils.UUID(),
// 		WikiParentUUID:  "",
// 		ParentPageId:    "",
// 		Children:        []*Page{},
// 		ChildrenPageIds: []string{},
// 		WizEditorObject: NewWizEditorObjectWithTitleAndPlainText(title, content),
// 		Type:            ConfluenceEntityTypePage,
// 		Title:           title,
// 		IgnoreTransfer:  true,
// 		WikiOwnerUUID:   GetUserUUIDFromContext(),
// 		EncryptStatus:   utilsModel.PageEncryptStatusNone,
// 	}
// 	return hiddenRoot
// }

// func NewPlainPage(title string) *Page {
// 	var hiddenRoot = &Page{
// 		Base: Base{
// 			Id:                   "",
// 			CreatorId:            "",
// 			LastModificationTime: 0,
// 		},
// 		WikiUUID:        utils.UUID(),
// 		WikiParentUUID:  "",
// 		ParentPageId:    "",
// 		Children:        []*Page{},
// 		ChildrenPageIds: []string{},
// 		WizEditorObject: NewWizEditorObjectWithTitleAndPlainText(title, " "),
// 		Type:            ConfluenceEntityTypePage,
// 		Title:           title,
// 		IgnoreTransfer:  true,
// 		WikiOwnerUUID:   GetUserUUIDFromContext(),
// 		EncryptStatus:   utilsModel.PageEncryptStatusNone,
// 	}
// 	return hiddenRoot
// }

// type FileUploadInfo struct {
// 	ZipFile       *zip.File `json:"-"`
// 	FileName      string    `json:"file_name"`
// 	FilePathInZip string    `json:"file_path_in_zip"`
// 	Hash          string    `json:"hash,omitempty"`          // 导入图片的时候使用，附件不使用
// 	Size          int64     `json:"size,omitempty"`          // 导入图片的时候使用，附件不使用
// 	ResourceUUID  string    `json:"resource-uuid,omitempty"` // 附件在 project 数据表 resource 中记录的 uuid
// 	Uploaded      bool      `json:"uploaded"`
// }

// func (p *Page) GetId() string {
// 	return p.Id
// }

// func (p *Page) GetType() ConfluenceEntityType {
// 	return p.Type
// }

// func NewPage() *Page {
// 	page := &Page{
// 		WizEditorObject: NewWizEditorObject(),
// 	}
// 	return page
// }

// func NewPageByConfluenceData(data *ObjectData) (ConfluenceDataModel, error) {
// 	// 填充属性
// 	properties := confluenceXMLPropertySliceToMap(data.Properties)
// 	var p *Page
// 	if t, ok := properties["originalVersionId"]; ok {
// 		if t.Value != "" {
// 			// 不处理历史版本
// 			return nil, nil
// 		}
// 	}
// 	if t, ok := properties["originalVersion"]; ok {
// 		// 兼容旧 confluence 中没有 originalVersionId 的情况，eg: 6.0.6
// 		if t.Id != "" {
// 			// 不处理历史版本
// 			return nil, nil
// 		}
// 	}
// 	if t, ok := properties["contentStatus"]; ok {
// 		if t.Value == ConfluencePageContentTypeDraft || t.Value == ConfluencePageContentTypeDeleted {
// 			// 不处理草稿和已删除状态的页面
// 			// 应该可以只导入状态为 current 的页面，但是没有官方文档可以依据
// 			return nil, nil
// 		}
// 	}
// 	if p == nil {
// 		p = NewPage()
// 	}
// 	p.ImagesNeedUpload = make(map[string]*FileUploadInfo)
// 	p.AttachmentsNeedUpload = make(map[string]*FileUploadInfo)
// 	p.Id = data.Id
// 	p.WikiUUID = utils.UUID()
// 	if data.Type == ConfluenceEntityClassPage {
// 		p.Type = ConfluenceEntityTypePage
// 	} else if data.Type == ConfluenceEntityClassBlogPost {
// 		p.Type = ConfluenceEntityTypeBlogPost
// 	} else {
// 		return nil, nil
// 	}
// 	if t, ok := properties["title"]; ok {
// 		p.Title = t.Value
// 	} else {
// 		return nil, errors.Trace(errors.MalformedError(errors.XML))
// 	}
// 	if t, ok := properties["parent"]; ok {
// 		p.ParentPageId = t.Id
// 	}
// 	if t, ok := properties["creator"]; ok {
// 		p.CreatorId = t.Id
// 	}
// 	// 创建时间
// 	if t, ok := properties["creationDate"]; ok && t.Value != "" {
// 		milli, err := parseTimeToMilli(t.Value)
// 		if err != nil {
// 			return nil, errors.Trace(err)
// 		}
// 		p.CreationTime = milli
// 	}
// 	// 修改时间
// 	if t, ok := properties["lastModificationDate"]; ok && t.Value != "" {
// 		milli, err := parseTimeToMilli(t.Value)
// 		if err != nil {
// 			return nil, errors.Trace(err)
// 		}
// 		p.LastModificationTime = milli
// 	}
// 	if t, ok := properties["position"]; ok {
// 		var v int64
// 		var err error
// 		if t.Value != "" {
// 			v, err = strconv.ParseInt(t.Value, 10, 64)
// 			if err != nil {
// 				return nil, errors.Wrap(err, errors.MalformedError(errors.XML))
// 			}
// 		}
// 		p.Position = v
// 	}
// 	// space id
// 	if t, ok := properties["space"]; ok {
// 		p.SpaceId = t.Id
// 	}
// 	// 填充关联关系
// 	collections := confluenceXMLCollectionSliceToMap(data.Collections)
// 	// 正文
// 	if body, ok := collections["bodyContents"]; ok {
// 		p.BodyContentIds = body
// 	}
// 	// 附件，包括正文内的图片
// 	if attachments, ok := collections["attachments"]; ok {
// 		p.AttachmentIds = attachments
// 	}
// 	// 子页面
// 	if children, ok := collections["children"]; ok {
// 		p.ChildrenPageIds = children
// 	}
// 	// 权限
// 	if permissionSets, ok := collections["contentPermissionSets"]; ok {
// 		p.ContentPermissionSetIds = permissionSets
// 	}
// 	return p, nil
// }

// // BodyContent 内容，正文内容或者评论内容
// type BodyContent struct {
// 	Id           string
// 	Type         ConfluenceEntityType
// 	ContentClass string
// 	Body         string // 正文内容
// }

// func (b *BodyContent) GetId() string {
// 	return b.Id
// }

// func (b *BodyContent) GetType() ConfluenceEntityType {
// 	return b.Type
// }

// func NewBodyContentByConfluenceData(data *ObjectData) (ConfluenceDataModel, error) {
// 	m := confluenceXMLPropertySliceToMap(data.Properties)
// 	b := &BodyContent{}
// 	b.Id = data.Id
// 	b.Type = ConfluenceEntityTypeBodyContent
// 	if p, ok := m["body"]; ok {
// 		b.Body = p.Value
// 	} else {
// 		return nil, errors.Trace(errors.MalformedError(errors.XML))
// 	}
// 	if p, ok := m["content"]; ok {
// 		b.ContentClass = p.Class
// 	}
// 	return b, nil
// }

// type AttachmentInfo struct {
// 	Base
// 	Type              ConfluenceEntityType `json:"-"`
// 	Title             string               `json:"title"`
// 	Version           string               `json:"version"`
// 	VersionComment    string               `json:"version_comment"`     // 版本注释，目前仅用来判断 drawio 的版本号
// 	OriginalVersionId string               `json:"original_version_id"` // 最新版本该属性为空
// 	OriginalVersion   string               `json:"original_version"`    // 部分 confluence 旧版本需要根据此字段判断是否最新版本，最新版本该属性为空
// 	SpaceId           string               `json:"-"`
// }

// func (a *AttachmentInfo) GetType() ConfluenceEntityType {
// 	return a.Type
// }

// func (a *AttachmentInfo) GetId() string {
// 	return a.Id
// }

// // IsNewestVersion 是否是最新版本
// func (a *AttachmentInfo) IsNewestVersion() bool {
// 	return a.OriginalVersionId == "" && a.OriginalVersion == ""
// }

// func NewAttachmentByConfluenceData(data *ObjectData) (ConfluenceDataModel, error) {
// 	m := confluenceXMLPropertySliceToMap(data.Properties)
// 	a := &AttachmentInfo{}
// 	a.Id = data.Id
// 	a.Type = ConfluenceEntityTypeAttachment
// 	if t, ok := m["title"]; ok {
// 		a.Title = t.Value
// 	} else {
// 		return nil, errors.Trace(errors.MalformedError(errors.XML))
// 	}
// 	if t, ok := m["version"]; ok && t.Value != "" {
// 		a.Version = t.Value
// 	} else {
// 		return nil, errors.Trace(errors.MalformedError(errors.XML))
// 	}
// 	if t, ok := m["versionComment"]; ok && t.Value != "" {
// 		a.VersionComment = t.Value
// 	}
// 	if t, ok := m["originalVersionId"]; ok {
// 		a.OriginalVersionId = t.Value
// 	}
// 	if t, ok := m["originalVersion"]; ok {
// 		a.OriginalVersion = t.Id
// 	}
// 	// space id
// 	if t, ok := m["space"]; ok {
// 		a.SpaceId = t.Id
// 	}
// 	return a, nil
// }

// // Label 标签
// type Label struct {
// 	Base
// 	Name string
// 	Type ConfluenceEntityType
// }

// func (l *Label) GetType() ConfluenceEntityType {
// 	return l.Type
// }

// func (l *Label) GetId() string {
// 	return l.Id
// }

// func NewLabelByConfluenceData(data *ObjectData) (ConfluenceDataModel, error) {
// 	m := confluenceXMLPropertySliceToMap(data.Properties)
// 	label := &Label{
// 		Type: ConfluenceEntityTypeLabel,
// 	}
// 	label.Id = data.Id
// 	if t, ok := m["name"]; ok && t.Value != "" {
// 		label.Name = t.Value
// 	} else {
// 		return nil, nil
// 	}
// 	return label, nil
// }

// // Labelling 页面标签关系
// type Labelling struct {
// 	Base
// 	Type      ConfluenceEntityType
// 	LabelId   string
// 	ContentId string
// }

// func (l *Labelling) GetType() ConfluenceEntityType {
// 	return l.Type
// }

// func (l *Labelling) GetId() string {
// 	return l.Id
// }

// func NewLabelingByConfluenceData(data *ObjectData) (ConfluenceDataModel, error) {
// 	m := confluenceXMLPropertySliceToMap(data.Properties)
// 	label := &Labelling{
// 		Type: ConfluenceEntityTypeLabelling,
// 	}
// 	if t, ok := m["content"]; ok {
// 		if t.Class != ConfluenceEntityClassPage {
// 			// 非页面标签，忽略
// 			return nil, nil
// 		}
// 		if t.Id != "" {
// 			label.ContentId = t.Id
// 		}
// 	}
// 	label.Id = data.Id
// 	if t, ok := m["label"]; ok && t.Id != "" {
// 		label.LabelId = t.Id
// 	} else {
// 		return nil, nil
// 	}
// 	return label, nil
// }

// // SpacePermission 空间权限
// type SpacePermission struct {
// 	Base
// 	Type            ConfluenceEntityType
// 	PermissionType  string
// 	SpaceId         string
// 	GroupName       string
// 	UserSubject     string
// 	AllUsersSubject string
// 	CreationDate    string
// 	LastModifier    string
// }

// func (l *SpacePermission) GetType() ConfluenceEntityType {
// 	return l.Type
// }

// func (l *SpacePermission) GetId() string {
// 	return l.Id
// }

// func NewSpacePermissionByConfluenceData(data *ObjectData) (ConfluenceDataModel, error) {
// 	m := confluenceXMLPropertySliceToMap(data.Properties)
// 	spacePermission := &SpacePermission{
// 		Type: ConfluenceEntityTypeSpacePermission,
// 	}
// 	spacePermission.Id = data.Id

// 	if t, ok := m["space"]; ok {
// 		if t.Id != "" {
// 			spacePermission.SpaceId = t.Id
// 		}
// 	}
// 	if t, ok := m["group"]; ok && t.Value != "" {
// 		spacePermission.GroupName = t.Value
// 	}
// 	if t, ok := m["userSubject"]; ok {
// 		if t.Id != "" {
// 			spacePermission.UserSubject = t.Id
// 		}
// 	}
// 	if t, ok := m["allUsersSubject"]; ok {
// 		if t.Id != "" {
// 			spacePermission.AllUsersSubject = t.Id
// 		}
// 	}
// 	if t, ok := m["type"]; ok {
// 		if t.Value != "" {
// 			spacePermission.PermissionType = t.Value
// 		}
// 	}
// 	if t, ok := m["creator"]; ok {
// 		if t.Id != "" {
// 			spacePermission.CreatorId = t.Id
// 		}
// 	}
// 	if t, ok := m["creationDate"]; ok && t.Value != "" {
// 		spacePermission.CreationDate = t.Value

// 	}
// 	if t, ok := m["lastModifier"]; ok {
// 		if t.Id != "" {
// 			spacePermission.LastModifier = t.Id
// 		}
// 	}

// 	// 修改时间
// 	if t, ok := m["lastModificationDate"]; ok && t.Value != "" {
// 		milli, err := parseTimeToMilli(t.Value)
// 		if err != nil {
// 			return nil, errors.Trace(err)
// 		}
// 		spacePermission.LastModificationTime = milli
// 	}

// 	return spacePermission, nil
// }

// type ContentPermissionSet struct {
// 	Base
// 	Type                  ConfluenceEntityType `json:"-"`
// 	OwningContentClass    string               `json:"owning_content_class"` // 权限归属的对象类型，eg：page
// 	OwningContentId       string               `json:"owning_content_id"`    // 权限归属的 id，eg：如果 class 是 page，则为 page id
// 	PermissionType        string               `json:"permission_type"`
// 	ContentPermissions    []string             `json:"content_permissions"` // ContentPermission 的 id
// 	ContentPermissionData []*ContentPermission `json:"-"`                   // 解析后填充
// }

// func (cps *ContentPermissionSet) GetType() ConfluenceEntityType {
// 	return cps.Type
// }

// func (cps *ContentPermissionSet) GetId() string {
// 	return cps.Id
// }

// func NewContentPermissionSetConfluenceData(data *ObjectData) (ConfluenceDataModel, error) {
// 	m := confluenceXMLPropertySliceToMap(data.Properties)
// 	s := &ContentPermissionSet{}
// 	s.Id = data.Id
// 	s.Type = ConfluenceEntityTypeContentPermissionSet
// 	if p, ok := m["type"]; ok {
// 		s.PermissionType = p.Value
// 	} else {
// 		return nil, errors.Trace(errors.MalformedError(errors.XML))
// 	}
// 	if p, ok := m["owningContent"]; ok {
// 		s.OwningContentId = p.Id
// 		s.OwningContentClass = p.Class
// 	} else {
// 		return nil, errors.Trace(errors.MalformedError(errors.XML))
// 	}
// 	// 填充关联关系
// 	collections := confluenceXMLCollectionSliceToMap(data.Collections)
// 	// 正文
// 	if permissions, ok := collections["contentPermissions"]; ok {
// 		s.ContentPermissions = permissions
// 	}
// 	return s, nil
// }

// type ContentPermission struct {
// 	Base
// 	Type           ConfluenceEntityType `json:"-"`
// 	UserSubject    string               `json:"owning_user"`
// 	GroupName      string               `json:"owning_group"`
// 	PermissionType string               `json:"permission_type"`
// }

// func (cps *ContentPermission) GetType() ConfluenceEntityType {
// 	return cps.Type
// }

// func (cps *ContentPermission) GetId() string {
// 	return cps.Id
// }

// func (cp *ContentPermission) GetPlatUsers(cd *ConfluenceData, filterOnes bool) (users []*User) {
// 	if cp.IsSingleUser() {
// 		user := cd.GetUserById(cp.UserSubject, filterOnes)
// 		if user != nil {
// 			users = append(users, user)
// 		}
// 	} else if cp.IsGroup() {
// 		users = cd.GetGroupUsers(cp.GroupName, filterOnes)
// 	}
// 	return
// }

// func (cp *ContentPermission) IsSingleUser() bool {
// 	return len(cp.UserSubject) != 0
// }

// func (cp *ContentPermission) IsGroup() bool {
// 	return len(cp.GroupName) != 0
// }

// func NewContentPermissionConfluenceData(data *ObjectData) (ConfluenceDataModel, error) {
// 	m := confluenceXMLPropertySliceToMap(data.Properties)
// 	s := &ContentPermission{}
// 	s.Id = data.Id
// 	s.Type = ConfluenceEntityTypeContentPermission
// 	if p, ok := m["type"]; ok {
// 		s.PermissionType = p.Value
// 	} else {
// 		return nil, errors.Trace(errors.MalformedError(errors.XML))
// 	}
// 	if p, ok := m["userSubject"]; ok {
// 		s.UserSubject = p.Id
// 	}
// 	if p, ok := m["groupName"]; ok {
// 		s.GroupName = p.Value
// 	}
// 	return s, nil
// }

// type Descriptor map[string]string

// func (d Descriptor) getOrPanic(key string) string {
// 	if v, ok := d[key]; ok {
// 		return v
// 	}
// 	panic(fmt.Errorf("no key: %s in descriptor", key))
// }

// func (d Descriptor) SpaceKey() string {
// 	return d.getOrPanic(ConfluenceDescriptorSpaceKey)
// }

// func (d Descriptor) DefaultUsersGroup() string {
// 	return d.getOrPanic(ConfluenceDescriptorDefaultUsersGroup)
// }

// func (d Descriptor) IsDefaultUsersGroup(groupName string) bool {
// 	return d.DefaultUsersGroup() == groupName
// }

// //关注者
// type Notification struct {
// 	Base
// 	Type         ConfluenceEntityType
// 	ContentId    string
// 	Receiver     string
// 	CreationDate int64
// 	LastModifier string
// 	Digest       string
// 	Network      string
// }

// func (l *Notification) GetType() ConfluenceEntityType {
// 	return l.Type
// }

// func (l *Notification) GetId() string {
// 	return l.Id
// }
// func NewNotificationByConfluenceData(data *ObjectData) (ConfluenceDataModel, error) {
// 	m := confluenceXMLPropertySliceToMap(data.Properties)
// 	notification := &Notification{
// 		Type: ConfluenceEntityTypeNotification,
// 	}
// 	notification.Id = data.Id
// 	if t, ok := m["content"]; ok {
// 		if t.Class != ConfluenceEntityClassPage {
// 			// 非页面标签，忽略
// 			return nil, nil
// 		}
// 		if t.Id != "" {
// 			notification.ContentId = t.Id
// 		}
// 	}
// 	if t, ok := m["receiver"]; ok {
// 		if t.Class != ConfluenceEntityClassUser {
// 			return nil, nil
// 		}
// 		if t.Id != "" {
// 			notification.Receiver = t.Id
// 		}
// 	}
// 	if t, ok := m["creator"]; ok {
// 		if t.Class != ConfluenceEntityClassUser {
// 			// 非页面标签，忽略
// 			return nil, nil
// 		}
// 		if t.Id != "" {
// 			notification.CreatorId = t.Id
// 		}
// 	}

// 	if t, ok := m["creationDate"]; ok && t.Value != "" {
// 		milli, err := parseTimeToMilli(t.Value)
// 		if err != nil {
// 			return nil, errors.Trace(err)
// 		}
// 		notification.CreationDate = milli
// 	}

// 	if t, ok := m["lastModifier"]; ok {
// 		if t.Class != ConfluenceEntityClassUser {
// 			// 非页面标签，忽略
// 			return nil, nil
// 		}
// 		if t.Id != "" {
// 			notification.LastModifier = t.Id
// 		}
// 	}
// 	// 修改时间
// 	if t, ok := m["lastModificationDate"]; ok && t.Value != "" {
// 		milli, err := parseTimeToMilli(t.Value)
// 		if err != nil {
// 			return nil, errors.Trace(err)
// 		}
// 		notification.LastModificationTime = milli
// 	}

// 	if t, ok := m["digest"]; ok && t.Value != "" {
// 		notification.Digest = t.Value
// 	}

// 	if t, ok := m["network"]; ok && t.Value != "" {
// 		notification.Network = t.Value
// 	}

// 	return notification, nil
// }
