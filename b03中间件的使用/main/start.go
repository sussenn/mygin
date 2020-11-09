package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	//使用自定义中间件
	//engine.Use(RequestInfos())	//写在这里,即当前main函数内所有接口都应用此中间件
	//自定义中间件可作为参数,仅针对当前Get函数生效
	engine.GET("/findAll", RequestInfos(), findAllHandler)
	engine.Run(":9001")
}

//自定义 打印请求方式和路径的中间件
//返回值必须是 gin.HandlerFunc
func RequestInfos() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		path := context.FullPath()
		fmt.Printf("先获取 请求方式: %s, 路径: %s\n", method, path)

		//中断函数 即中间件执行到这一步会停止,待后续代码执行完毕,再执行以下代码
		context.Next()

		fmt.Println("最后获取 响应状态码:", context.Writer.Status())
	}
}

func findAllHandler(context *gin.Context) {
	context.JSON(http.StatusOK, map[string]interface{}{
		"code": 1,
		"msg":  "success",
		"data": context.FullPath(),
	})
}
