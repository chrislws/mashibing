package content

//control

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//action
func Get(ctx *gin.Context)  {
	ctx.JSON(http.StatusOK, gin.H{
		"id": ctx.Param("id"),
		"message": "内容查询",
	})
}

func Delete(ctx *gin.Context)  {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "内容删除",
	})
}

func Put(ctx *gin.Context)  {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "内容更新",
	})
}

func Post(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "内容添加",
	})
}