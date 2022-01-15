package setu

import (
	"context"
	"reflect"
	"testing"
)

func TestGetPic(t *testing.T) {
	tests := []struct {
		name    string
		wantPic Picture
		wantErr bool
	}{
		{
			name:    "",
			wantPic: Picture{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPic, err := GetPic()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPic, tt.wantPic) {
				t.Errorf("GetPic() = %v, want %v", gotPic, tt.wantPic)
			}
		})
	}
}

func TestHandler(t *testing.T) {
	type args struct {
		ctx   context.Context
		event interface{}
	}
	tests := []struct {
		name     string
		args     args
		wantResp interface{}
		wantErr  bool
	}{
		{
			name: "",
			args: args{
				ctx:   nil,
				event: nil,
			},
			wantResp: nil,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := Handler(tt.args.ctx, tt.args.event)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Handler() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
