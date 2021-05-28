package handler

import (
	"context"
	"reflect"
	"testing"
)

func TestBilibiliWeeklyRemind(t *testing.T) {
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
			gotResp, err := BilibiliWeeklyRemind(tt.args.ctx, tt.args.event)
			if (err != nil) != tt.wantErr {
				t.Errorf("BilibiliWeeklyRemind() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("BilibiliWeeklyRemind() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
