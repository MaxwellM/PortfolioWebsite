package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"PortfolioWebsite/goResources/db"
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
	router.GET("/loadAngularJSExampleTableResults", LoadAngularJSExampleTableResults)
	router.GET("/setClickedRow", SetClickedRow)

	// This is the port that runs
	router.Run(":8080")
}

func AddCharacterToDB(data *gin.Context) {
	type PduIpAddress struct {
		Name        string   `json:"name"`
		Homeworld   string   `json:"homeworld"`
		Born        string   `json:"born"`
		Died        string   `json:"died"`
		Gender      string   `json:"gender"`
		Species     string   `json:"species"`
		Affiliation []string `json:"affiliation"`
		Associated  []string `json:"associated"`
		Masters     []string `json:"masters"`
		Apprentices []string `json:"apprentices"`
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

func LoadAngularJSExampleTableResults(data *gin.Context) {
	rows, err := db.ConnPool.Query(
		"select id, name, home_world, born, died, species, gender, affiliation, associated, masters, apprentices from star_wars_characters")

	if err != nil {
		fmt.Println("There was an error reading the star_wars_characters table from the database:", err)
	}

	type CharacterResult struct {
		Id          int      `json:"id"`
		Name        string   `json:"name"`
		Homeworld   string   `json:"homeworld"`
		Born        string   `json:"born"`
		Died        string   `json:"died"`
		Species     string   `json:"species"`
		Gender      string   `json:"gender"`
		Affiliation []string `json:"affiliation"`
		Associated  []string `json:"associated"`
		Masters     []string `json:"masters"`
		Apprentices []string `json:"apprentices"`
	}

	characterResultsArray := []*CharacterResult{}

	defer rows.Close()

	for rows.Next() {

		var characterResult CharacterResult

		err = rows.Scan(
			&characterResult.Id,
			&characterResult.Name,
			&characterResult.Homeworld,
			&characterResult.Born,
			&characterResult.Died,
			&characterResult.Species,
			&characterResult.Gender,
			&characterResult.Affiliation,
			&characterResult.Associated,
			&characterResult.Masters,
			&characterResult.Apprentices)

		if err != nil {
			fmt.Println("There was an error querying the database for the Star Wars Character Results:", err)
			continue
		}

		characterResultsArray = append(characterResultsArray, &characterResult)
	}

	data.JSON(http.StatusOK, characterResultsArray)
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
