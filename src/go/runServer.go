package main

import (
	"PortfolioWebsite/src/go/routes"
	"PortfolioWebsite/src/go/visitorCounter"
	"PortfolioWebsite/src/go/weatherExample"
	"github.com/gin-gonic/gin"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	//sets the root of the directory to access
	router.Static("/css", "./public/css")
	router.Static("/js", "./public/js")
	router.Static("/images", "./public/images")
	router.Static("/unityGames", "./public/unityGames")
	router.StaticFile("", "./public/index.html")

	// Linked to the Star Wars Character Example
	router.POST("/addCharacterToDB", routes.AddCharacterToDB)
	router.POST("/updateCharacter", routes.UpdateCharacter)
	router.GET("/loadAngularJSExampleTableResults", routes.LoadAngularJSExampleTableResults)
	router.GET("/setClickedRow", routes.SetClickedRow)

	// Linked to the Weather Example
	router.GET("/getWeather", routes.GetWeather)
	router.GET("/getWeatherConditions", routes.GetWeatherConditions)
	router.GET("/getLocalWeather", routes.GetLocalWeather)
	router.GET("/getLocalCurrentConditions", routes.GetLocalCurrentConditions)

	// Linked to the Visitor Counter
	router.GET("/visitorCounter", routes.VisitorCounter)
	router.GET("/readVisitors", routes.ReadVisitors)
	router.GET("/readMonthlyVisitors", routes.ReadMonthlyVisitors)
	router.GET("/getIPLocation", routes.GetIPLocation)

	// Linked to the Contact Me Page
	router.POST("/sendMessage", routes.SendMessage)

	// Linked to the Go Examples Page!
	router.GET("/getOccurrences", routes.GetOccurrences)
	router.GET("/translate", routes.Translate)
	//router.POST("/postTweet", routes.PostTweet)

	// Timed functions!
	go weatherExample.InitRequestCount()
	go weatherExample.InitUpdateCurrentConditions()
	go weatherExample.InitUpdateForecast()
	go visitorCounter.InitCreateMonth()
	//go weatherExample.UpdateAllWeather(true)

	// This is the port that runs
	router.Run(":8080")
}
