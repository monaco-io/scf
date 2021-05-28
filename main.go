package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"scf/config"
	"scf/handler"
	"time"

	"github.com/tencentyun/scf-go-lib/cloudfunction"
)

type EventTimer struct {
	Type        string
	TriggerName string
	Time        time.Time
	Message     string
}

type EventHttp struct {
	HTTPMethod            string         `json:"httpMethod"`
	Path                  string         `json:"path"`
	Headers               Headers        `json:"headers"`
	PathParameters        interface{}    `json:"pathParameters"`
	QueryString           interface{}    `json:"queryString"`
	QueryStringParameters interface{}    `json:"queryStringParameters"`
	HeaderParameters      interface{}    `json:"headerParameters"`
	RequestContext        RequestContext `json:"requestContext"`
	IsBase64Encoded       bool           `json:"isBase64Encoded"`
}

type Headers struct {
	Accept        string `json:"accept"`
	Host          string `json:"host"`
	RequestSource string `json:"requestsource"`
	UserAgent     string `json:"user-agent"`
	XAPIRequestID string `json:"x-api-requestid"`
	XAPIScheme    string `json:"x-api-scheme"`
	XB3TraceId    string `json:"x-b3-traceid"`
	XQualifier    string `json:"x-qualifier"`
}

type RequestContext struct {
	HTTPMethod string      `json:"httpMethod"`
	Identity   interface{} `json:"identity"`
	Path       string      `json:"path"`
	ServiceId  string      `json:"serviceId"`
	SourceIp   string      `json:"sourceIp"`
	Stage      string      `json:"stage"`
}

func _handler(ctx context.Context, event interface{}) (resp interface{}, err error) {
	os.Setenv("APP_ENV", "prod")
	config.Init()

	var (
		eHttp   EventHttp
		eTimer  EventTimer
		success = func() bool {
			return err == nil
		}
	)

	j, _ := json.Marshal(event)

	err = json.Unmarshal(j, &eHttp)
	if success() && eHttp.HTTPMethod != "" {
		log.Println("Method", eHttp.HTTPMethod, "Path", eHttp.Path)
		switch eHttp.Path {
		case "/gogo/bilibili_weekly_remind":
			return handler.BilibiliWeeklyRemind(ctx, event)
		case "/gogo/http":
			return handler.HTTPSource(ctx, event)
		case "/gogo/alwd":
			return handler.ALWD(ctx, event)
		default:
			log.Printf("unknown method: %+v", eHttp)
			return
		}
	}

	err = json.Unmarshal(j, &eTimer)
	if success() && eTimer.TriggerName != "" {
		log.Println("TriggerName", eTimer.TriggerName, "Type", eTimer.Type)
		switch eTimer.TriggerName {
		case "bilibili_weekly_remind":
			return handler.BilibiliWeeklyRemind(ctx, event)
		default:
			log.Printf("unknown timer: %+v", eTimer)
			return
		}
	}

	log.Printf("unknown event: %+v", event)
	return
}

func main() {
	// Make the handler available for Remote Procedure Call by Cloud Function
	cloudfunction.Start(_handler)
}
