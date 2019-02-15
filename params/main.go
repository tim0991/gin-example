package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// test router named param
	// url: localhost:3000/user/tim
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, name)
	})

	// url: localhost:3000/user/tim/create
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		c.String(http.StatusOK, name+" "+action)
	})

	// test query string
	// url /welcome?firstname=Jane&lastname=Doe
	r.GET("/welcome", func(c *gin.Context) {
		// 下面两个都是通过getQuery实现
		fname := c.Query("firstname")
		lname := c.DefaultQuery("lastname", "tim")
		queryArr, _ := c.GetQueryArray("firstname")
		c.String(http.StatusOK,"%s %s %#v", fname, lname, queryArr)
	})

	r.Run(":3000")
}
