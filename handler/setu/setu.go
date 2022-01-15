package setu

// doc: https://api.lolicon.app
import (
	"context"
	"fmt"
	"log"
	"os"
	"scf/config"
	"scf/pkg/feishu"

	"github.com/monaco-io/request"
	"github.com/pkg/errors"
)

func Handler(ctx context.Context, event interface{}) (resp interface{}, err error) {
	pic, err := GetPic()
	if err != nil {
		return
	}
	picPath, err := SaveToTmpFile(pic)
	log.Printf("INFO download pic %s %v", picPath, err)
	if config.Config.Setu.CacheDir == "" {
		defer os.Remove(picPath)
	}
	if err != nil {
		err = errors.Wrap(err, "SaveToTmpFile(pic)")
		return
	}
	seResp, err := feishu.UploadSetuPic(picPath)
	log.Printf("INFO upload pic to feishu %+v %v", seResp, err)
	if err != nil {
		err = errors.Wrap(err, "feishu.UploadSetuPic(picPath)")
		return
	}
	groups, err := feishu.GetRobotGroupsSetu()
	if err != nil {
		err = errors.Wrap(err, "feishu.GetRobotGroupsSetu()")
		return
	}
	for _, g := range groups {
		msgResp, err2 := feishu.SendSetuMessage(g.ChatID, seResp.Data.ImageKey)
		log.Printf("INFO send pic to feishu grup %+v %v", msgResp, err)
		if err != nil {
			err = errors.Wrap(err2, "feishu.SendSetuMessage(seResp.Data.ImageKey)")
			return
		}
	}

	resp = event
	return
}

func GetPic() (pic Picture, err error) {
	var (
		req  Request
		resp Response
	)

	req.R18 = config.Config.Setu.R18
	req.Size = AllSize
	req.Proxy = config.Config.Setu.Proxy

	c := request.Client{
		URL:    SeTuURL,
		Method: request.POST,
		JSON:   req,
	}

	out := c.Send().Scan(&resp)
	if !out.OK() {
		log.Println(out.Error())
		return
	}
	log.Printf("INFO resp: %+v", resp)

	if len(resp.Data) > 0 {
		pic.Title = fmt.Sprintf("%s-%s-%d", resp.Data[0].Author, resp.Data[0].Title, resp.Data[0].PID)
		pic.Tags = resp.Data[0].Tags
		pic.URL = resp.Data[0].URLs.Original
		pic.Ext = resp.Data[0].Ext
	} else {
		err = errors.New(resp.Error)
	}
	return
}
