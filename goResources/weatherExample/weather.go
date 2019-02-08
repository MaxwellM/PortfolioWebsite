package weatherExample

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

func InitRequstCount() {
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

func GetWeather(location string) (map[string]interface{}, error) {
	if CountRequest() {
		key := "6kXpxa4RNqkTAgbRNc4ZFaZvcCOLcrM3"
		url := fmt.Sprintf(`http://dataservice.accuweather.com/forecasts/v1/daily/5day/`+location+`?apikey=`+key)

		fmt.Println("URL: ", url)

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

		return test, err
	} else {
		return nil, fmt.Errorf("Reached reached request limit for the day. Come back tomorrow when it resets!")
	}
}
