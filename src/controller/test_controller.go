package controller

import (
	"ChaoXing_SignIn_GoRebuild/src/mysql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type TestController struct {
	dataSource *mysql.DataSouce
}

func NewTestController(dataSource *mysql.DataSouce, router *gin.Engine) TestController {
	controller := TestController{
		dataSource: dataSource,
	}
	controller.RegisterRoutes(router)

	return controller
}

func (this *TestController) RegisterRoutes(router *gin.Engine) {
	// 前缀
	pre := "/test"
	router.GET(pre+"/", this.helloWorld)

}

func (this *TestController) helloWorld(ctx *gin.Context) {
	bookDao := this.dataSource.TestDao()
	id, err := bookDao.ListTestById(1)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(id)
	ctx.JSON(http.StatusOK, id)
}
