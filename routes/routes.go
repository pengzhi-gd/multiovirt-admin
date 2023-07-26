package routes

import (
	"example/fundemo01/web-app/controllers"
	"example/fundemo01/web-app/logger"
	"example/fundemo01/web-app/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitRoutes() *gin.Engine{
	start_mode := viper.GetString("app.start_mode")
	switch start_mode {
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(),logger.GinRecovery(true))
	v1 := r.Group("/api/v1")
	//用户注册路由
	v1.POST("/signup",controllers.SignUpHandler)
	v1.POST("/login",controllers.LoginHandler)
	v1.GET("/ping",middleware.JwtAuthMiddleware(),func(c *gin.Context) {
		controllers.ResponseSuccess(c,controllers.CodeResponseSuccess)
	})
	v1.Use(middleware.JwtAuthMiddleware())
	{
		v1.GET("/community",controllers.CommunityList)
		v1.GET("/community/:id",controllers.CommunityDetail)
		v1.POST("/post",controllers.CreatePost)
		v1.POST("/createovirtconf",controllers.CreateOvirtConf)
		v1.GET("/listovirtconf",controllers.ListOvirtConf)
		v1.GET("/ovirtconfdetail/:aliasname",controllers.GetOvirtConfDetail)
	}
	r.NoRoute(func(c *gin.Context){
		controllers.ResponseError(c,controllers.CodeNoRoute)
	})

	return r
}
