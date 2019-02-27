package weatherExample

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"
)

var requestCount = 0

func CountRequest() bool {
	requestCount++
	fmt.Println("Request NUM: ", requestCount)
	if requestCount < 50 {
		return true
	} else {
		fmt.Println("TOO MANY!")
		// Too many requests today
		return false
	}
}

func resetRequestCount() {
	requestCount = 0
}

func ReadLocalCurrentConditions() (map[string]interface{}, error) {
	b, err := ioutil.ReadFile("goResources/weatherExample/weatherReport/currentConditions.json") // just pass the file name
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	var dat map[string]interface{}
	err = json.Unmarshal(b, &dat)
	if err != nil {
		fmt.Println("Error unmarshaling Weather Report JSON", err)
	}

	return dat, nil
}

func ReadLocalWeatherReport() (map[string]interface{}, error) {
	b, err := ioutil.ReadFile("goResources/weatherExample/weatherReport/weather.json") // just pass the file name
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	var dat map[string]interface{}
	err = json.Unmarshal(b, &dat)
	if err != nil {
		fmt.Println("Error unmarshaling Weather Report JSON", err)
	}

	return dat, nil
}

func UpdateCurrentConditions(location string) (map[string]interface{}, error) {
	currentConditionsReturn, err := GetCurrentConditions(location)
	if err != nil {
		fmt.Println("Error obtaining current conditions!", err)
		return nil, err
	}
	currentConditionsReturnJSON, err := json.Marshal(currentConditionsReturn)
	if err != nil {
		fmt.Println("Error writing weather to file", err)
		return nil, err
	}
	currentConditionsReturnJSONReturn := ioutil.WriteFile("goResources/weatherExample/weatherReport/currentConditions.json", currentConditionsReturnJSON, 0644)
	fmt.Println(currentConditionsReturnJSONReturn)
	// Now we return our weather report back to the frontend
	allWeather := map[string]interface{}{
		"Current": currentConditionsReturn,
	}

	return allWeather, nil
}

func UpdateForecast(location string) (map[string]interface{}, error) {
	weatherReturn, err := GetWeather(location)
	if err != nil {
		fmt.Println("Error obtaining weather!")
		return nil, err
	}
	weatherReturnJSON, err := json.Marshal(weatherReturn)
	if err != nil {
		fmt.Println("Error writing weather to file", err)
		return nil, err
	}
	weatherReturnJSONReturn := ioutil.WriteFile("goResources/weatherExample/weatherReport/weather.json", weatherReturnJSON, 0644)
	fmt.Println(weatherReturnJSONReturn)
	// Now we return our weather report back to the frontend
	allWeather := map[string]interface{}{
		"Forecast": weatherReturn,
	}

	return allWeather, nil
}


func UpdateAllWeather(location string) (map[string]interface{}, error){
		weatherReturn, err := GetWeather(location)
		if err != nil {
			fmt.Println("Error obtaining weather!")
			return nil, err
		}
		//fmt.Println(weatherReturn)
		fmt.Println(reflect.TypeOf(weatherReturn))

		weatherReturnJSON, err := json.Marshal(weatherReturn)
		if err != nil {
			fmt.Println("Error writing weather to file", err)
			return nil, err
		}
		weatherReturnJSONReturn := ioutil.WriteFile("goResources/weatherExample/weatherReport/weather.json", weatherReturnJSON, 0644)
		fmt.Println(weatherReturnJSONReturn)

		currentConditionsReturn, err := GetCurrentConditions(location)
		if err != nil {
			fmt.Println("Error obtaining current conditions!", err)
			return nil, err
		}
		currentConditionsReturnJSON, err := json.Marshal(currentConditionsReturn)
		currentConditionsReturnJSONReturn := ioutil.WriteFile("goResources/weatherExample/weatherReport/currentConditions.json", currentConditionsReturnJSON, 0644)
		fmt.Println(currentConditionsReturnJSONReturn)

		// Now we return our weather report back to the frontend
		allWeather := map[string]interface{}{
			"Forecast": weatherReturn,
			"Current": currentConditionsReturn,
		}

		return allWeather, nil
}


