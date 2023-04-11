package importserv

const (
	SourceTypeConfluence = "confluence"
	SourceTypeWord       = "word"
	SourceTypePage       = "page"
	SourceTypeTemplate   = "template"

	ImportDoingDataType    = "ImportDoing"
	ImportFailedDataType   = "ImportFailed"
	ImportFinishedDataType = "ImportFinish"
)

const (
	ImportTypeFullSite   = "full_site"
	ImportTypeSpace      = "space"
	ImportTypeConfluence = "confluence"
	ImportTypePage       = "page"
	ImportTypeWord       = "word"
	ImportTypeWiz        = "wiz"
)

const (
	ParseTimeOutSecond = 55
)

const (
	ImportRetryCountMax = 5
)

const (
	ConfluenceDirAttachments = "attachments"
	ConfluenceDirImages      = "images"
	ConfluencePageIndex      = "index.html"
	ConfluenceHomePage       = "home_page"

	ConfluenceResourceID          = "data-linked-resource-id"
	ConfluenceResourceContentType = "data-linked-resource-content-type"
	ConfluenceDataEmoticonName    = "data-emoticon-name"
	ConfluenceDrawioDiagramImage  = "drawio-diagram-image"
	ConfluenceEmoticonClass       = "emoticon"

	ImgSrc = "src"
)

const (
	ErrorRecordFileNamePrefix = "errors_import"
)

// func ErrorConfluenceRecordExcelHeaders() []string {
// 	return []string{
// 		i18n.TG("service.import.import_file_name"),
// 		i18n.TG("service.import.space_name"),
// 		i18n.TG("service.import.page_name"),
// 		i18n.TG("service.import.reference_info"),
// 	}
// }
// func ErrorWordRecordExcelHeaders() []string {
// 	return []string{
// 		i18n.TG("service.import.import_file_name"),
// 		i18n.TG("service.import.page_group_name"),
// 		i18n.TG("service.import.doc_name"),
// 		i18n.TG("service.import.reference_info"),
// 	}
// }

const (
	ErrAttachmentUpload  = "err_attachment_upload"
	ErrLimitExceeded     = "err_limit_exceeded"
	ErrPageUpload        = "err_page_upload"
	ErrPageParse         = "err_page_parse"
	ErrCardSave          = "err_card_save"
	ErrPageSave          = "err_page_save"
	ErrAttachmentImport  = "err_attachment_import"
	ErrImageImport       = "err_image_import"
	ErrPageRestriction   = "err_page_restriction"
	ErrUnknown           = "err_unknown"
	ErrImportWordForWiz  = "err_import_word_for_wiz"
	ErrImportWordForWiki = "err_import_word_for_wiki"
	ErrMarshal           = "err_marshal"
)

var MapErrMessage = map[string]string{
	ErrAttachmentUpload:  "service.import.err_attachment_upload",
	ErrLimitExceeded:     "service.import.err_limit_exceeded",
	ErrPageUpload:        "service.import.err_page_upload",
	ErrPageParse:         "service.import.err_page_parse",
	ErrUnknown:           "service.import.err_unknown",
	ErrPageSave:          "service.import.err_page_save",
	ErrAttachmentImport:  "service.import.err_attachment_save",
	ErrImageImport:       "service.import.err_image_save",
	ErrMarshal:           "service.import.err_marshal",
	ErrImportWordForWiz:  "service.import.err_import_word_for_wiz",
	ErrImportWordForWiki: "service.import.err_import_word_for_wiki",
}

const (
	ONESWikiFigureSizeSmall  = "small"
	ONESWikiFigureSizeMedium = "medium"
	ONESWikiFigureSizeLarge  = "large"
	ONESWikiFigure           = `
<figure class="ones-image-figure" data-size="%s">
	<div class="image-wrapper"><img data-mime="%s" data-or="1" data-ref-id="%s" data-ref-type="space"
			data-size="%s" data-uuid="%s" src="data:image/gif;base64,R0lGODlhAQABAAD/ACwAAAAAAQABAAACADs=" />
	</div>
	<figcaption></figcaption>
</figure>
`
	ONESWikiFileBlock = `
<p><span class="ones-file-block" data-name="%s" data-ref-id="%s" data-ref-type="space"
		data-uuid="%s">&nbsp;</span></p>
`
)

const (
	PostActionBindTask = "bind_task"
)

const (
	ConfluenceXMLExportEntityXML          = "entities.xml"                // xml 形式导出数据文件名
	ConfluenceXMLExportExportDescriptor   = "exportDescriptor.properties" // 导出包的属性
	ConfluenceDescriptorSpaceKey          = "spaceKey"                    // 导出空间的 key
	ConfluenceDescriptorDefaultUsersGroup = "defaultUsersGroup"           // 导出空间的默认用户组名
)

const (
	ConfluenceEntityElementObject = "object" // entities.xml 中 object 元素名
)

const (
	ConfluenceEntityElementAttrPackage = "package" // object 元素 package
	ConfluenceEntityElementAttrClass   = "class"   // object 元素 class属性
)

