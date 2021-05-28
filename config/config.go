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
}

type bilibliWeekly struct {
	RobotURL string `yaml:"robot_url"`
	Users    []user `yaml:"users"`
}

type user struct {
	Name   string `yaml:"name"`
	Mobile string `yaml:"mobile"`
}

func panicError(err error) {
	if err != nil {
		panic(err)
	}
}
