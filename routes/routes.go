package routes

import (
	"yuequan/controller"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) {
	r.Static("/statics", "./statics")
	r.LoadHTMLGlob("templates/*")
	r.StaticFile("/favicon.ico", "./statics/favicon.ico")
	r.GET("/", controller.IndexHandler)                                     // index页面
	r.GET("/api/clear", controller.Clearhandler(), controller.IndexHandler) // 清除扫描结果
	r.POST("/", controller.Scanhandler(), controller.IndexHandler)          // 生成扫描任务
	r.GET("/api/getdata", controller.ScanResultHandler())                   // 获取扫描结果
	r.GET("/api/getconfig", controller.GetConfigHandler)                    // 获取配置
	r.GET("/api/stopscan", controller.StopScanHandler)                      // 暂停扫描任务
	r.POST("/api/replay", controller.ReplayScanHandler)                     // 操作 - 重放
	r.POST("/api/addwhite", controller.AddWhiteHandler)                     // 操作 - 加白
}
