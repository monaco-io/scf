package http

import (
	"context"
)

func Handler(ctx context.Context, event interface{}) (resp interface{}, err error) {
	resp = event
	return
}
