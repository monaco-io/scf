package feishu

import (
	"reflect"
	"scf/config"
	"testing"

	"github.com/monaco-io/request"
)

func TestGetRobotGroups(t *testing.T) {
	token, err := GetTenantToken(config.Config.Setu.FsBot.AppID, config.Config.Setu.FsBot.AppSecrect)
	if err != nil {
		t.Fatal(err)
	}
	c := request.Client{
		URL:    getRobotGroups,
		Bearer: token.TenantAccessToken,
	}
	t.Log(c.Send().String())
}

func TestSendSetuMessage(t *testing.T) {
	type args struct {
		chatID   string
		imageKey string
	}
	tests := []struct {
		name     string
		args     args
		wantResp SendMessageResponse
		wantErr  bool
	}{
		{
			name: "",
			args: args{
				chatID:   "oc_56b894cead097418361efaf0a943fcd5",
				imageKey: "img_v2_9c17d3af-4c7e-469e-bcc7-c2aae8abc29g",
			},
			wantResp: SendMessageResponse{},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := SendSetuMessage(tt.args.chatID, tt.args.imageKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendSetuMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("SendSetuMessage() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestSendMessage(t *testing.T) {
	type args struct {
		appID      string
		appSecrect string
		in         SendMessageRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp SendMessageResponse
		wantErr  bool
	}{
		{
			name: "",
			args: args{
				appID:      config.Config.Setu.FsBot.AppID,
				appSecrect: config.Config.Setu.FsBot.AppSecrect,
				in: SendMessageRequest{
					ReceiveIDType: "chat_id",
					Body: SendMessageRequest_Body{
						ReceiveID:    "oc_56b894cead097418361efaf0a943fcd5",
						Content:      `{"text":"test"}`,
						MesssageType: "text",
					},
				},
			},
			wantResp: SendMessageResponse{},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := SendMessage(tt.args.appID, tt.args.appSecrect, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("SendMessage() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
