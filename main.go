package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	_ "go_cms_reptile/base" // 导入缓存使用
	"go_cms_reptile/routers"
	"go_cms_reptile/tools"
	"io"
	"log"
	"os"
)

//启动获取资源数据
func main() {
	Conrs := cron.New() //创建一个cron实例

	//执行定时任务（每1小时执行一次）
	err := Conrs.AddFunc("0 0 0/1 * * ? ", tools.DDnsByNameSilo)
	if err != nil {
		fmt.Println(err)
	}

	//启动/
	Conrs.Start()
	defer Conrs.Stop()
	log.Println("定时任务 Conrs 启动完成")
	ginServer()

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
