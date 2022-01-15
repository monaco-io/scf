package feishu

import (
	"reflect"
	"scf/config"
	"testing"
)

func TestGetRobotGroupsSetu(t *testing.T) {
	tests := []struct {
		name       string
		wantGroups GetRobotGroupsResponse
		wantErr    bool
	}{
		{
			name:       "",
			wantGroups: GetRobotGroupsResponse{},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGroups, err := GetRobotGroupsSetu()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRobotGroupsSetu() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, v := range gotGroups {
				t.Log(v.ChatID, v.Name)
			}
		})
	}
}

func TestSearchRobotGroups(t *testing.T) {
	type args struct {
		appID      string
		appSecrect string
	}
	tests := []struct {
		name       string
		args       args
		wantGroups []GroupItem
		wantErr    bool
	}{
		{
			name: "",
			args: args{
				appID:      config.Config.Setu.FsBot.AppID,
				appSecrect: config.Config.Setu.FsBot.AppSecrect,
			},
			wantGroups: []GroupItem{},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGroups, err := SearchRobotGroups(tt.args.appID, tt.args.appSecrect)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchRobotGroups() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotGroups, tt.wantGroups) {
				t.Errorf("SearchRobotGroups() = %v, want %v", gotGroups, tt.wantGroups)
			}
		})
	}
}
