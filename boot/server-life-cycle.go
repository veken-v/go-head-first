//服务器的生命周期
package boot

import "github.com/gin-gonic/gin"

type ServerLifeCycle interface {
	//启动之前
	BeforeStart(Environment) Environment
	//创建完成
	Created(ginInstance *gin.Engine)
	//运行中
	Running(ginInstance *gin.Engine)
}

//生命周期模板
type ServerLifeCycleTemplate struct {
}

func (t ServerLifeCycleTemplate) BeforeStart(env Environment) Environment {
	return env
}

func (t ServerLifeCycleTemplate) Created(ginInstance *gin.Engine) {

}

func (t ServerLifeCycleTemplate) Running(ginInstance *gin.Engine) {

}
