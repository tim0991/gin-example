package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	api := r.Group("/api")
	{
		// 这个get是 RouterGroup的方法, 对engine的路由方法做了代理
		api.GET("/users", func(c *gin.Context) {
			c.JSON(http.StatusOK, []string{"tim", "wang"})
		})
	}

	r.Run(":3000")
}
