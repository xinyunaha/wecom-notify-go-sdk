/**
 * @Author: xinyunaha
 * @File:  message
 * @Date: 2024/1/11 16:40
 * @Version: 1.0.0
 * @Description:
 */

package wecom_bot_notify

func NewTextMessage(content string, mentionedList []string, mentionedMobileList []string) *TextMessage {
	return &TextMessage{
		MsgType: TEXT,
		Text: Text{
			Content:             content,
			MentionedList:       mentionedList,
			MentionedMobileList: mentionedMobileList,
		},
	}
}

func NewMarkdownMessage(content string) *MarkdownMessage {
	return &MarkdownMessage{
		MsgType: MARKDOWN,
		Markdown: Markdown{
			Content: content,
		},
	}
}

func NewNewsMessage(articles []Articles) *NewsMessage {
	return &NewsMessage{
		MsgType: NEWS,
		News: News{
			Articles: articles,
		},
	}
}

func NewImageMessage(imgBase64, md5 string) *ImageMessage {
	return &ImageMessage{
		MsgType: IMAGE,
		Image: Image{
			Base64: imgBase64,
			Md5:    md5,
		},
	}
}

func NewFileMessage(mediaId string) *FileMessage {
	return &FileMessage{
		MsgType: FILE,
		File: File{
			MediaId: mediaId,
		},
	}
}

func NewVoiceMessage(mediaId string) *VoiceMessage {
	return &VoiceMessage{
		MsgType: VOICE,
		Voice: Voice{
			MediaId: mediaId,
		},
	}
}

func NewTextTemplateCardMessage(source Source, title MainTitle, emphasis Emphasis, quote Quote, subTitle string, content []HorizontalContent, jump []Jump, action CardAction) *TextTemplateCardMessage {
	return &TextTemplateCardMessage{
		MsgType: CARD,
		TemplateCard: TextTemplateCard{
			CardType:              TextCard,
			Source:                source,
			MainTitle:             title,
			EmphasisContent:       emphasis,
			QuoteArea:             quote,
			SubTitleText:          subTitle,
			HorizontalContentList: content,
			JumpList:              jump,
			CardAction:            action,
		},
	}
}

func NewPictureTemplateCardMessage(source Source, title MainTitle, image ImgText, quote Quote, vertical []VerticalContent, horizontal []HorizontalContent, jump []Jump, action CardAction) *PictureTemplateCardMessage {
	return &PictureTemplateCardMessage{
		MsgType: CARD,
		TemplateCard: PictureTemplateCard{
			CardType:              PicCard,
			Source:                source,
			MainTitle:             title,
			CardImage:             CardImage{},
			ImageTextArea:         image,
			QuoteArea:             quote,
			VerticalContentList:   vertical,
			HorizontalContentList: horizontal,
			JumpList:              jump,
			CardAction:            action,
		},
	}
}
