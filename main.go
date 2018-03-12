package PortfolioWebsite

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("/css", "./web/css")
	router.Static("/js", "./web/js")
	router.Static("/html", "./web/html")
	router.StaticFile("/", "./web/html/index.html")

	router.Run(":8080")
}