package feishu

import (
	"reflect"
	"testing"
)

func TestGetTenantToken(t *testing.T) {
	type args struct {
		appID      string
		appSecrect string
	}
	tests := []struct {
		name     string
		args     args
		wantResp TenantResponse
		wantErr  bool
	}{{
		name: "",
		args: args{
			appID:      "cli_9e780088f1f7100d",
			appSecrect: "T0FjsUFgVdG4gFYEULou1dnVnwQXyL0A",
		},
		wantResp: TenantResponse{
			Code:              0,
			Msg:               "",
			TenantAccessToken: "",
			Expire:            0,
		},
		wantErr: false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := GetTenantToken(tt.args.appID, tt.args.appSecrect)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTenantToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("GetTenantToken() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
