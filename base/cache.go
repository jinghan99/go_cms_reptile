// Package base 初始插件
package base

import (
	"github.com/patrickmn/go-cache"
	"time"
) // 使用前先import包

// Cache 定义全局变量 cache
var Cache *cache.Cache

//初始化 缓存框架 go-cache
func init() {
	//创建一个cache对象，默认ttl 5分钟，每10分钟对过期数据进行一次清理
	Cache = cache.New(5*time.Minute, 60*time.Second)

	// 设置缓存值并带上过期时间
	Cache.Set("foo", "bar", cache.DefaultExpiration)

	// 设置没有过期时间的KEY，这个KEY不会被自动清除，想清除使用：c.Delete("baz")
	Cache.Set("lastVideo", "雪中悍刀行", cache.NoExpiration)
}

func GetCacheOne(key string) interface{} {
	i, found := Cache.Get(key)
	if found {
		return i
	}
	return nil
}
