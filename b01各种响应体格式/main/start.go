package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()

	//string类型的响应
	engine.GET("/hello00", func(context *gin.Context) {
		context.Writer.WriteString(context.FullPath())
	})

	//json map类型
	engine.GET("/hello01", func(context *gin.Context) {
		path := context.FullPath()
		context.JSON(200, map[string]interface{}{
			"code":    20001,
			"message": "ok",
			"data":    path,
		})
	})

	//json struct类型
	engine.GET("/hello02", func(context *gin.Context) {
		path := context.FullPath()
		resp := Response{
			Code: 20000,
			Msg:  "success",
			Data: path,
		}
		context.JSON(http.StatusOK, &resp)
	})

	//响应html页面及图片
	//设置html文件目录
	engine.LoadHTMLGlob("./html/*")
	//设置静态文件目录		p1:访问的目录 p2:对应的工程目录
	engine.Static("/img", "./img")

	engine.GET("/hello03", func(context *gin.Context) {
		//响应内容通过template模板输出的页面
		context.HTML(200, "index.html", gin.H{
			"myPath": context.FullPath(),
		})
	})

	engine.Run(":8010")
}

type Response struct {
	Code int
	Msg  string
	Data interface{}
}
