package main

import (
	"github.com/gin-gonic/gin"
	"go_cms_reptile/routers"
)

//启动获取资源数据
func main() {
	r := gin.Default()
	//静态文件
	r.Static("/static", "./static")

	// 路由抽离单独文件 分组
	routers.ReptileInit(r)

	err := r.Run(":8999")
	if err != nil {
		return
	}
}
