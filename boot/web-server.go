package boot

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

func ServerStart(lifeCycle ServerLifeCycle) {
	fmt.Println("服务启动了！！！")
	if lifeCycle == nil {
		lifeCycle = ServerLifeCycleTemplate{}
	}

	env = lifeCycle.BeforeStart(env)

	gin.SetMode(env.Server.Mode)

	var ginInstance *gin.Engine

	if env.Server.Mode == gin.ReleaseMode {
		ginInstance = gin.New()
	} else {
		ginInstance = gin.Default()
	}
	lifeCycle.Created(ginInstance)
	//跨域处理
	//ginInstance.Use(CrossOrigin())
	//privat := public.Group("/", gin.BasicAuth(gin.Accounts{Storage.ServerHTTPLogin(): Storage.ServerHTTPPassword()}))
	//html 模板
	//ginInstance.LoadHTMLGlob("/templates/*")
	//静态资源
	ginInstance.StaticFS("/static", http.Dir("/static"))
	//https 证书配置
	//if Storage.ServerHTTPS() {
	//	go func() {
	//		err := public.RunTLS(Storage.ServerHTTPSPort(), Storage.ServerHTTPSCert(), Storage.ServerHTTPSKey())
	//		if err != nil {
	//			log.WithFields(logrus.Fields{
	//				"module": "http_router",
	//				"func":   "HTTPSAPIServer",
	//				"call":   "ServerHTTPSPort",
	//			}).Fatalln(err.Error())
	//			os.Exit(1)
	//		}
	//	}()
	//}
	port := strconv.Itoa(env.Server.Port)
	err := ginInstance.Run(":" + port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		os.Exit(1)
	}

	lifeCycle.Running(ginInstance)
}

//func CrossOrigin() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
//		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
//		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
//		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
//		if c.Request.Method == "OPTIONS" {
//			c.AbortWithStatus(204)
//			return
//		}
//		c.Next()
//	}
//}
