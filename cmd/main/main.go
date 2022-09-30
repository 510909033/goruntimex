package main

import (
	"context"
	"github.com/510909033/goruntimex/pkg/controller"
	_ "github.com/felixge/fgprof"
	"log"
	"runtime"
	"strings"
	"time"
)

func main() {
	ctx := context.Background()
	//go func() {
	//	(&model.UserService{}).Visit(ctx, time.Now())
	//}()

	go func() {
		for {
			go (&controller.UserInfoController{}).GetUserInfoAction(ctx)
			//go (&controller.UserInfoController{}).SetNameAction(ctx)
			(&controller.UserInfoController{}).SetNameAction(ctx)
			//time.Sleep(time.Second)
		}
	}()

	time.Sleep(time.Millisecond * 10)

	monotor()

	//spew.Printf("%+v", profile)

	//for k,v:=range profile.stacks

}

func monotor() {
	//var profile = &wallclockProfile{stacks: map[[32]uintptr]*wallclockStack{}}

	ticker := time.NewTicker(time.Millisecond * 100)
	stop := time.After(time.Second)

	for {
		select {
		case <-ticker.C:
			_monotor()
		case <-stop:
			log.Println("stop")
			goto over
		}
	}
over:

	log.Println("汇总")
	for _, stackRecord := range profile.stacks {
		log.Printf("count=%d, names=%s\n", stackRecord.count, strings.Join(stackRecord.Names, ","))
	}

	log.Println("monitor over")
}

var profile = &wallclockProfile{stacks: map[[32]uintptr]*wallclockStack{}}

func _monotor() {
	var stacks []runtime.StackRecord

	fori := 0
	for {
		fori++
		/**
		GoroutineProfile返回n，活动goroutine堆栈配置文件中的记录数量。
		如果len(p) >= n, GoroutineProfile将配置文件复制到p中并返回n，为true。
		如果len(p) < n, GoroutineProfile不改变p并返回n, false。
		大多数客户端应该使用runtime/pprof包，而不是直接调用GoroutineProfile。
		*/
		n, ok := runtime.GoroutineProfile(stacks)
		log.Printf("GoroutineProfile ,第%d次返回的n=%d, ok=%t\n", fori, n, ok)
		if !ok {
			stacks = make([]runtime.StackRecord, int(float64(n)*1.1))
		} else {
			stacks = stacks[0:n]
			break
		}
	}
	log.Printf("stacks的长度为 %d\n", len(stacks))

	//spew.Printf("%#v", stacks)

	profile.Add(stacks)
}

type wallclockProfile struct {
	stacks map[[32]uintptr]*wallclockStack
	ignore []*runtime.Frame
}

type wallclockStack struct {
	frames []*runtime.Frame
	count  int
	Names  []string
}

func (p *wallclockProfile) Add(stackRecords []runtime.StackRecord) {

	log.Println("stackRecords长度为： ", len(stackRecords))

	for k, stackRecord := range stackRecords {
		log.Printf("开始遍历第%d个stackRecord (%#v), \n", k+1, stackRecord)
		if _, ok := p.stacks[stackRecord.Stack0]; !ok {
			ws := &wallclockStack{}
			// symbolize pcs into frames
			frames := runtime.CallersFrames(stackRecord.Stack())
			for {
				frame, more := frames.Next()
				ws.frames = append(ws.frames, &frame)
				log.Printf("function=%s\n", frame.Function)
				ws.Names = append(ws.Names, frame.Function)
				if !more {
					break
				}
			}
			p.stacks[stackRecord.Stack0] = ws
		}
		p.stacks[stackRecord.Stack0].count++
		log.Printf("count=%d, names=%s\n", p.stacks[stackRecord.Stack0].count, strings.Join(p.stacks[stackRecord.Stack0].Names, ","))
	}
}

//var stackNames = make(map[[32]uintptr]string)
