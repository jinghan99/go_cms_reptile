package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "go_cms_reptile/base" // 导入缓存使用
	"go_cms_reptile/routers"
	"go_cms_reptile/tools"
	"io"
	"os"
)

//启动获取资源数据
func main() {

	nameSilo()
}

//开启 ginserver
func ginServer() {
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	gin.DisableConsoleColor()

	// 记录到文件。
	f, _ := os.Create("go_cms_reptile.log")

	// 同时将日志写入文件和控制台。
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

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

func nameSilo() {
	domain := "fengyueqing.top"
	apikey := "26c28d52a431307d50bd9"
	records, err := tools.DnsListRecords(domain, apikey)
	if err != nil {
		fmt.Println("nameSilo failed, err:", err)
		return
	}

	fmt.Println("nameSilo records success ", records)

	ip, _ := tools.MyIp()
	fmt.Println("MyIp success ", ip.IP)
}
