package feishu

import (
	"io/ioutil"
	"scf/config"

	"github.com/monaco-io/request"
	"github.com/pkg/errors"
)

func UploadSetuPic(picPath string) (resp UploadPicResponse, err error) {
	return UploadPic(config.Config.Setu.FsBot.AppID, config.Config.Setu.FsBot.AppSecrect, picPath)
}

func UploadPic(appID, appSecrect, picPath string) (resp UploadPicResponse, err error) {
	token, err := GetTenantToken(appID, appSecrect)
	if err != nil {
		err = errors.Wrap(err, "UploadPic.GetTenantToken")
		return
	}
	f, err := ioutil.ReadFile(picPath)
	if err != nil {
		err = errors.Wrap(err, "readfile from picPath")
		return
	}
	form := request.MultipartForm{
		Fields: map[string]string{
			"image_type": "message",
			"image":      string(f),
		},
	}
	c := request.Client{
		URL:           uploadImageURL,
		Method:        request.POST,
		Bearer:        token.TenantAccessToken,
		MultipartForm: form,
	}
	out := c.Send().Scan(&resp)
	err = out.Error()
	return
}

type UploadPicResponse struct {
	Code int64                  `json:"code"`
	Data UploadPicResponse_Data `json:"data"`
	Msg  string                 `json:"msg"`
}
type UploadPicResponse_Data struct {
	ImageKey string `json:"image_key"`
}
