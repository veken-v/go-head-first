package main

import (
	"fmt"
	"go-head-first/boot"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	go boot.ServerStart()

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
