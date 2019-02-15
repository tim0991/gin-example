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
	// localhost:3000/welcome?firstname=Jane&lastname=Doe&ids[]=1&ids[]=2&names[a]=tim&names[b]=wang
	r.GET("/welcome", func(c *gin.Context) {
		// 下面两个都是通过getQuery实现
		fname := c.Query("firstname")
		// alias of Query
		lname := c.DefaultQuery("lastname", "tim")

		fnameArr, _ := c.GetQueryArray("firstname")
		ids := c.QueryArray("ids[]")
		names := c.QueryMap("names")
		// 和query结果一致
		qs := c.Request.URL.Query()
		fnameFromUrlQuery := qs.Get("firstname")

		c.String(http.StatusOK, "firstname:%s, lastname:%s, firstnameArr:%#v, qs:%#v,  fnameFromUrlQuery:%#v names: %#v, ids2:%#v", fname, lname, fnameArr, qs, fnameFromUrlQuery, ids, names)
	})


	// test post 与get查询方法名几乎一致
	r.POST("/form", func(c *gin.Context) {
		name := c.PostForm("name")
		// c.PostFormArray(key string)
		c.String(http.StatusOK, name)
	})

	r.Run(":3000")
}
