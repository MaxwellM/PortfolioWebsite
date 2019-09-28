package routes

import (
	"PortfolioWebsite/src/go/contactMe"
	"PortfolioWebsite/src/go/goExamples"
	"PortfolioWebsite/src/go/starWarsCharacterTableEample"
	"PortfolioWebsite/src/go/visitorCounter"
	"PortfolioWebsite/src/go/weatherExample"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func VisitorCounter(data *gin.Context) {
	//IP := data.Request.RemoteAddr
	//IP := data.ClientIP()
	IP := data.GetHeader("X-Real-IP")

	fmt.Println("FOUND IP: ", IP)

	header := data.Request.Header
	fmt.Println("HEADER: ", header)

	//IP, err := visitorCounter.GetClientIPHelper(data.Request)
	//if err != nil {
	//	fmt.Println("Error parsing for IP!")
	//	data.JSON(http.StatusBadRequest, err)
	//}

	fmt.Println("IP: ", IP)

	ips, err := visitorCounter.CheckIfIPExists(IP)
	if err != nil {
		fmt.Println("Error returning IPs who visited the site! ", err)
		data.JSON(http.StatusBadRequest, err.Error())
	} else {
		fmt.Println("IPs: ", ips)
		data.JSON(http.StatusOK, ips)
	}

}

func ReadVisitors(data *gin.Context) {
	visitorsReturn, err := visitorCounter.ReadIPDB()
	if err != nil {
		data.JSON(http.StatusBadRequest, err)
		fmt.Println("Error obtaining visitors report", err.Error())
	} else {
		data.JSON(http.StatusOK, visitorsReturn)
	}
}

func ReadMonthlyVisitors(data *gin.Context) {
	visitorsReturn, err := visitorCounter.ReadMonthlyVisitorsDB()
	if err != nil {
		data.JSON(http.StatusBadRequest, err)
		fmt.Println("Error obtaining visitors report", err.Error())
	} else {
		data.JSON(http.StatusOK, visitorsReturn)
	}
}

func AddCharacterToDB(data *gin.Context) {
	type PduIpAddress struct {
		Name        string `json:"name"`
		Homeworld   string `json:"homeworld"`
		Born        string `json:"born"`
		Died        string `json:"died"`
		Gender      string `json:"gender"`
		Species     string `json:"species"`
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
		data.String(http.StatusBadRequest, "Failed to add character.", err.Error())
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
		data.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("Error re-submitting the character", err)
	} else {
		data.JSON(http.StatusOK, characterReturn)
	}
}

func LoadAngularJSExampleTableResults(data *gin.Context) {
	Name := data.DefaultQuery("Name", "")
	Species := data.DefaultQuery("Species", "")

	starWarsCharactersReturn, err := starWarsCharacterTableEample.LoadAllStarWarsCharacters(Name, Species)
	if err != nil {
		data.JSON(http.StatusBadRequest, err.Error())
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
		data.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("Error obtaining a particular Quote", err)
	} else {
		data.JSON(http.StatusOK, quoteBuilderReturn)
	}
}

func GetWeather(data *gin.Context) {
	location := data.DefaultQuery("location", "351219")

	fmt.Println("Location: ", location)

	weatherReturn, err := weatherExample.UpdateAllWeather(location)
	if err != nil {
		data.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("Error obtaining a weather report", err)
	} else {
		data.JSON(http.StatusOK, weatherReturn)
	}
}

func GetWeatherConditions(data *gin.Context) {
	location := data.DefaultQuery("location", "351219")

	fmt.Println("Location: ", location)

	weatherReturn, err := weatherExample.GetCurrentConditions(location)
	if err != nil {
		data.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("Error obtaining a weather conditions report", err)
	} else {
		data.JSON(http.StatusOK, weatherReturn)
	}
}

func GetLocalWeather(data *gin.Context) {
	weatherReturn, err := weatherExample.ReadLocalWeatherReport()
	if err != nil {
		data.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("Error reading local weather report")
	} else {
		data.JSON(http.StatusOK, weatherReturn)
	}
}

func GetLocalCurrentConditions(data *gin.Context) {
	currentConditionsReturn, err := weatherExample.ReadLocalCurrentConditions()
	if err != nil {
		data.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("Error reading local weather report")
	} else {
		data.JSON(http.StatusOK, currentConditionsReturn)
	}
}

func GetIPLocation(data *gin.Context) {
	//ip := data.DefaultQuery("ip", "")

	//fmt.Println("IP: ", ip)

	ipLocationReturn, err := visitorCounter.ReadIPLocationDB()
	if err != nil {
		data.JSON(http.StatusBadRequest, err)
		fmt.Println("Error obtaining a location for an IP", err.Error())
	} else {
		data.JSON(http.StatusOK, ipLocationReturn)
	}
}

func SendMessage(data *gin.Context) {
	type MessageContents struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Phone   string `json:"phone"`
		Message string `json:"message"`
	}

	var info MessageContents
	data.Bind(&info)

	mailErr := contactMe.SendEmail(info.Name, info.Email, info.Phone, info.Message)
	if mailErr != nil {
		data.JSON(http.StatusBadRequest, mailErr.Error())
	} else {
		data.JSON(http.StatusOK, nil)
	}
}

func GetOccurrences(data *gin.Context) {
	type SearchString struct {
		SplitString []string `json:"splitString"`
	}

	var info SearchString
	data.Bind(&info)

	fmt.Println("STRING 1: ", info.SplitString)

	stringOccurrenceReturn, err := goExamples.GetStringOccurrences(info.SplitString)
	if err != nil {
		data.JSON(http.StatusBadRequest, err.Error())
	} else {
		data.JSON(http.StatusOK, stringOccurrenceReturn)
	}
}

// Translate will take a string and a language and parse it to the specified language
func Translate(data *gin.Context) {
	type TranslateString struct {
		SplitString string `json:"splitString"`
		Lang        string `json:"lng"`
	}

	var info TranslateString
	data.Bind(&info)

	fmt.Println("String Before: ", info.SplitString)
	fmt.Println("Lang Before: ", info.Lang)

	translationReturn, err := goExamples.TranslateString(info.SplitString, info.Lang)
	if err != nil {
		data.JSON(http.StatusBadRequest, err.Error())
	} else {
		data.JSON(http.StatusOK, translationReturn)
	}
}

// PostTweet will post a string to twitter
func PostTweet(data *gin.Context) {
	type Tweet struct {
		Tweet string `json:"tweet"`
	}

	var info Tweet
	data.Bind(&info)

	fmt.Println("INFO: ", info.Tweet)

	SubmitTweet, err := goExamples.SubmitTweet(info.Tweet)
	if err != nil {
		data.JSON(http.StatusBadRequest, err.Error())
	} else {
		data.JSON(http.StatusOK, SubmitTweet)
	}
}