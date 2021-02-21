package main

import (
	"github.com/EikoNakashima/test-go3/db"
	"github.com/EikoNakashima/test-go3/server"
)

func main() {
	// r := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 	c.String(200, "Hello, World")
	// })
	db.Init()
	server.Init()
	// r.Run()
	db.Close()
}
