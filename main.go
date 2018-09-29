package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	//sets the root of the directory to access
	router.Static("/css", "./web/css")
	router.Static("/js", "./web/js")
	router.Static("/images", "./web/images")
	router.Static("/html", "./web/html")
	router.StaticFile("/", "./web/html/index.html")

	router.Run(":8080")
}