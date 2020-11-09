package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()

	//localhost:8090/hello?name=xxx&classes=aaa
	engine.GET("/hello", func(context *gin.Context) {
		var student Student
		err := context.ShouldBindQuery(&student)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		context.Writer.Write([]byte("name:" + student.Name + "classes:" + student.Classes))
	})

	//post: x-www
	engine.POST("/register", func(context *gin.Context) {
		var user User
		if err := context.ShouldBind(&user); err != nil {
			log.Fatal(err.Error())
			return
		}
		context.Writer.Write([]byte("username:" + user.UserName))
	})

	//post: json
	engine.POST("/add", func(context *gin.Context) {
		var person Person
		if err := context.BindJSON(&person); err != nil {
			log.Fatal(err.Error())
			return
		}
		context.Writer.Write([]byte("name:" + person.Name))
	})

	engine.Run(":8090")
}

type Student struct {
	Name    string `form:"name"`
	Classes string `form:"classes"`
}

type User struct {
	UserName string `from:"username"`
	Phone    string `from:"phone"`
	Password string `form:"password"`
}

type Person struct {
	Name string `form:"name"`
	Sex  string `form:"sex"`
	Age  int    `form:"age"`
}
