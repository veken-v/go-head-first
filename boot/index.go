package boot

import "fmt"

func init() {

	fmt.Println("---配置加载---")
	loadEnvironment()

}
