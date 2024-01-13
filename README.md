# wecom-notify-go-sdk
> Golang 企业微信群bot推送通知 SDK

## 支持的功能
- [x] [文本](https://developer.work.weixin.qq.com/document/path/91770#%E6%96%87%E6%9C%AC%E7%B1%BB%E5%9E%8B)
- [x] [markdown](https://developer.work.weixin.qq.com/document/path/91770#markdown%E7%B1%BB%E5%9E%8B)
- [x] [图片](https://developer.work.weixin.qq.com/document/path/91770#%E5%9B%BE%E7%89%87%E7%B1%BB%E5%9E%8B)
- [x] [图文](https://developer.work.weixin.qq.com/document/path/91770#%E5%9B%BE%E7%89%87%E7%B1%BB%E5%9E%8B)
- [x] [文件](https://developer.work.weixin.qq.com/document/path/91770#%E6%96%87%E4%BB%B6%E7%B1%BB%E5%9E%8B)
- [x] [语音](https://developer.work.weixin.qq.com/document/path/91770#%E8%AF%AD%E9%9F%B3%E7%B1%BB%E5%9E%8B)
- [x] [文本通知模板卡片](https://developer.work.weixin.qq.com/document/path/91770#%E6%96%87%E6%9C%AC%E9%80%9A%E7%9F%A5%E6%A8%A1%E7%89%88%E5%8D%A1%E7%89%87)
- [x] [图片通知模板卡片](https://developer.work.weixin.qq.com/document/path/91770#%E5%9B%BE%E6%96%87%E5%B1%95%E7%A4%BA%E6%A8%A1%E7%89%88%E5%8D%A1%E7%89%87)
- [x] [文件上传](https://developer.work.weixin.qq.com/document/path/91770#%E6%96%87%E4%BB%B6%E4%B8%8A%E4%BC%A0%E6%8E%A5%E5%8F%A3)

## 安装使用

### 环境要求
- go 1.18+

### 安装
```shell
go get -u github.com/xinyunaha/wecom-notify-go-sdk
```
### 引入包
```go
import (
    "github.com/xinyunaha/wecom-notify-go-sdk"
)
```

### 发送文本消息
```go
notify := wecomBotNotify.NewWecomBotClient("YOUR_TOKEN")
err := notify.SendTextMessage("Test Message", []string{wecomBotNotify.AtAll}, nil)
if err != nil {
    // error
}
```

### 发送markdown消息
```go
notify := wecomBotNotify.NewWecomClient("YOUR_TOKEN")
err := notify.SendMarkdownMessage("这是一条Markdown测试消息\n<font color=\"comment\">测试</font>")
if err != nil {
    // error
}
```

### 发送图片消息
```go
    notify := wecomBotNotify.NewWecomClient("YOUR_TOKEN")
    err := notify.SendImageMessage("IMAGE_BASE64", "IMAGE_MD5")
    if err != nil {
        // error
    }
```

###  发送图文消息
```go
notify := wecomBotNotify.NewWecomClient("YOUR_TOKEN")
notify.SendNewsMessage([]Articles{
    {
        Title:       "图文消息测试",
        Description: "这是一条测试消息",
        URL:         "https://www.baidu.com",
        PicUrl:      "https://img2.baidu.com/it/u=3992244728,3167300513&fm=253&fmt=auto&app=138&f=PNG?w=1068&h=455",
    }, {
        Title:       "图文消息测试",
        Description: "这是一条测试消息",
        URL:         "https://www.baidu.com",
        PicUrl:      "https://img2.baidu.com/it/u=3992244728,3167300513&fm=253&fmt=auto&app=138&f=PNG?w=1068&h=455",
    }, {
        Title:       "图文消息测试",
        Description: "这是一条测试消息",
        URL:         "https://www.baidu.com",
        PicUrl:      "https://img2.baidu.com/it/u=3992244728,3167300513&fm=253&fmt=auto&app=138&f=PNG?w=1068&h=455",
    },
}
```

### 发送语音消息
> 语音消息仅能上传amr格式文件
- 已知media_id
```go
notify := wecomBotNotify.NewWecomClient("YOUR_TOKEN")
// 给定media_id发送语音，media_id可通过上传文件获取
mediaID, err := notify.SendVoiceMessage("MEDIA_ID", "")
if err != nil {
    // error
}
```
- 未知media_id，上传文件并发送
```go
notify := wecomBotNotify.NewWecomClient("YOUR_TOKEN")
// media_id不填，给定path将自动上传并返回media_id
mediaID, err := notify.SendVoiceMessage("", "VOICE_PATH")
if err != nil {
    // error
}
```

### 发送文件
- 已知media_id
```go
notify := wecomBotNotify.NewWecomClient("YOUR_TOKEN")
// 给定media_id发送文件，media_id可通过上传文件获取
mediaID, err := notify.SendFileMessage("MEDIA_ID", "")
if err != nil {
    // error
}
```
- 未知media_id，上传文件并发送
```go
notify := wecomBotNotify.NewWecomClient("YOUR_TOKEN")
// media_id不填，给定path将自动上传并返回media_id
mediaID, err := notify.SendFileMessage("", "FILE_PATH")
if err != nil {
    // error
}
```

### 文字卡片模板消息
```go
notify := wecomBotNotify.NewWecomClient("YOUR_TOKEN")
err := notify.SendTextTemplateMessage(
    Source{
        IconUrl:   "https://wecom.qpic.cn/wwpic/252813_jOfDHtcISzuodLa_1629280209/0",
        Desc:      "企业微信通知",
        DescColor: 0,
    },
    MainTitle{
        Title: "欢迎使用企业微信",
        Desc:  "这是一条测试消息",
    },
    Emphasis{
        Title: "99.98%",
        Desc:  "24h可用率",
    },
    Quote{
        Type:      1,
        Url:       "https://work.weixin.qq.com/?from=openApi",
        Appid:     "APPID",
        PagePath:  "PAGEPATH",
        Title:     "引用标题测试",
        QuoteText: "引用文本测试\nbot引用文本测试",
    },
    "子标题测试",
    []HorizontalContent{
        {
            KeyName: "测试人",
            Value:   "bot",
        },
        {
            KeyName: "企微官网",
            Value:   "点击访问",
            Type:    1,
            Url:     "https://work.weixin.qq.com/?from=openApi",
        },
    },
    []Jump{
        {
            Type:  1,
            Url:   "https://work.weixin.qq.com/?from=openApi",
            Title: "企业微信官网",
        },
    },
    CardAction{
        Type:     1,
        Url:      "https://work.weixin.qq.com/?from=openApi",
        Appid:    "APPID",
        PagePath: "PAGEPATH",
    },
)
if err != nil {
    // error
}
```

### 图片卡片模板消息
```go
notify := wecomBotNotify.NewWecomClient("YOUR_TOKEN")
err := notify.SendPictureTemplateMessage(
    Source{
        IconUrl:   "https://wecom.qpic.cn/wwpic/252813_jOfDHtcISzuodLa_1629280209/0",
        Desc:      "企业微信通知",
        DescColor: 0,
    },
    MainTitle{
        Title: "欢迎使用企业微信",
        Desc:  "这是一条测试消息",
    },
    ImgText{
        Type:     1,
        Url:      "https://work.weixin.qq.com",
        Title:    "欢迎使用企业微信",
        Desc:     "这是一条测试消息",
        ImageUrl: "https://img2.baidu.com/it/u=3992244728,3167300513&fm=253&fmt=auto&app=138&f=PNG?w=1068&h=455",
    },
    Quote{
        Type:      1,
        Url:       "https://work.weixin.qq.com/?from=openApi",
        Title:     "引用标题测试",
        QuoteText: "引用文本测试\nbot引用文本测试",
    },
    []VerticalContent{
        {
            Title: "垂直区域二级标题",
            Desc:  "测试消息",
        },
    },
    []HorizontalContent{
        {
            KeyName: "测试人",
            Value:   "bot",
        },
        {
            KeyName: "企微官网",
            Value:   "点击访问",
            Type:    1,
            Url:     "https://work.weixin.qq.com/?from=openApi",
        },
    },
    []Jump{
        {
            Type:  1,
            Url:   "https://work.weixin.qq.com/?from=openApi",
            Title: "企业微信官网",
        },
    },
    CardAction{
        Type:     1,
        Url:      "https://work.weixin.qq.com/?from=openApi",
        Appid:    "APPID",
        PagePath: "PAGEPATH",
    },
)
if err != nil {
    // error
}
```

### 上传文件
```go
notify := wecomBotNotify.NewWecomClient("YOUR_TOKEN")
// 给定文件类型( FILE | VOICE)及文件路径进行上传，返回带有media_id的结构体
mediaInfo, err := notify.UploadFile(FILE, filePath)
if err != nil {
	log.Errorln(mediaInfo.ErrCode,mediaInfo.ErrMsg)
}
log.Info(mediaInfo.MediaId)
```