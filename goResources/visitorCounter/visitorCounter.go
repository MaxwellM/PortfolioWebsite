package visitorCounter

import (
	"PortfolioWebsite/goResources/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type IpResult struct {
	Id        int       `json:"id"`
	Ip        string    `json:'ip"`
	Timestamp time.Time `json:"timestamp"`
}

type VisitorResult struct {
	Id int `json:"id"`
	Month string `json:"month"`
	Count int `json:"count"`
}

//{
//	"ip": "110.170.64.0",
//	"location": {
//	"country": "TH",
//	"region": "Bangkok",
//	"city": "Bangkok",
//	"lat": 13.7083,
//	"lng": 100.4562,
//	"postalCode": "10120",
//	"timezone": "Asia/Bangkok"
//},
//	"isp": "True Internet",
//	"domains": [
//		"110-170-64-0.static.asianet.co.th"
//	],
//	"as": {
//		"asn": 7470,
//		"name": "TRUE INTERNET Co.,Ltd.",
//		"route": "110.170.64.0/18",
//		"domain": "trueinternet.co.th"
//	}
//}

type VisitorLocation struct {
	VisitorLocationIPInfo
}

type VisitorLocationIPInfo struct {
	IP string `json:"ip"`
	Location VisitorLocationLocationInfo `json:"location"`
	ISP string `json:"isp"`
	Domains []VisitorLocationDomainsInfo `json:"domains"`
	AS VisitorLocationASInfo `json:"as"`
}

type VisitorLocationLocationInfo struct {
	County string `json:"county"`
	Region string `json:"region"`
	City string `json:"city"`
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
	PostalCode string `json:"postalCode"`
	Timezone string `json:"timezone"`
}

type VisitorLocationDomainsInfo struct {
	Domain string `json:"domain"`
}

type VisitorLocationASInfo struct {
	ASN int `json:"asn"`
	Name string `json:"name"`
	Route string `json:"route"`
	Domain string `json:"domain"`
}

func InitCreateMonth() {
	for {
		now := time.Now()
		// We rest the counter once a day, at midnight. Well, midnight for our server which depends on the time zone.
		next := time.Date(now.Year(), now.Month()+1, now.Day(), 1, 0, 0, 0, now.Location())

		sleepDur := next.Sub(now)
		fmt.Printf("Creating Month in DB in %s on %s\n", sleepDur.String(), next)
		time.Sleep(sleepDur)

		createMonthReturn, err := CreateMonth()
		if err != nil {
			fmt.Println("Error creating new month!")
		} else {
			fmt.Println(createMonthReturn)
		}
		emptyVisitorsReturn, err := EmptyVisitors()
		if err != nil {
			fmt.Println("Error clearing the ips table!", err)
		} else {
			fmt.Println(emptyVisitorsReturn)
		}
	}
}

func EmptyVisitors() (string, error) {
	commandTag, err := db.ConnPool.Exec(
		`DELETE FROM 
				ips`)
	if err != nil {
		return "", err
	}
	if commandTag.RowsAffected() != 1 {
		return "", fmt.Errorf("No row found to delete")
	}
	return "Success!", nil
}

func CreateMonth() (string, error) {
	lastInsertId := 0
	current := time.Now().UTC()
	y, m, _:= current.Date()
	// This inserts our quote and accompanying data into our table!
	err := db.ConnPool.QueryRow(
		`INSERT INTO 
				monthly_visitors(
					month,
					count,
					year) 
			VALUES(
				$1, $2, $3) 
			RETURNING 
				id`,
		m.String(), 0, y).Scan(&lastInsertId)
	if err != nil {
		fmt.Println("Error saving new month to database: ", err)
		return "", err
	}

	fmt.Println("LAST INSERT ID: ", lastInsertId)

	return fmt.Sprintf("Added Month: %s to DB!!", m.String()), nil
}

func IncrementMonthlyVisitors() (string, error) {
	current := time.Now().UTC()
	_, m, _ := current.Date()
	_, err := db.ConnPool.Exec(
		`UPDATE 
				monthly_visitors
			SET 
				count = count + 1 
			WHERE 
				month = $1`,
		m.String())
	if err != nil {
		fmt.Println("There was an error updating the monthly visitors table in the database 1:", err)
		return "", err
	} else {
		return "SUCCESS!", nil
	}
}

func CheckIfIPExists(ip string) (string, error) {
	unique := true

	rows, err := db.ConnPool.Query(
		`SELECT
				id,
				ip,
				timestamp
			FROM
				ips`)
	if err != nil {
		fmt.Println("There was an error reading the ips table from the database 1:", err)
		return "", err
	}

	counter := 0;

	for rows.Next() {
		dbID := 0
		dbIP := ""
		dbTimestamp := time.Now()
		counter ++
		err := rows.Scan(&dbID, &dbIP, &dbTimestamp)
		if err != nil {
			fmt.Println("Error scanning row: ", err)
			return "", err
		}
		fmt.Println("DBIP: ", dbIP)
		fmt.Println("IP: ", ip)
		if dbIP == ip {
			unique = false
		}
	}

	fmt.Println("COUNTER: ", counter)

	if unique {
		fmt.Println("IP UNIQUE!")
		message, err := WriteIPToDatabase(ip)
		if err != nil {
			fmt.Println("Error inserting IP to DB", message)
			return message, err
		}
	}
	fmt.Println("IP NOT UNIQUE!")
	return "Not Unique", nil
}

