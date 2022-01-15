package setu

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"scf/config"
	"time"

	"github.com/google/uuid"
	"github.com/monaco-io/request"
)

func SaveToTmpFile(pic Picture) (fPath string, err error) {
	tmpDir := config.Config.Setu.CacheDir
	if tmpDir == "" {
		tmpDir = "/tmp"
	}
	fName := pic.Title
	if fName == "" {
		fName = uuid.NewString()
	}
	fPath = fmt.Sprintf("%s/%s.%s", tmpDir, fName, pic.Ext)
	if fileExists(fPath) {
		return
	}
	c := request.Client{
		URL:         pic.URL,
		Timeout:     time.Minute,
		TLSTimeout:  time.Minute,
		DialTimeout: time.Minute,
		Header: map[string]string{
			"Referer": "https://www.pixiv.net",
			// "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.1 Safari/605.1.15",
		},
	}
	resp := c.Send()
	if !resp.OK() {
		err = resp.Error()
		return
	}
	err = ioutil.WriteFile(fPath, resp.Bytes(), fs.ModePerm)
	return
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
