package routers

import (
	"net/http"
	"recruit/controller"
	"recruit/settings"

	"github.com/gin-gonic/gin"
)

func SetupRouter(config settings.AppConfig) *gin.Engine {
	if config.Mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	settings.StaticPath = config.StaticFilePath
	r := gin.New()

	api := r.Group("/api")
	{
		api.POST("/register/sendMsg", controller.RegisterMsg)
		api.POST("/register", controller.Register)
		api.POST("/login", controller.Login)
	}
	withToken := r.Group("/withToken" /*, middlewares.ValidateToken*/)
	{

		withToken.POST("/postavator/:uid", controller.PostAvator)

		withToken.GET("/mainPageJS/:uid", controller.JobSeekPage)         //找工作页面
		withToken.GET("/jobSeekDetails/:jsId", controller.JobSeekDetails) //求职贴详情
		withToken.POST("/postJobSeek/:uid", controller.PostJobSeek)       //上传求职贴
		withToken.POST("/updateJsStatus", controller.UpdateJsStatus)      //修改求职贴状态
		withToken.POST("/updateJsDetails", controller.UpdateJsDetails)    //修改求职贴详情

		withToken.GET("/mainPagePS/:uid", controller.PeopleSeekPage)
		withToken.GET("/peopleSeekDetails/:psId", controller.PeopleSeekDetails)
		withToken.POST("/postPeopleSeek/:uid", controller.PostPeopleSeek)
		withToken.POST("/updatePsStatus", controller.UpdatePsStatus)
		withToken.POST("/updatePsDetails", controller.UpdatePsDetails)

		community := withToken.Group("/community")
		{
			community.POST("/postCommodities", controller.PostCommodities)
			community.GET("/getMyCommodities", controller.GetMyCommodities)
			community.GET("/getCommodityDetails", controller.GetCommodityDetails)
			community.POST("/updateCommodity", controller.UpdateCommodity)
			community.POST("/updateCmdtImage", controller.UpdateCmdtImage)
			community.GET("/getCommodities", controller.GetCommodities)
			community.DELETE("/deleteCommodity", controller.DeleteCommodity)
		}
		rank := withToken.Group("/rank")
		{
			rank.GET("/rankingList", controller.Rank)
		}
		personality := withToken.Group("personality")
		{
			personality.GET("/me", controller.Me)
		}
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "40412323156",
		})
	})
	return r
}
