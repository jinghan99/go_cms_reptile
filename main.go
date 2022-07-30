package main

import (
	"github.com/gin-gonic/gin"
	_ "go_cms_reptile/base" // 导入缓存使用
	"go_cms_reptile/routers"
)

//启动获取资源数据
func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	//静态文件
	r.Static("/static", "./static")

	// 路由抽离单独文件 分组
	routers.ReptileInit(r)

	err := r.Run(":8999")
	if err != nil {
		return
	}
}
