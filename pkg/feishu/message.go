package feishu

import (
	"fmt"
	"scf/config"

	"github.com/monaco-io/request"
	"github.com/pkg/errors"
)

func SendSetuMessage(chatID, imageKey string) (resp SendMessageResponse, err error) {
	content := fmt.Sprintf(`{"image_key":"%s"}`, imageKey)
	externals := config.Config.Setu.FsBot.ExternalGroups
	for _, v := range externals {
		query := map[string]string{
			"msg_type": "image",
			"content":  content,
		}
		c := request.Client{
			URL:   v,
			Query: query,
		}
		_ = c.Send()
	}
	in := SendMessageRequest{
		ReceiveIDType: "chat_id",
		Body: SendMessageRequest_Body{
			ReceiveID:    chatID,
			Content:      content,
			MesssageType: "image",
		},
	}
	return SendMessage(config.Config.Setu.FsBot.AppID, config.Config.Setu.FsBot.AppSecrect, in)
}

func SendMessage(appID, appSecrect string, in SendMessageRequest) (resp SendMessageResponse, err error) {
	token, err := GetTenantToken(appID, appSecrect)
	if err != nil {
		err = errors.Wrap(err, "SendMessage.GetTenantToken")
		return
	}

	c := request.Client{
		URL:    sendMessageURL,
		Method: request.POST,
		Bearer: token.TenantAccessToken,
		Query:  map[string]string{"receive_id_type": in.ReceiveIDType},
		JSON:   in.Body,
	}
	out := c.Send().Scan(&resp)
	err = out.Error()
	return
}

type SendMessageRequest struct {
	// 消息接收者id类型 open_id/user_id/union_id/email/chat_id
	ReceiveIDType string                  `json:"receive_id_type"`
	Body          SendMessageRequest_Body `json:"body"`
}
type SendMessageRequest_Body struct {
	// 依据receive_id_type的值，填写对应的消息接收者id 示例值："ou_7d8a6e6df7621556ce0d21922b676706ccs"
	ReceiveID string `json:"receive_id"`
	// 消息内容，json结构序列化后的字符串。
	// 不同msg_type对应不同内容。
	// 消息类型 包括：text、post、image、file、audio、media、sticker、interactive、share_chat、share_user等，
	// 具体格式说明参考：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json
	Content string `json:"content"`
	// 消息类型 包括：text、post、image、file、audio、media、sticker、interactive、share_chat、share_user等，类型定义请参考
	MesssageType string `json:"msg_type"`
}
type SendMessageResponse struct {
	Code    int64       `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}
