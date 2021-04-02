package handler

import (
	"context"
)

func HTTPSource(ctx context.Context, event interface{}) (resp interface{}, err error) {
	resp = event
	return
}
