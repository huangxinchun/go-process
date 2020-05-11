package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/huangxinchun/go-demo/testpackage"
	"github.com/huangxinchun/hxcgo"
	"github.com/huangxinchun/hxcgo/my"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
	testpackage.TestNew()
	fmt.Println("main")
	my.TestMy()
	hxcgo.TestGin()
}
