package feishu

import "github.com/monaco-io/request"

func GetTenantToken(appID, appSecrect string) (resp TenantResponse, err error) {
	jsonBody := TenantRequest{
		AppID:      appID,
		AppSecrect: appSecrect,
	}
	c := request.Client{
		URL:    getInternalTokenURL,
		Method: request.POST,
		JSON:   jsonBody,
	}
	out := c.Send().Scan(&resp)
	err = out.Error()
	return
}

type TenantRequest struct {
	AppID      string `json:"app_id"`
	AppSecrect string `json:"app_secret"`
}

type TenantResponse struct {
	Code              int64  `json:"code"`
	Msg               string `json:"msg"`
	TenantAccessToken string `json:"tenant_access_token"`
	Expire            int64  `json:"expire"`
}