const (
	ConfluenceEntityPackagePages        = "com.atlassian.confluence.pages"             // 页面
	ConfluenceEntityPackageSpaces       = "com.atlassian.confluence.spaces"            // 空间
	ConfluenceEntityPackageUsers        = "com.atlassian.confluence.user"              // 用户
	ConfluenceEntityPackageCore         = "com.atlassian.confluence.core"              // 正文
	ConfluenceEntityPackageContent      = "com.atlassian.confluence.content"           // 正文
	ConfluenceEntityPackageLabel        = "com.atlassian.confluence.labels"            // 标签
	ConfluenceEntityPackageSecurity     = "com.atlassian.confluence.security"          // 权限
	ConfluenceEntityPackageNotification = "com.atlassian.confluence.mail.notification" // 标签
)

const (
	ConfluenceEntityClassPage                 = "Page"                 // 页面类名
	ConfluenceEntityClassBlogPost             = "BlogPost"             // 博文类型
	ConfluenceEntityClassSpace                = "Space"                // 空间类名
	ConfluenceEntityClassUser                 = "ConfluenceUserImpl"   // 用户类名
	ConfluenceEntityClassBodyContent          = "BodyContent"          // 页面正文、评论正文类
	ConfluenceEntityClassAttachment           = "Attachment"           // 附件类名
	ConfluenceEntityClassComment              = "Comment"              // 评论类名
	ConfluenceEntityClassContentProperty      = "ContentProperty"      // 正文属性类名
	ConfluenceEntityClassLabel                = "Label"                // 标签
	ConfluenceEntityClassLabelling            = "Labelling"            // 标签-页面关系
	ConfluenceEntityClassSpacePermission      = "SpacePermission"      // 空间权限
	ConfluenceEntityClassContentPermission    = "ContentPermission"    // 权限
	ConfluenceEntityClassContentPermissionSet = "ContentPermissionSet" // 权限集
	ConfluenceEntityClassNotification         = "Notification"         // 关注者
)

type ConfluenceEntityType string

const (
	ConfluenceEntityTypePage                 ConfluenceEntityType = "Page"                 // 页面
	ConfluenceEntityTypeBlogPost             ConfluenceEntityType = "BlogPost"             // 博文
	ConfluenceEntityTypeSpace                ConfluenceEntityType = "Space"                // 空间
	ConfluenceEntityTypeUser                 ConfluenceEntityType = "User"                 // 用户
	ConfluenceEntityTypeBodyContent          ConfluenceEntityType = "BodyContent"          // 正文
	ConfluenceEntityTypeAttachment           ConfluenceEntityType = "Attachment"           // 附件
	ConfluenceEntityTypeComment              ConfluenceEntityType = "Comment"              // 评论
	ConfluenceEntityTypeContentProperty      ConfluenceEntityType = "ContentProperty"      // 正文属性
	ConfluenceEntityTypeLabel                ConfluenceEntityType = "Label"                // 评论
	ConfluenceEntityTypeLabelling            ConfluenceEntityType = "Labelling"            // 评论
	ConfluenceEntityTypeSpacePermission      ConfluenceEntityType = "SpacePermission"      // 空间权限
	ConfluenceEntityTypeContentPermission    ConfluenceEntityType = "ContentPermission"    // 权限
	ConfluenceEntityTypeContentPermissionSet ConfluenceEntityType = "ContentPermissionSet" // 权限集
	ConfluenceEntityTypeNotification         ConfluenceEntityType = "Notification"         // 关注者
)

const (
	ConfluencePageContentTypeDraft   = "draft"
	ConfluencePageContentTypeDeleted = "deleted"
)

const (
	WizDocStructureBlocks = "blocks"
)

const (
	WizDocElementText          = "text"
	WizDocElementEmbed         = "embed"
	WizDocElementList          = "list"
	WizDocElementTask          = "list"
	WizDocElementCode          = "code"
	WizDocElementCodeLine      = "code-line"
	WizDocElementTable         = "table"
	WizDocElementInlineComment = "inline-comment"
	WizDocElementBr            = "br"
)

const (
	WizDocElementEmbedTypeImage    = "image"
	WizDocElementEmbedTypeDrawIO   = "drawio"
	WizDocElementEmbedTypeVideo    = "video"
	WizDocElementEmbedTypeAudio    = "audio"
	WizDocElementEmbedTypeMarkdown = "markdown"
	WizDocElementEmbedTypeUnknown  = "unknown"
)

const (
	WizDocElementBoxTypeStatus  = "status"
	WizDocElementBoxTypeUnknown = "unknown"
	WizDocElementBoxTypeAnchor  = "anchor"
	WizDocElementBoxTypeMention = "mention"
	WizDocElementBoxTypeBr      = "br"
	WizDocElementBoxTypeImage   = "image"
	WizDocElementBoxTypeFile    = "file"
)

const (
	WizDocTextAlignLeft   = "left"
	WizDocTextAlignCenter = "center"
)

const (
	ConfluenceExportTypeHTML = "html"
	ConfluenceExportTypeXML  = "xml"
)

const (
	AttachmentTypeImage = "image"
	AttachmentTypeFile  = "file"
)

const (
	CdataTagHead = "[CDATA["
	CdataTagTail = "]]"
)

const (
	WizFontColorGrey = "style-color-6"
)

const (
	WizStyleUnderline = "style-underline"
	WizStyleBold      = "style-bold"
)

const (
	ConfluencePagePermissionView = "View"
	ConfluencePagePermissionEdit = "Edit"
)

const (
	UserChannelLength = 32
)
