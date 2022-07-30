package routers

import (
	"github.com/gin-gonic/gin"
	"go_cms_reptile/controllers"
)

func ReptileInit(r *gin.Engine) {

	//指定路由
	r.GET("/", controllers.Index)
	r.GET("/video_Detail", controllers.VideoDetail)
	r.GET("/search_name", controllers.SearchName)
	r.GET("/video_playlist", controllers.VideoPlaylist)
}
