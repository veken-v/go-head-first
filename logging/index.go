package logging

import (
	"fmt"
	"go-head-first/boot"
)

var loggingConfigurer boot.LoggingConfigurer

func init() {
	fmt.Println("--日志初始化--")
	loggingConfigurer = boot.Env().Logging
	initZap()
}
