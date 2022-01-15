package config

import (
	"embed"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

var (
	//go:embed conf/*
	f embed.FS

	Config *_Config
)

func init() {
	Init()
}

func env() string {
	env := os.Getenv("APP_ENV")
	if env == "" {
		return "dev"
	}
	return env
}

func Init() {
	var err error

	conf, err := f.ReadFile(path.Join("conf", env()+".yaml"))
	panicError(err)

	err = yaml.Unmarshal(conf, &Config)
	panicError(err)
}

type _Config struct {
	BilibiliWeekly bilibliWeekly `yaml:"bilibili_weekly"`
	Setu           setu          `yaml:"setu"`
}

type bilibliWeekly struct {
	RobotURL string `yaml:"robot_url"`
	Users    []user `yaml:"users"`
}

type setu struct {
	CacheDir string `yaml:"cache_dir"`
	R18      int64  `yaml:"r18"`
	Proxy    string `yaml:"proxy"`
	URL      string `yaml:"url"`
	FsBot    fsBot  `yaml:"fs_bot"`
}
type user struct {
	Name   string `yaml:"name"`
	Mobile string `yaml:"mobile"`
}

type fsBot struct {
	URL        string `yaml:"url"`
	AppID      string `yaml:"app_id"`
	AppSecrect string `yaml:"app_secrect"`
}

func panicError(err error) {
	if err != nil {
		panic(err)
	}
}
