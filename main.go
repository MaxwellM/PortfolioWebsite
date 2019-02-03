package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"PortfolioWebsite/goResources"
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

	router.Run(":8080")
}

func AddCharacterToDB(data *gin.Context) {
	type PduIpAddress struct {
		Name        string                 `json:"name"`
		Born        string                 `json:"born"`
		Associated  map[string]interface{} `json:"associated"`
		Gender      string                 `json"gender"`
		Affiliation map[string]interface{} `json:"affiliation"`
		Masters     map[string]interface{} `json:"masters"`
	}

	var info PduIpAddress
	data.Bind(&info)

	//name := data.DefaultQuery("name", "")
	//born := data.DefaultQuery("born", "")
	//associated := data.DefaultQuery("associated", map[string]interface{})
	//gender := data.DefaultQuery("gender", "")
	//affiliation := data.DefaultQuery("affiliation", "")
	//masters := data.DefaultQuery("masters", "")

	characterReturn, err := goResources.AddCharacter(info.Name, info.Born, info.Associated, info.Gender, info.Affiliation, info.Masters)
	if err != nil {
		data.String(http.StatusBadRequest, "Failed to add character.", err)
	} else {
		data.String(http.StatusOK, "successfully added Charcter!", characterReturn)
	}
}
