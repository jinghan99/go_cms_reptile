package controllers

import (
	"github.com/gin-gonic/gin"
	"go_cms_reptile/tools"
)

// Index 首页目录
func Index(c *gin.Context) {
	//url 传值
	searchName := c.DefaultQuery("searchName", "")

	kanYuModel, err := tools.SearchName(searchName)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "searchName err" + err.Error(),
		})
	} else {
		c.JSON(200, kanYuModel)
	}
}

// VideoDetail 影片详情
func VideoDetail(c *gin.Context) {
	//url 传值
	id := c.DefaultQuery("id", "")
	if id == "" {
		c.JSON(500, gin.H{
			"message": "VideoDetail id is nil",
		})
	}
	kanYuDetailModel, err := tools.VideoDetail(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "VideoDetail err" + err.Error(),
		})
	} else {
		c.JSON(200, kanYuDetailModel)
	}
}

// VideoPlaylist 影片播放列表
func VideoPlaylist(c *gin.Context) {
	//url 传值
	id := c.DefaultQuery("id", "")
	if id == "" {
		c.JSON(500, gin.H{
			"message": "VideoDetail id is nil",
		})
	}
	videoPlayModel, err := tools.VideoPlaylist(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "VideoDetail err" + err.Error(),
		})
	} else {
		c.JSON(200, videoPlayModel)
	}
}
