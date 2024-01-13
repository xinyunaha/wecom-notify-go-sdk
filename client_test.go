/**
 * @Author: xinyunaha
 * @File:  client_test
 * @Date: 2024/1/12 9:09
 * @Version: 1.0.0
 * @Description:
 */

package wework_notify

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

var (
	filePath  = "F:\\project\\wework-notify\\test\\test.txt"
	voicePath = "F:\\project\\wework-notify\\test\\test.amr"
	notify    = NewWecomBotClient("YOUR_WEWORK_BOT_KEY")
)

func TestClient_SendTextMessage(t *testing.T) {
	assert.Equal(t, notify.SendTextMessage("这是一条消息测试", nil, nil), nil)
}

func TestClient_SendImageMessage(t *testing.T) {
	assert.Equal(t, notify.SendImageMessage("iVBORw0KGgoAAAANSUhEUgAAAMgAAADICAYAAACtWK6eAAAAAXNSR0IArs4c6QAAEA9JREFUeF7tXQ2y3rYN/HIStydxexK3J0lyksYnaX2S1idpjBfR0ZMlEViABCmuZjwvk48UyQWWWP7qpxcfIkAELhH4idgQASJwjQAJQu8gAjcIkCB0DyJAgtAHiACGACMIhhtzLYIACbKIodlMDAESBMONuRZBgARZxNBsJoYACYLhxlyLIECCLGJoNhNDgATBcGOuRRAgQfIN/ZfX61X+fXi9Xl+3Kv3v9XqVf/m1XLQGJEh/wwsZPn1z/n9sxNDUQIjyn9fr9Xn7q8nDNAEIkCABICpeUUjxiyJtLQnJUkMo8HcSJBDMk1dJlPjZECmstRGy/Prt/b9ZMzK9DgESRIeTJVVktNCWK0QRkogEk//mE4QACRIE5BYlZGwRIaPQWlF+ochd5CNB/IC2llFoDSm/UOR2+UgQDMQMGYXV9A/JRfkFokeC2ICbiRjHllF+2Wz9lpoE0YH2t202Sv4+4aH8UlqRBLkGClnQU8I+TDLKr4opSJAfAZpZRqHMo/ziLFbVd0Q+lS0g1cQPTkD5xVms7whky6gicb5sNZL9VlIneeTvR+OerUjelrrJSv2yz6oSK1tGWXvpzOi2tPxajSDZjuZdj5iN2NNHnhUIMoqMipQq0qZC9oyp52Xk15MJskpvmx0VH31O5YkEEWLIFnPZI9X7yVxXWKVD6GrTpxDkiTIKdQTKLxS5k3yzE4S95r0zUH45yTIrQTJllEAuZz5mOpzEjgQkykwEoYwCjbzLRvllxHAGgrD3MxpVmZzySwHUyAShjFIYMCBJdgdUpomHvHhiNIJQRgV4PPgKYj/wLFZ2L2bdGwX64DTZZA1JdjZnrdIPs/iYHUEoo8bmTHbHlS6/MgjCUD42Kc5qt6zNehIkuzeijIohZrb88u6INqHQgyDZxJhtUc9kwMTEmXbttpu4NUFkkPfvBCN2AzChbaMVmSm/mtu5JUFkR23vazgpo3LpkyW/xM8iz9t8R7EFQaRH+VfnKULKqFxiHEvPkF/SOf49+vLuaIL0JEfz8DqWz01Zm97yK5wkkQTpRQ7KqCm58naArcfiYyhJIgkig/GWK6+UUXMSI0N+yQKjyC33E0WQVgNyyii3iYd9QWv5FTJwjyBIi6lcyqhh/Tq8Yi3PqEgUkWgCPxEE+W/gN/goo2BTPiJj9OyXdLR/9SDjJYgMvGRK1/NQRnnQe2beSPn1T89HTr0E+b/TPiE60VkHZh8XgYiI4ooiHoJ4okfoVNy49mXNghDwLiHAHbGHIOjYI2wKLgh8vmYOBDwkgaMIShB05gqu6Bw2ZC0bIyAkkfU2+Wt9oBktlCAyMEeu9oQqaUWC6R+NgJBD1Iv1gWQWShBk1VwOusiMAh8i4EUA6aAhaY8SBJm9kvlokVh8iIAXASSKQPIeIQgy/mD08LoE8x8RQFSMWeIjBEGmd12LNfQNInCCAOKHJAhdaRkEECVj7qh7RRCOP5bx224NRcYh5pkshCDI1naknG5Is6BpEbBOFpEg05qaFUcQIEEQ1JhnGQRIkGVMzYYiCJAgCGrMswwCJMgypmZDEQRIEAQ15lkGARJkGVOzoQgCJAiCGvMsgwAJsoyp2VAEARIEQY15lkGABFnG1GwoggAJgqC2WB7Z1Sob9z5s7d6f15ZDQl+3g2quGwYHxXQZgsgGSDEschhfa7tymlE+tFI72VjuY5J3t7iQu1ytWqvHVduQ+6KkLM/nlUuZLfDYt1Pq+UV52dvjCYLs6dcS4ird3b2/nqtlrPVC7h+OwktOeGo6itImZAe3FY9jes0R2ccTBDk26QVe8l9dXNe7PhonaOWkWoJGkRKxW+2o9qMJghx4QUC+ynM8XYYc4YyoT+0YqOcuqFr9NPciIzeI1MrV/l7rQB5NkCyHLMY5Hp7Jqk/tGGjrqCZOKHW4GsSjN2hqSVBLd3ca9dEEyQzdYpQjQbJ6yjuC9NL+d3cltyZojSB3p1EfTZBsiXWUNhkEuZMQvTuQq4vVepH0jCi1y94eTRABJAv8M+B795Q1aWOpT/kevUyPynvLlPlH4xWxZ+OhlmOgWvSojc8eTxBkPr8Gau33q4P7Wof0LriJA5cZpLu6aowvdRGJdreeYsH4qsfuOf1dZhnvxkUFNw1Ge4ynvbShLBDeLRRKtNEsUt193fTOubUE6XFDi0buWYytdfDarJFmMVdwrD0yffv5IpGlA1qGIDVA5XeN00g61IFHIkht5qimzc/w1EolFD8pUzuutJDbG2UfEUFIkPcI1HrG2tTwFZ6aqeya7r+zFQkCDqw9vVIxyEoRpEYQ9KZKTRTx9O4kCAkCSzhNlJQ0tend2jihVk4tiqDRiRJrQx6ZmmUEqbntn7/XCCIp0QhSnPjuM2bed2u+/uSJUnska5H2iLq5XMRxSRC9syMpNQTxjBOkTldRpLY5sNYeSixKrOYSS+NkEY4sY7r9tLlXulFiUWK9IYBE3VrPe/xdIx0iSCIEkdV2zcKlpg0acst7zFLnonANTpzmPYCHOvBM6yClyZqVdI1jR6UhQSixukQQyzivnOmQfViWVegoUuzfQ4KQIF0IonW0o5Nnk0Vbb0qsINmzf81KC4XSbq3ku4oEhSyy5wm9HMIaZUgQRhCrz3xPX3bxigySCxNqT21Br5a//F7KFaLIwL7lQ4KQICH+pZ2BsoxFNBXzXv1TK4MEIUFqPqL+XbPYp92mri50l9B69Y+mDBKEBNH4iSqNdqCq2WCoKvAkkeZWE8u7SRASxOIvt2m1BJGXWE4FIhVEzpeclUOCkCCI/53mQXbNysD9k/JkpbWi2nHR3XtJEBLE6nen6b37nloRxUsSEoQEcROkdpuJpQDZSyURRQgT9WgmD67KIkFIkLdbRNCnTLOi+e8cs2w89JLFMx4hQUiQLltNPAQSJ/WSBY0iJAgJMjxB9uQqZLEO7C0zbMfyeKLwGyLlJj9tT4duQd+/f7W9WFpsteksW1fQSQRGEEaQrhGkXNQWtY1d68BCOuRsuvb9aIQ6dgY8MLVDZKUIcjyX7hk4H51KG0kQgkhZGqclQQ5WocTSCqHra3+8axSlBtqtK+hAvXYrpNQDWSQ9Q1BDxn0+MzERx0V2mCLlHAFZJYLcnQNBnfaIpcaGaFmacyxodKLEuuloVyHIXa8YFUU0BEF7+dq1RZFykRFksTFIbZCLzi4h4xA0gtS26KPvpcSqyPQVIkiNIFH6XRNBPDLoavdxJDm0EwIcgwRNCGi0sxQVMZ666wtqg9yIKKLpbCLaWRYqW22xocRaTGJJczVE9fbENRL26Aj083rXKUmQBQmikT+ewXrr90c4vvYdJMiCBNGMQwQW85y+4tMKBW50Bkvr2FHpSJABCSKDV+9Tu6dKI7Os58k1kaO0K2L8Ud4lhK+1F8WTBBmQIKgxj/nuHLy2lrB/V/kg5tV+rdq067FeHvm2f9eekIUgmi/XWvAlQR5MkJqc0USRvTOVC+KKM5YZJIvDSVrvBIC8467uiDS8agMJ0pEgmqlPq7Np0l9N22rHIpoytGkiyFGTcpFHjkmQnWVrwHv1s/b9WmezpLtalNPuvLWUdZU2aguIZgo5KoqQIEaCeIzc0xmPTnq3at2DuBGLj6VNGqf12GmPnaasfXozMZHZCsRgSDlHJ9IMXD0DzAxJo416CObaqBLlrFKeFkOPnUiQC8tqzjF45+8zxiFaZxGSSJQTHKIec4+qKFgzuRBVLiPIwSB3MihCJrS+4vPoX9Y6l/p5iSIO2uqLVDWpam3zHSdJkBN0zuRGJOhFKoik+xDcY5fmSH2/Or7XgdxWYl1YVASLyyRXkjByBksKJ0Fu5FZx4FY9ocdBeuctsqt81lmILY9gI0/UpQ+WdkmdhCilbtoPB1nKIEEsaDHtcgiQIMuZnA22IECCWNBi2uUQeAxBIrYwLGd9NvgWAe2ay/4l5mUAZAGvNo131ioShN4ejQDih10IolnRPoIRtTAUDTLfNy8CCEHMHTUSQZDQFrmdYV6TsuaRCGhW7I/lmW9qQQgihWp2bB4rZ2ZvJJp816MQQFSMAGD2d3OGDWaEvYwij/LR1MYg/qfd8/auYShB0N2ljCKpfvWIwtHoYR6gQyFngxgZh0jW6L04j7A4G6FGAPU7KcA8/vAQRPIiYa6QRCJJq5su1Ggz4VQIeMgBySsvQdBQV0gilf51KhOxslkIeHxN6gzJKy9BrNfJnIErUYTRJMvtxi83ysfgu8zQQXqB1svsEk1khutz0rbs8d1kvRqWrfLeb74Lcq6JIS9BPGORq4hC6bUeIUqLo48Vu5cWIgjiGTxduUI5/SZRhYP5ZxOm5TFnV/TwjkH2ZkP2xWjMXr4rQfmlQWuuNC2J4ZZWBcqICLIPj7IpsdUjZJFZL5FgfOZFIFpGnSEBT+seXxZJkNY9Qqk75dd85OjlG4KMe9yxhzeSIPLenkBQfo1PlJ7+EE6OyDHI3lS9QZGyKb/GIouMSfc3mvSoXWjkaDEGOYKAbmj0gNnz7idPPZ+YN+pSOwQbeKW8Vli0xDqWF7ESWmvD2e+UXwhqWJ4MxbAfjzbdidGaIKUhEnI/bd/Iw8yA56L8wrG7yxm52m2tYTeb9iJIGcDL1pRMonCV3uqK79NnyqiUsWZPghwH8mU+3Gcye27KLztmmTJKaivraym7KrIIUkyUDXzZJMnFx3PSZMuo9IifTZA9USi/7D17qxw9Vruv6j7UTOQoBBlJfpXPKq+2STI7mpdoMRTuIxJkFPlV+/54q96793tHkFEp4wsN0CMThPJLY0E8DWWUArsZCLJvRhmnRJw0U8DzLskTNklmy6hu6xdW416ln40gI8ivGaeJM4kxdccyK0GORPF+1BLtcIaacTlpBGUUatkt3+wE2Tc/ezvLKLNfmdFC7DGdjLrj0JMIsrr8yiZG2mq3M0jcZn8iQUaSXz2OCGecvSgYjy4x3dx5MkFGmyaOnOvPjhaPklGrSay79mY6VsTsV2b9BddHyigS5EcEZtu2PcJq95L3KK8gsWoRJXuTpOwo/rLN/uz3IQkpPm69tltLAy94/PhCg8nqBNljlNlLa2zVK81yMooSy+Za2fLLVtuY1IwWFzgyglw7mBAlU37FuP79W0iMCsokiM4Nnya/ZNwjV+UMdfZCZ4q+qUgQG97Z06y22r5PPfWmQU/DPXlJEAy9meQXZRRm47dcJIgDvC3rqPJrmdVuvwmv30CCxKE7gvyijIqzJyNIMJbldRnyizKqkTEZQRoBu7229RFhyqi29uMYpDG++6hSrlwV0ngeyigPesa8jCBGwAKSl7GKEEVLFpIiAHjkFSQIglpsHiGMPPJX/n3YNi/K/5MFPT6JCJAgieCz6PERIEHGtxFrmIgACZIIPoseHwESZHwbsYaJCJAgieCz6PERIEHGtxFrmIgACZIIPoseHwESZHwbsYaJCJAgieCz6PERIEHGtxFrmIgACZIIPoseHwESZHwbsYaJCPwOj1+9BaSBR6MAAAAASUVORK5CYII=", "0ea576aed8e2c0ef8121dc205f511cfe"), nil)
}

