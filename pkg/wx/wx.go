package wx

import (
	"github.com/monaco-io/request"
)

type wxMessage struct {
	MessageType string `json:"msgtype"`
	Text        text   `json:"text"`
}

type text struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list"`
	MentionedMobileList []string `json:"mentioned_mobile_list"`
}

// WeComWebHookTextMsg 企业微信 使用 webhook 发送text信息
func WeComWebHookTextMsg(webHookURL string, content string, mentionedList []string, mentionedMobileList []string) error {

	weMsgBody := wxMessage{
		MessageType: "text",
		Text: text{
			Content:             content,
			MentionedList:       mentionedList,
			MentionedMobileList: mentionedMobileList,
		},
	}
	c := request.Client{
		URL:    webHookURL,
		Method: "POST",
		JSON:   weMsgBody,
	}
	return c.Send().Error()
}
