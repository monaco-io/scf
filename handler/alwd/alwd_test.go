package alwd

import (
	"context"
	"testing"
)

func TestWords_GetOne(t *testing.T) {
	tests := []struct {
		name   string
		fields *Words
		want   string
	}{
		{
			fields: N2(),
			want:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Words{
				Data:  tt.fields.Data,
				Bloom: tt.fields.Bloom,
			}
			if got := w.GetOne(); got != tt.want {
				t.Errorf("Words.GetOne() = %v, want %v", got, tt.want)
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
		wantResp string
		wantErr  bool
	}{
		{
			args: args{
				ctx:   context.TODO(),
				event: nil,
			},
			wantResp: "",
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
			if gotResp != tt.wantResp {
				t.Errorf("Handler() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
