package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"go_cms_reptile/base"
	"net/http"
)

// Index 首页
func Index(c *gin.Context) {

	// 设置没有过期时间的KEY，这个KEY不会被自动清除，想清除使用：c.Delete("baz")
	base.Cache.Set("lastVideo", "雪中悍刀行", cache.NoExpiration)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"lastVideo": base.GetCacheOne("lastVideo"),
	})
}
