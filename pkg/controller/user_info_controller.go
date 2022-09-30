package controller

import (
	"context"
	"github.com/510909033/goruntimex/pkg/util"
	"sync"
	"time"
)

type UserInfoController struct {
}

func (controller *UserInfoController) GetUserInfoAction(ctx context.Context) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		util.Sleep500()
	}()
	//go func() {
	//	defer wg.Done()
	//	util.Sleep200()
	//}()

	wg.Wait()
}

func (controller *UserInfoController) SetNameAction(ctx context.Context) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		util.Sleep500()
	}()
	//go func() {
	//	defer wg.Done()
	//	util.Sleep200()
	//}()

	wg.Wait()

	//util.Sleep500()
	//util.Sleep200()
	//util.Sleep1500()

	time.Sleep(time.Second)
}
