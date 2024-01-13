/**
 * @Author: xinyunaha
 * @File:  message
 * @Date: 2024/1/11 15:22
 * @Version: 1.0.0
 * @Description:
 */

package wecom_bot_notify

const (
	TEXT     = "text"
	MARKDOWN = "markdown"
	IMAGE    = "image"
	NEWS     = "news"
	FILE     = "file"
	VOICE    = "voice"
	CARD     = "template_card"
	PicCard  = "news_notice"
	TextCard = "text_notice"
	AtAll    = "@all"
)

// TextMessage 文本类型消息
type TextMessage struct {
	MsgType string `json:"msgtype"`
	Text    Text   `json:"text"`
}
type Text struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list"`
	MentionedMobileList []string `json:"mentioned_mobile_list"`
}

// MarkdownMessage markdown类型消息
type MarkdownMessage struct {
	MsgType  string   `json:"msgtype"`
	Markdown Markdown `json:"markdown"`
}
type Markdown struct {
	Content string `json:"content"`
}

// NewsMessage 图文消息
type NewsMessage struct {
	MsgType string `json:"msgtype"`
	News    News   `json:"news"`
}
type News struct {
	Articles []Articles `json:"articles"`
}
type Articles struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	PicUrl      string `json:"picurl"`
}

// ImageMessage 图片消息
type ImageMessage struct {
	MsgType string `json:"msgtype"`
	Image   Image  `json:"image"`
}
type Image struct {
	Base64 string `json:"base64"`
	Md5    string `json:"md5"`
}

// FileMessage 文件消息
type FileMessage struct {
	MsgType string `json:"msgtype"`
	File    File   `json:"file"`
}
type File struct {
	MediaId string `json:"media_id"`
}

// FileUploadResponse 文件上传返回
type FileUploadResponse struct {
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
	Type      string `json:"type"`       // 文件类型，分别有语音(voice)和普通文件(file)
	MediaId   string `json:"media_id"`   // 媒体文件上传后获取的唯一标识，3天内有效
	CreatedAt string `json:"created_at"` // 媒体文件上传时间戳
}

// VoiceMessage 语音消息
type VoiceMessage struct {
	MsgType string `json:"msgtype"`
	Voice   Voice  `json:"voice"`
}
type Voice struct {
	MediaId string `json:"media_id"`
}

// TextTemplateCardMessage 文本模板消息，详见 https://work.weixin.qq.com/api/doc/90000/90135/92135
type TextTemplateCardMessage struct {
	MsgType      string           `json:"msgtype"`
	TemplateCard TextTemplateCard `json:"template_card"`
}

// PictureTemplateCardMessage 图片模板消息，详见 https://work.weixin.qq.com/api/doc/90000/90135/92135
type PictureTemplateCardMessage struct {
	MsgType      string              `json:"msgtype"`
	TemplateCard PictureTemplateCard `json:"template_card"`
}
type TextTemplateCard struct {
	CardType              string              `json:"card_type"`
	Source                Source              `json:"source"`
	MainTitle             MainTitle           `json:"main_title"`
	EmphasisContent       Emphasis            `json:"emphasis_content"`
	QuoteArea             Quote               `json:"quote_area"`
	SubTitleText          string              `json:"sub_title_text"`
	HorizontalContentList []HorizontalContent `json:"horizontal_content_list"`
	JumpList              []Jump              `json:"jump_list"`
	CardAction            CardAction          `json:"card_action"`
}
type PictureTemplateCard struct {
	CardType              string              `json:"card_type"`
	Source                Source              `json:"source"`
	MainTitle             MainTitle           `json:"main_title"`
	CardImage             CardImage           `json:"card_image"`
	ImageTextArea         ImgText             `json:"image_text_area"`
	QuoteArea             Quote               `json:"quote_area"`
	VerticalContentList   []VerticalContent   `json:"vertical_content_list"`
	HorizontalContentList []HorizontalContent `json:"horizontal_content_list"`
	JumpList              []Jump              `json:"jump_list"`
	CardAction            CardAction          `json:"card_action"`
}
type Source struct {
	IconUrl   string `json:"icon_url"`
	Desc      string `json:"desc"`
	DescColor int    `json:"desc_color"`
}
type MainTitle struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}
type Emphasis struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}
type Quote struct {
	Type      int    `json:"type"`
	Url       string `json:"url"`
	Appid     string `json:"appid"`
	PagePath  string `json:"pagepath"`
	Title     string `json:"title"`
	QuoteText string `json:"quote_text"`
}
type HorizontalContent struct {
	KeyName string `json:"keyname"`
	Value   string `json:"value"`
	Type    int    `json:"type,omitempty"`
	Url     string `json:"url,omitempty"`
	MediaId string `json:"media_id,omitempty"`
}
type Jump struct {
	Type     int    `json:"type"`
	Url      string `json:"url,omitempty"`
	Title    string `json:"title"`
	Appid    string `json:"appid,omitempty"`
	PagePath string `json:"pagepath,omitempty"`
}
type CardAction struct {
	Type     int    `json:"type"`
	Url      string `json:"url"`
	Appid    string `json:"appid"`
	PagePath string `json:"pagepath"`
}
type VerticalContent struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}
type ImgText struct {
	Type     int    `json:"type"`
	Url      string `json:"url"`
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	ImageUrl string `json:"image_url"`
}
type CardImage struct {
	Url         string  `json:"url"`
	AspectRatio float64 `json:"aspect_ratio"`
}
