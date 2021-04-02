package main

import (
	"context"
	"log"
	"scf/handler"

	"github.com/tencentyun/scf-go-lib/cloudfunction"
)

type Event struct {
	HeaderParameters      interface{}    `json:"headerParameters"`
	Headers               Headers        `json:"headers"`
	HTTPMethod            string         `json:"httpMethod"`
	IsBase64Encoded       bool           `json:"isBase64Encoded"`
	Path                  string         `json:"path"`
	PathParameters        interface{}    `json:"pathParameters"`
	QueryString           interface{}    `json:"queryString"`
	QueryStringParameters interface{}    `json:"queryStringParameters"`
	RequestContext        RequestContext `json:"requestContext"`
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

func _handler(ctx context.Context, event Event) (resp interface{}, err error) {
	log.Println("Method", event.HTTPMethod, "Path", event.Path)

	switch event.Path {
	case "/gogo/http":
		return handler.HTTPSource(ctx, event)
	case "/gogo/alwd":
		return handler.ALWD(ctx, event)
	}

	return
}

func main() {
	// Make the handler available for Remote Procedure Call by Cloud Function
	cloudfunction.Start(_handler)
}
