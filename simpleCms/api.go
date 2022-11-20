package main

import (
	"github.com/gin-gonic/gin"
	"simpleCms/app/content"
)

var router = gin.Default()

func init()  {
	router.GET("/content/:id", content.Get)
	router.DELETE("/content/:id", content.Delete)
	router.PUT("/content/:id", content.Put)
	router.POST("/content/:id", content.Post)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//定义任意的url, api
}