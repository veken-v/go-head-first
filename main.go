package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-head-first/boot"
	_ "go-head-first/logging"
	"go-head-first/router"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type ServerInitializer struct {
	boot.ServerLifeCycleTemplate
}

func (l ServerInitializer) Created(ginInstance *gin.Engine) {
	router.RouteMount(ginInstance)
}

func main() {
	go boot.ServerStart(ServerInitializer{})

	//构建一个停机信号
	signalChanel := make(chan os.Signal, 1)
	signal.Notify(signalChanel, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)
	go func() {
		sin := <-signalChanel
		fmt.Println("收到退出信息", sin)
		done <- true
	}()
	<-done
	//做一些清理工作
	time.Sleep(2 * time.Second)
	fmt.Println("系统已经关闭！！！")
}
