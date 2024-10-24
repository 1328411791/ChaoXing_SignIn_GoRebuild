package controller

import (
	"ChaoXing_SignIn_GoRebuild/src/mysql"
	"github.com/gin-gonic/gin"
	"log"
)

type ChaoxingController struct {
	dataSource *mysql.DataSouce
}

func NewChaoxingControllerController(dataSource *mysql.DataSouce, router *gin.Engine) ChaoxingController {
	controller := ChaoxingController{
		dataSource: dataSource,
	}
	controller.RegisterRoutes(router)

	return controller
}

func (this *ChaoxingController) RegisterRoutes(router *gin.Engine) {
	router.GET("/login", this.login)
	router.POST("/activity", this.activity)
	router.POST("/activityAll", this.activityAll)
	router.POST("/qrcode", this.qrcode)
	router.POST("/location", this.location)
	router.POST("/general", this.general)
	router.POST("/getSlideCaptcha", this.getSlideCaptcha)
	router.POST("/checkSlideCaptcha", this.checkSlideCaptcha)
	router.POST("/uvtoken", this.uvtoken)
	router.POST("/upload", this.upload)
}

// login 接口
func (this *ChaoxingController) login(ctx *gin.Context) {
	log.Fatal("未实现")
}

// activity 接口
func (this *ChaoxingController) activity(ctx *gin.Context) {
	// 未实现
	log.Fatal("未实现")
}

// activityAll 接口
func (this *ChaoxingController) activityAll(context *gin.Context) {
	log.Fatal("未实现")
}

// qrcode 接口
func (this *ChaoxingController) qrcode(context *gin.Context) {
	log.Fatal("未实现")
}

// location 接口
func (this *ChaoxingController) location(context *gin.Context) {
	log.Fatal("未实现")
}

// general 接口
func (this *ChaoxingController) general(context *gin.Context) {
	log.Fatal("未实现")
}

// getSlideCaptcha 接口
func (this *ChaoxingController) getSlideCaptcha(context *gin.Context) {
	log.Fatal("未实现")
}

// checkSlideCaptcha 接口
func (this *ChaoxingController) checkSlideCaptcha(context *gin.Context) {
	log.Fatal("未实现")
}

// uvtoken 接口
func (this *ChaoxingController) uvtoken(context *gin.Context) {
	log.Fatal("未实现")
}

// upload 接口
func (this *ChaoxingController) upload(context *gin.Context) {
	log.Fatal("未实现")
}