func InitRequestCount() {
	for {
		now := time.Now()
		// We rest the counter once a day, at midnight. Well, midnight for our server which depends on the time zone.
		next := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())

		sleepDur := next.Sub(now)
		fmt.Printf("Reseting Weather Request Counter in %s on %s\n", sleepDur.String(), next)
		time.Sleep(sleepDur)

		resetRequestCount()
	}
}

func InitUpdateForecast() {
	for {
		now := time.Now()
		next := time.Date(now.Year(), now.Month(), now.Day()+1, 12, 0, 0, 0, now.Location())

		sleepDur := next.Sub(now)
		fmt.Printf("Updating Forecast in %s on %s\n", sleepDur.String(), next)
		time.Sleep(sleepDur)
		fmt.Println("UPDATING THE FORECAST!")
		_, err := UpdateForecast("351219")
		if err != nil {
			fmt.Println("Error fetching spotter validator stuff")
		}
	}
}

func InitUpdateCurrentConditions() {
	for {
		t := time.Now()
		if t.Minute() == 00 {
			fmt.Println("UPDATING THE WEATHER!")
			_, err := UpdateCurrentConditions("351219")
			if err != nil {
				fmt.Println("Error updating the weather report: ", err)
			}
		}
		time.Sleep(1 * time.Minute)
	}
}

func GetCurrentConditions(location string) (map[string]interface{}, error) {
	if CountRequest() {
		key := "6kXpxa4RNqkTAgbRNc4ZFaZvcCOLcrM3"
		url := fmt.Sprintf(`http://dataservice.accuweather.com/currentconditions/v1/`+location+`?apikey=`+key)

		//fmt.Println("URL: ", url)

		resp, err := http.Get(url)

		fmt.Println("CURRENT CONDITIONS RESP: ", resp.Body)

		if err != nil {
			fmt.Println("There was an error getting the Weather Report you were looking for... ", err)
			return nil, err
		}

		defer resp.Body.Close()

		// This worked because we need to convert our *Reader to []Bytes
		// https://stackoverflow.com/questions/38673673/access-http-response-as-string-in-go
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Couldn't convert RESP to []Byte for your Current Conditions Request: ", err)
		}

		var test []map[string]interface{}
		err = json.Unmarshal(bytes, &test)
		if err != nil {
			return nil, err
		}

		finalResult := map[string]interface{}{
			"Result":          test,
			"Location":        location,
		}

		fmt.Println("CURRENT CONDITIONS: ", finalResult)

		return finalResult, err
	} else {
		return nil, fmt.Errorf("Reached reached request limit for the day. Come back tomorrow when it resets!")
	}
}

func GetWeather(location string) (map[string]interface{}, error) {
	if CountRequest() {
		key := "6kXpxa4RNqkTAgbRNc4ZFaZvcCOLcrM3"
		url := fmt.Sprintf(`http://dataservice.accuweather.com/forecasts/v1/daily/5day/`+location+`?apikey=`+key)

		//fmt.Println("URL: ", url)

		resp, err := http.Get(url)

		//fmt.Println("RESP: ", resp)

		if err != nil {
			fmt.Println("There was an error getting the Weather Report you were looking for... ", err)
			return nil, err
		}

		defer resp.Body.Close()

		// This worked because we need to convert our *Reader to []Bytes
		// https://stackoverflow.com/questions/38673673/access-http-response-as-string-in-go
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Couldn't convert RESP to []Byte for your Weather Request: ", err)
		}

		var test map[string]interface{}
		err = json.Unmarshal(bytes, &test)
		if err != nil {
			return nil, err
		}

		//fmt.Println("WEATHER: ", test)

		finalResult := map[string]interface{}{
			"Result":          test,
			"Location":        location,
		}

		return finalResult, err
	} else {
		return nil, fmt.Errorf("Reached reached request limit for the day. Come back tomorrow when it resets!")
	}
}
