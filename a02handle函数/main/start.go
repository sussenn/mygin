package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	//get请求,?参数的接收
	engine.Handle("GET", "/hello", func(context *gin.Context) {
		fmt.Println("请求路径:", context.FullPath())
		//localhost:8090/hello?name=xxx
		//接收 name 参数	p2:设置默认值,如name为空则,name=aaa
		name := context.DefaultQuery("name", "aaa")
		//输出
		context.Writer.Write([]byte("name," + name))
	})

	//post请求,x-www表单参数的接收
	engine.Handle("POST", "/login", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")
		context.Writer.Write([]byte("name:" + username + "pwd:" + password))
	})

	engine.Run(":8090")
}
