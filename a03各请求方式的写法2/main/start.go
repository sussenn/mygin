package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	//get: ?name=xxx
	engine.GET("/hello", func(context *gin.Context) {
		name := context.Query("name")
		context.Writer.Write([]byte("name:" + name))
	})

	//post: x-www
	engine.POST("/login", func(context *gin.Context) {
		//exist: 若key:username 参数不存在,exist为false
		username, exist := context.GetPostForm("username")
		if exist {
			fmt.Println("username=", username)
		}
		context.Writer.Write([]byte("username=" + username))
	})

	//del: rustFull风格
	engine.DELETE("/del/:id", func(context *gin.Context) {
		userId := context.Param("id")
		context.Writer.Write([]byte("userId=" + userId))
	})

	engine.Run(":8090")
}