func WriteIPToDatabase(ip string) (string, error) {
	lastInsertId := 0
	now := time.Now()
	// This inserts our quote and accompanying data into our table!
	err := db.ConnPool.QueryRow(
		`INSERT INTO 
				ips(
					ip,
					timestamp) 
			VALUES(
				$1, $2) 
			RETURNING 
				id`,
		ip, now).Scan(&lastInsertId)
	if err != nil {
		fmt.Println("Error saving IP to database: ", err)
		return "", err
	}

	// Now we're going to increment our Monthly Visitors DB!
	monthlyVisitorsReturn, err := IncrementMonthlyVisitors()
	if err != nil {
		return monthlyVisitorsReturn, err
	} else {
		fmt.Println("LAST INSERT ID: ", lastInsertId)

		return "NEW IP, ADDED TO DB!", nil
	}
}

func ReadMonthlyVisitorsDB() ([]*VisitorResult, error) {
	rows, err := db.ConnPool.Query(
		`SELECT
				id,
				month,
				count
			FROM
				monthly_visitors`)

	if err != nil {
		fmt.Println("There was an error reading the ips table from the database 2:", err)
		return nil, err
	}

	monthlyVisitorsResultsArray := []*VisitorResult{}

	defer rows.Close()

	for rows.Next() {

		var res VisitorResult

		err = rows.Scan(
			&res.Id,
			&res.Month,
			&res.Count)

		if err != nil {
			fmt.Println("There was an error querying that database for the Monthly Visitors Results:", err)
			continue
		}

		monthlyVisitorsResultsArray = append(monthlyVisitorsResultsArray, &res)
	}

	return monthlyVisitorsResultsArray, nil
}

func ReadIPDB() ([]*IpResult, error) {

	rows, err := db.ConnPool.Query(
		`SELECT
				id,
				ip,
				timestamp
			FROM
				ips`)

	if err != nil {
		fmt.Println("There was an error reading the ips table from the database 2:", err)
		return nil, err
	}

	ipResultsArray := []*IpResult{}

	defer rows.Close()

	for rows.Next() {

		var res IpResult

		err = rows.Scan(
			&res.Id,
			&res.Ip,
			&res.Timestamp)

		if err != nil {
			fmt.Println("There was an error querying that database for the IP Results:", err)
			continue
		}

		ipResultsArray = append(ipResultsArray, &res)
	}

	return ipResultsArray, nil

}

func GetIPLocationRequest(ip string) (map[string]interface{}, error) {
		fmt.Println("LOCATION IP: ", ip)
		//key := "at_pruWCmEUi97TIwBtqGswfJokFFZ6M"
		url := fmt.Sprintf(`https://geoipify.whoisxmlapi.com/api/v1?apiKey=at_pruWCmEUi97TIwBtqGswfJokFFZ6M&ipAddress=`+ip)

		//fmt.Println("URL: ", url)

		resp, err := http.Get(url)

		fmt.Println("IP LOCATION RESP: ", resp.Body)

		if err != nil {
			fmt.Println("There was an error getting the IP Location... ", err)
			return nil, err
		}

		defer resp.Body.Close()

		// This worked because we need to convert our *Reader to []Bytes
		// https://stackoverflow.com/questions/38673673/access-http-response-as-string-in-go
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Couldn't convert RESP to []Byte for your Current Conditions Request: ", err)
		}

		var test map[string]interface{}
		err = json.Unmarshal(bytes, &test)
		if err != nil {
			return nil, err
		}

		fmt.Println("TEST: ", test)

		//finalResult := map[string]interface{}{
		//	"Result":          test,
		//	"Location":        location,
		//}
		//
		//fmt.Println("CURRENT CONDITIONS: ", finalResult)

		return test, err
}

func GetIPLocation(ip string) (map[string]interface{}, error) {
	ipLocationReturn, err := GetIPLocationRequest(ip)
	if err != nil {
		fmt.Println("Error obtaining IP location!", err)
		return nil, err
	}
	ipLocationReturnJSON, err := json.Marshal(ipLocationReturn)
	if err != nil {
		fmt.Println("Error writing weather to file", err)
		return nil, err
	}
	currentConditionsReturnJSONReturn := ioutil.WriteFile("goResources/visitorCounter/ipLocations.json", ipLocationReturnJSON, 0644)
	fmt.Println(currentConditionsReturnJSONReturn)
	// Now we return our weather report back to the frontend
	allWeather := map[string]interface{}{
		"Current": ipLocationReturn,
	}

	return allWeather, nil
}
