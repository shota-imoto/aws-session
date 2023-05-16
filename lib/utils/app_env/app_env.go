package app_env

import (
	"errors"
	"os"
)

var LineAppEnv string
var RootPath string = os.Getenv("LINE_APP_ROOT")
var ChannelAccessToken string = os.Getenv("LINE_ACCESS_TOKEN")
var UrlHost string = os.Getenv("URL_HOST")

func init() {
	LineAppEnv = os.Getenv("app_env")
	if LineAppEnv == "" {
		LineAppEnv = "development"
	}

	if RootPath == "" {
		panic(errors.New("env variables is not configed: LINE_APP_ROOT"))
	}
}