func TestClient_SendMarkdownMessage(t *testing.T) {
	assert.Equal(t, notify.SendMarkdownMessage("这是一条Markdown测试消息\n<font color=\"comment\">测试</font>"), nil)
}

func TestClient_SendNewsMessage(t *testing.T) {
	assert.Equal(t, notify.SendNewsMessage([]Articles{
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
	}), nil)
}

func TestClient_UploadFile(t *testing.T) {
	info, err := notify.UploadFile(FILE, filePath)
	assert.Equal(t, err, nil)
	assert.Equal(t, info.ErrCode, 0)
}

func TestClient_SendVoiceMessage(t *testing.T) {
	mediaID, errFromUploadFile := notify.SendVoiceMessage("", voicePath)
	assert.Equal(t, errFromUploadFile, nil)
	_, errFromMediaID := notify.SendVoiceMessage(mediaID, "")
	assert.Equal(t, errFromMediaID, nil)
}

func TestClient_SendFileMessage(t *testing.T) {
	mediaID, errFromUploadFile := notify.SendFileMessage("", filePath)
	assert.Equal(t, errFromUploadFile, nil)
	_, errFromMediaID := notify.SendFileMessage(mediaID, "")
	assert.Equal(t, errFromMediaID, nil)
}

func TestClient_SendTextTemplateMessage(t *testing.T) {
	err := notify.SendTextTemplateMessage(
		Source{
			IconUrl:   "https://wework.qpic.cn/wwpic/252813_jOfDHtcISzuodLa_1629280209/0",
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
	assert.Equal(t, err, nil)
}

func TestClient_SendPictureTemplateMessage(t *testing.T) {
	err := notify.SendPictureTemplateMessage(
		Source{
			IconUrl:   "https://wework.qpic.cn/wwpic/252813_jOfDHtcISzuodLa_1629280209/0",
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
	assert.Equal(t, err, nil)
}
