package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"os"
	"scf/config"
	"scf/handler/alwd"
	"scf/handler/biliweekly"
	xhttp "scf/handler/http"
	"scf/handler/setu"
	"time"

	"github.com/robfig/cron/v3"
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
			return biliweekly.Handler(ctx, event)
		case "/gogo/http":
			return xhttp.Handler(ctx, event)
		case "/gogo/alwd":
			return alwd.Handler(ctx, event)
		case "/gogo/setu":
			return setu.Handler(ctx, event)
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
			return biliweekly.Handler(ctx, event)
		case "setu":
			return setu.Handler(ctx, event)
		default:
			log.Printf("unknown timer: %+v", eTimer)
			return
		}
	}

	log.Printf("unknown event: %+v", event)
	return
}

func main() {
	var (
		app = flag.String("app", "", "loli")
	)
	flag.Parse()

	switch *app {
	case "setu":
		log.Println("app=setu")
		c := cron.New()
		_, err := c.AddFunc("0 10-19/4 * * *", func() { _handler(context.Background(), EventTimer{TriggerName: "setu"}) })
		if err != nil {
			log.Fatal(err)
		}
		c.Start()
	default:
		log.Println("app=default")
		// Make the handler available for Remote Procedure Call by Cloud Function
		cloudfunction.Start(_handler)
		return
	}
	select {}
}
