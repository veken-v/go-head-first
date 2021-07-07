package router

import (
	"github.com/gin-gonic/gin"
	. "go-head-first/loging"
)

func queryUser(c *gin.Context) {
	//c.HTML(http.StatusOK, "query-user", gin.H{
	//	"version": time.Now().String(),
	//	"page":    "add_stream",
	//})
	ZapLog.Info("我的用户请求")
	ZapSugarLog.Error("这是错误的处理")
	c.String(200, "json")

}
