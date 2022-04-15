package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yaji1122/goDict/internal/handler"
)

//var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	// load template
	r.LoadHTMLGlob("views/*")
	// admin index
	r.GET("/index", handler.Index)
	// api
	r.GET("/api/dict/:word", handler.Query)

	return r
}
