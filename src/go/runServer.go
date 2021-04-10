package main

import (
	"PortfolioWebsite/src/go/routes"
	"PortfolioWebsite/src/go/visitorCounter"
	"PortfolioWebsite/src/go/weatherExample"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.ForwardedByClientIP = false

	gin.SetMode(gin.ReleaseMode)

	//sets the root of the directory to access
	router.Static("/css", "./public/css")
	router.Static("/js", "./public/js")
	router.Static("/html", "./public/html")
	router.Static("/images", "./public/images")
	router.Static("/unityGames", "./public/unityGames")
	// For our single page website... Minus the Unity stuff as it cannot be
	// closed without reloading the page...
	router.StaticFile("", "./public/html/index.html")

    // We'll house all of our API calls here
    api := router.Group("/api")
    {
        // Cookie stuff
        api.GET("/setCookie", routes.SetCookie)
        api.GET("/getCookie", routes.GetCookie)
        api.GET("/clearCookie", routes.ClearCookie)

        router.GET("/getGithubInfo/:url", routes.GetGithubInfo)

        // Linked to the Star Wars Character Example
        api.POST("/addCharacterToDB", routes.AddCharacterToDB)
        api.POST("/updateCharacter", routes.UpdateCharacter)
        api.GET("/loadAngularJSExampleTableResults", routes.LoadAngularJSExampleTableResults)
        api.GET("/setClickedRow", routes.SetClickedRow)

        // Linked to the Weather Example
        api.GET("/getWeather", routes.GetWeather)
        api.GET("/getWeatherConditions", routes.GetWeatherConditions)
        api.GET("/getLocalWeather", routes.GetLocalWeather)
        api.GET("/getLocalCurrentConditions", routes.GetLocalCurrentConditions)

        // Linked to the Visitor Counter
        api.GET("/visitorCounter", routes.VisitorCounter)
        api.GET("/readVisitors", routes.ReadVisitors)
        api.GET("/readMonthlyVisitors", routes.ReadMonthlyVisitors)
        api.GET("/getIPLocation", routes.GetIPLocation)

        // Linked to the Contact Me Page
        api.POST("/sendMessage", routes.SendMessage)

        // Linked to the Go Examples Page!
        api.GET("/getOccurrences", routes.GetOccurrences)
        api.GET("/translate", routes.Translate)
        api.POST("/postTweet", routes.PostTweet)

        // Linked to the Stock Tracker Page
        api.GET("/getNewInventory", routes.GetNewInventory)

        api.GET("/ping", routes.SendPong)
        api.GET("/readIP", routes.ReadIP)
    }

	// Timed functions!
	go weatherExample.InitRequestCount()
	go weatherExample.InitUpdateCurrentConditions()
	go weatherExample.InitUpdateForecast()
	go visitorCounter.InitCreateMonth()

    // This will have the frontend handle all of the routing
    router.NoRoute(func(c *gin.Context) {
        c.File("./public/html/index.html")
    })

	// This is the port that runs
	router.Run(":8080")
}