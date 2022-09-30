package model

import (
	"context"
	"time"
)

type UserService struct {
}

func (service *UserService) Visit(ctx context.Context, t time.Time) {

	time.Sleep(time.Second)
}
