package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"PortfolioWebsite/goResources/starWarsCharacterTableEample"
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

	router.POST("/addCharacterToDB", AddCharacterToDB)
	router.POST("/updateCharacter", UpdateCharacter)
	router.GET("/loadAngularJSExampleTableResults", LoadAngularJSExampleTableResults)
	router.GET("/setClickedRow", SetClickedRow)

	// This is the port that runs
	router.Run(":80")
}

func AddCharacterToDB(data *gin.Context) {
	type PduIpAddress struct {
		Name        string   `json:"name"`
		Homeworld   string   `json:"homeworld"`
		Born        string   `json:"born"`
		Died        string   `json:"died"`
		Gender      string   `json:"gender"`
		Species     string   `json:"species"`
		Affiliation string `json:"affiliation"`
		Associated  string `json:"associated"`
		Masters     string `json:"masters"`
		Apprentices string `json:"apprentices"`
	}

	var info PduIpAddress
	data.Bind(&info)

	fmt.Println("INFO: ", info)

	characterReturn, err := starWarsCharacterTableEample.AddCharacter(
		info.Name,
		info.Homeworld,
		info.Born,
		info.Died,
		info.Gender,
		info.Species,
		info.Affiliation,
		info.Associated,
		info.Masters,
		info.Apprentices)
	if err != nil {
		data.String(http.StatusBadRequest, "Failed to add character.", err)
	} else {
		data.String(http.StatusOK, "successfully added Charcter!", characterReturn)
	}
}

func UpdateCharacter(data *gin.Context) {
	type temp struct {
		Character map[string]interface{} `json:"character"`
	}

	var info temp
	data.Bind(&info)

	fmt.Println("EDITED CHARACTER: ", info.Character)

	characterReturn, err := starWarsCharacterTableEample.ResubmitCharacter(info.Character)
	if err != nil {
		data.JSON(http.StatusBadRequest, err)
		fmt.Println("Error re-submitting the character", err)
	} else {
		data.JSON(http.StatusOK, characterReturn)
	}
}

func LoadAngularJSExampleTableResults(data *gin.Context) {
	starWarsCharactersReturn, err := starWarsCharacterTableEample.LoadAllStarWarsCharacters()
	if err != nil {
		data.JSON(http.StatusBadRequest, err)
		fmt.Println("Error obtaining all Star Wars Characters", err)
	} else {
		data.JSON(http.StatusOK, starWarsCharactersReturn)
	}
}

func SetClickedRow(data *gin.Context) {
	id := data.DefaultQuery("id", "")

	fmt.Println("ID: ", id)

	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("YOU SUCK MAX", err)
	}

	fmt.Println("ID AFTER: ", idInt)
	quoteBuilderReturn, err := starWarsCharacterTableEample.RetreiveCharacter(idInt)
	if err != nil {
		data.JSON(http.StatusBadRequest, err)
		fmt.Println("Error obtaining a particular Quote", err)
	} else {
		data.JSON(http.StatusOK, quoteBuilderReturn)
	}
}
