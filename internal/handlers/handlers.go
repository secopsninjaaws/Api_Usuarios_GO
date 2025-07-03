package handlers

import "github.com/gin-gonic/gin"

func Handlers() {
	router := gin.Default()
	{
		v1 := router.Group("/v1")
		v1.POST("/post", PostCreate)
		v1.GET("/list", RetornarTodos)
	}
	router.Run() // listen and serve on 0.0.0.0:8080
}
