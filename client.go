/**
 * @Author: xinyunaha
 * @File:  client
 * @Date: 2024/1/11 15:12
 * @Version: 1.0.0
 * @Description:
 */

package wework_notify

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	weworknotify "github.com/xinyunaha/wework-notify-go-sdk/utils"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Client struct {
	token string
}

func NewWecomBotClient(token string) *Client {
	return &Client{token: token}
}

// Send 发送消息
func (client *Client) Send(msg any) error {
	url := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + client.token

	messageContent, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	data := strings.NewReader(string(messageContent))
	request, err := http.NewRequest(http.MethodPost, url, data)
	if err != nil {
		return err
	}

	httpClient := http.Client{}
	request.Header.Add("Content-Type", "application/json")
	response, err := httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	var resp struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}

	all, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(all, &resp); err != nil {
		return err
	}
	if resp.ErrCode != 0 {
		return errors.New(fmt.Sprintf("bad response: %s", resp.ErrMsg))
	}
	return nil
}

// UploadFile 上传文件 获取media_id
func (client *Client) UploadFile(fileType, filePath string) (FileUploadResponse, error) {
	if !weworknotify.Exists(filePath) {
		return FileUploadResponse{}, errors.New("file not exists")
	}
	var mediaInfo FileUploadResponse
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/upload_media?key=%s&type=%s", client.token, fileType)

	file, err := os.Open(filePath)
	if err != nil {
		return mediaInfo, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_ = writer.WriteField("filename", filepath.Base(filePath))
	_ = writer.WriteField("filelength", "")
	_ = writer.WriteField("name", "media")
	part, err := writer.CreateFormFile("media", filepath.Base(filePath))
	if _, err = io.Copy(part, file); err != nil {
		return mediaInfo, err
	}

	if err := writer.Close(); err != nil {
		return mediaInfo, err
	}

	request, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return mediaInfo, err
	}

	httpClient := http.Client{}
	request.Header.Add("Content-Type", writer.FormDataContentType())
	response, err := httpClient.Do(request)
	if err != nil {
		return mediaInfo, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return mediaInfo, errors.New("bad http status: " + response.Status)
	}
	resp, err := io.ReadAll(response.Body)
	if err != nil {
		return mediaInfo, err
	}
	if err := json.Unmarshal(resp, &mediaInfo); err != nil {
		return mediaInfo, err
	}
	if mediaInfo.ErrCode != 0 {
		return mediaInfo, errors.New(mediaInfo.ErrMsg)
	}
	return mediaInfo, nil
}

// SendTextMessage 发送文本消息
func (client *Client) SendTextMessage(content string, mentionedList []string, mentionedMobileList []string) error {
	message := NewTextMessage(content, mentionedList, mentionedMobileList)
	return client.Send(message)
}

// SendMarkdownMessage 发送markdown消息
func (client *Client) SendMarkdownMessage(content string) error {
	message := NewMarkdownMessage(content)
	return client.Send(message)
}

// SendImageMessage 发送图片消息
func (client *Client) SendImageMessage(base64, md5 string) error {
	message := NewImageMessage(base64, md5)
	return client.Send(message)
}

// SendNewsMessage 发送图文消息
func (client *Client) SendNewsMessage(articles []Articles) error {
	message := NewNewsMessage(articles)
	return client.Send(message)
}

// SendFileMessage 发送文件消息 如果mediaId为空，通过给定filePath进行上传并返回mediaId
func (client *Client) SendFileMessage(mediaId string, filePath string) (string, error) {
	if mediaId == "" {
		if !weworknotify.Exists(filePath) {
			return "", errors.New("file not exists")
		}
		mediaInfo, err := client.UploadFile(FILE, filePath)
		if err != nil {
			return "", err
		}
		mediaId = mediaInfo.MediaId
	}
	message := NewFileMessage(mediaId)
	return mediaId, client.Send(message)
}

// SendVoiceMessage 发送语音消息 如果mediaId为空，通过给定filePath进行上传并返回mediaId
func (client *Client) SendVoiceMessage(mediaId, filePath string) (string, error) {
	if mediaId == "" {
		if !weworknotify.Exists(filePath) {
			return "", errors.New("file not exists")
		}
		mediaInfo, err := client.UploadFile(VOICE, filePath)
		if err != nil {
			return "", err
		}
		mediaId = mediaInfo.MediaId
	}
	message := NewVoiceMessage(mediaId)
	return mediaId, client.Send(message)
}

// SendTextTemplateMessage 发送文本卡片模板消息 参数详见 https://developer.work.weixin.qq.com/document/path/91770#%E6%96%87%E6%9C%AC%E9%80%9A%E7%9F%A5%E6%A8%A1%E7%89%88%E5%8D%A1%E7%89%87
func (client *Client) SendTextTemplateMessage(source Source, title MainTitle, emphasis Emphasis, quote Quote, subTitle string, content []HorizontalContent, jump []Jump, action CardAction) error {
	message := NewTextTemplateCardMessage(source, title, emphasis, quote, subTitle, content, jump, action)
	return client.Send(message)
}

// SendPictureTemplateMessage 发送图片卡片模板消息 参数详见 https://developer.work.weixin.qq.com/document/path/91770#%E5%9B%BE%E6%96%87%E5%B1%95%E7%A4%BA%E6%A8%A1%E7%89%88%E5%8D%A1%E7%89%87
func (client *Client) SendPictureTemplateMessage(source Source, title MainTitle, image ImgText, quote Quote, vertical []VerticalContent, horizontal []HorizontalContent, jump []Jump, action CardAction) error {
	message := NewPictureTemplateCardMessage(source, title, image, quote, vertical, horizontal, jump, action)
	return client.Send(message)
}
