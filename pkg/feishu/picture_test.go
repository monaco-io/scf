package feishu

import (
	"reflect"
	"testing"
)

func TestUploadSetuPic(t *testing.T) {
	type args struct {
		piPathc string
	}
	tests := []struct {
		name     string
		args     args
		wantResp UploadPicResponse
		wantErr  bool
	}{
		{
			name: "",
			args: args{
				piPathc: "/Users/monaco/Desktop/1635468351852.jpeg",
			},
			wantResp: UploadPicResponse{},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := UploadSetuPic(tt.args.piPathc)
			if (err != nil) != tt.wantErr {
				t.Errorf("UploadPic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("UploadPic() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

// img_v2_9c17d3af-4c7e-469e-bcc7-c2aae8abc29g
