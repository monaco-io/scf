package feishu

import (
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
