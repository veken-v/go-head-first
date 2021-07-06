package controller

import (
	"github.com/gin-gonic/gin"
)

func queryUser(c *gin.Context) {
	//c.HTML(http.StatusOK, "query-user", gin.H{
	//	"version": time.Now().String(),
	//	"page":    "add_stream",
	//})
	c.String(200, "json")

}
