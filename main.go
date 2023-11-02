package main

import (
	"gee/gin"
	"log"
)

func main() {
	//http.HandleFunc("/hello", helloHandler)
	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	})

	r.GET("/good", func(c *gin.Context) {
		c.String(200, "good %s", c.Query("name"))
	})

	r.POST("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	log.Fatal(r.Run(":8080"))
}
