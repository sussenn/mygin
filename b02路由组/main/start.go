package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default()

	routerGroup := engine.Group("/user")
	//加上路由组前缀,完整的路径即: /user/register
	routerGroup.POST("/register", registerHandler)
	routerGroup.DELETE("/delById/:id", delHandler)

	engine.Run(":8090")
}

//提取出来,解耦合
func registerHandler(context *gin.Context) {
	context.Writer.Write([]byte(context.FullPath()))
}

func delHandler(context *gin.Context) {
	id := context.Param("id")
	context.Writer.WriteString("删除的id是:" + id)
}
