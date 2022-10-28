package main

import (
	"go-casbin-study/lib"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(lib.Middlewares()...)
	r.GET("/depts", func(context *gin.Context) {
		context.JSON(200, gin.H{"result": "部门列表"})
	})
	r.POST("/depts", func(context *gin.Context) {
		context.JSON(200, gin.H{"result": "批量修改部门列表"})
	})
	r.GET("/depts/:id", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"result": "详情"})
	})
	r.Run(":8080")

}
