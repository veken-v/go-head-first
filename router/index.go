package router

import "github.com/gin-gonic/gin"

func RouteMount(ginInstance *gin.Engine) {
	ginInstance.GET("/ping", queryUser)
}
