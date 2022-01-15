package feishu

import (
	"scf/config"

	"github.com/monaco-io/request"
)

func GetRobotGroups(appID, appSecrect string) (groups []GroupItem, err error) {
	token, err := GetTenantToken(appID, appSecrect)
	if err != nil {
		return
	}
	for flag := true; flag; {
		var resp GetRobotGroupsResponse
		c := request.Client{
			URL:    getRobotGroups,
			Bearer: token.TenantAccessToken,
		}
		_ = c.Send().Scan(&resp)
		groups = append(groups, resp.Data.Items...)
		flag = resp.Data.HasMore
	}
	return
}

func SearchRobotGroups(appID, appSecrect string) (groups []GroupItem, err error) {
	token, err := GetTenantToken(appID, appSecrect)
	if err != nil {
		return
	}
	for flag := true; flag; {
		var resp GetRobotGroupsResponse
		c := request.Client{
			URL:    searchRobotGroups,
			Bearer: token.TenantAccessToken,
			Query: map[string]string{
				"query": "",
			},
		}
		_ = c.Send().Scan(&resp)
		groups = append(groups, resp.Data.Items...)
		flag = resp.Data.HasMore
	}
	return
}

func GetRobotGroupsSetu() (groups []GroupItem, err error) {
	data, err := GetRobotGroups(config.Config.Setu.FsBot.AppID, config.Config.Setu.FsBot.AppSecrect)
	if err != nil {
		return
	}
	for _, v := range data {
		if v.Name != "" {
			groups = append(groups, v)
		}
	}
	return
}

type GetRobotGroupsResponse struct {
	Code int64                       `json:"code"`
	Data GetRobotGroupsResponse_Data `json:"data"`
}
type GetRobotGroupsResponse_Data struct {
	HasMore bool        `json:"has_more"`
	Items   []GroupItem `json:"items"`
}

type GroupItem struct {
	Avatar      string `json:"avatar"`
	ChatID      string `json:"chat_id"`
	Name        string `json:"name"`
	OwnerID     string `json:"owner_id"`
	OwnerIDType string `json:"owner_id_type"`
	TenantKey   string `json:"tenant_key"`
	External    bool   `json:"external"`
	Description string `json:"description"`
}
