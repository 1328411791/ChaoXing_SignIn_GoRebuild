package main

import (
	"ChaoXing_SignIn_GoRebuild/src/controller"
	"ChaoXing_SignIn_GoRebuild/src/mysql"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ginSetUpRouter() *gin.Engine {
	dataSouce := mysql.NewDataSource()
	S := gin.Default()
	controller.NewTestController(dataSouce, S)

	return S
}

func main() {
	S := ginSetUpRouter()
	err := S.Run(":8088")
	if err != nil {
		fmt.Println("服务器启动失败！")
	}
}
