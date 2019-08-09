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
	Ip        string    `json:"ip"`
	Timestamp time.Time `json:"timestamp"`
}

type VisitorResult struct {
	Id    int    `json:"id"`
	Month string `json:"month"`
	Count int    `json:"count"`
}

type VisitorLocationResult struct {
	As       VisitorLocationAs       `json:"as"`
	Ip       string                  `json:"ip"`
	Isp      string                  `json:"isp"`
	Location VisitorLocationLocation `json:"location"`
}

type VisitorLocationAs struct {
	Asn    int    `json:"asn"`
	Domain string `json:"domain"`
	Name   string `json:"name"`
	Route  string `json:"route"`
}

type VisitorLocationLocation struct {
	City       string  `json:"city"`
	Country    string  `json:"country"`
	Lat        float64 `json:"lat"`
	Lng        float64 `json:"lng"`
	PostalCode string  `json:"postalcode"`
	Region     string  `json:"region"`
	Timezone   string  `json:"timezone"`
}

type VisitorLocationDB struct {
	Id                 int     `json:"id"`
	AsAsn              int     `json:"as_asn"`
	AsDomain           string  `json:"as_domain"`
	AsName             string  `json:"as_name"`
	AsRoute            string  `json:"as_route"`
	Ip                 string  `json:"ip"`
	LocationCity       string  `json:"location_city"`
	LocationCountry    string  `json:"location_country"`
	LocationLat        float64 `json:"location_lat"`
	LocationLng        float64 `json:"location_lng"`
	LocationPostalcode string  `json:"location_postalcode"`
	LocationRegion     string  `json:"location_region"`
	LocationTimezone   string  `json:"location_timezone"`
}


func InitCreateMonth() {
	for {
		now := time.Now()
		// We rest the counter once a day, at midnight. Well, midnight for our server which depends on the time zone.
		next := time.Date(now.Year(), now.Month()+1, 1, 3, 0, 0, 0, now.Location())

		sleepDur := next.Sub(now)
		fmt.Printf("Creating Month in DB in %s on %s\n", sleepDur.String(), next)
		time.Sleep(sleepDur)

		createMonthReturn, err := CreateMonth()
		if err != nil {
			fmt.Println("Error creating new month!")
		} else {
			fmt.Println("Successfully created month in DB", createMonthReturn)
		}
		emptyVisitorsReturn, err := EmptyVisitors()
		if err != nil {
			fmt.Println("Error clearing the ips table!", err)
		} else {
			fmt.Println("Successfully emptied visitors in DB", emptyVisitorsReturn)
		}
		emptyIPlocationsReturn ,err := EmptyIPLocations()
		if err != nil {
			fmt.Println("Error clearing the ip_locations table!", err)
		} else {
			fmt.Println("Successfully emptied IP Locations in DB", emptyIPlocationsReturn)
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
	} else {
		return "Successfully deleted the ips table!", nil
	}
}

func EmptyIPLocations() (string, error) {
	commandTag, err := db.ConnPool.Exec(
		`DELETE FROM
				ip_locations
			RETURNING *`)
	if err != nil {
		return "", err
	}
	if commandTag.RowsAffected() != 1 {
		return "", fmt.Errorf("No row found to delete")
	} else {
		return "Successfully deleted the ip_locations table!", nil
	}
}

func CreateMonth() (string, error) {
	lastInsertId := 0
	current := time.Now().UTC()
	y, m, _:= current.Date()

	// First lets check to make sure that the month we're looking to add isn't already in the table
	var exists bool
	err := db.ConnPool.QueryRow(
		`SELECT EXISTS(
				SELECT
					month
				FROM
					monthly_visitors
				WHERE
					month = $1
					AND 
					year = $2)`,
		m.String(), y).Scan(&exists)
	fmt.Println("EXISTS: ", exists)
	// If the row doesn't exist, we're going to create it!
	if !exists {
		// This inserts our quote and accompanying data into our table!
		err = db.ConnPool.QueryRow(
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
	} else {
		return "Row already in table!", nil
	}
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
		updateIPLocation, err := WriteIPLocationToDB(ip)
		if err != nil {
			fmt.Println("Error updating the IP Location to DB", updateIPLocation)
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

func ReadIPLocationDB() ([]*VisitorLocationResult, error) {

	rows, err := db.ConnPool.Query(
		`SELECT
				as_asn,
				as_domain,
				as_name,
				as_route,
				ip,
				location_city,
				location_country,
				location_lat,
				location_lng,
				location_postalcode,
				location_region,
				location_timezone,
				isp
			FROM
				ip_locations`)

	if err != nil {
		fmt.Println("There was an error reading the ip_locations table from the database:", err)
		return nil, err
	}

	ipLocationResultsArray := []*VisitorLocationResult{}

	defer rows.Close()

	for rows.Next() {

		var res VisitorLocationResult

		err = rows.Scan(
			&res.As.Asn,
			&res.As.Domain,
			&res.As.Name,
			&res.As.Route,
			&res.Ip,
			&res.Location.City,
			&res.Location.Country,
			&res.Location.Lat,
			&res.Location.Lng,
			&res.Location.PostalCode,
			&res.Location.Region,
			&res.Location.Timezone,
			&res.Isp)

		if err != nil {
			fmt.Println("There was an error querying that database for the IP Location Results:", err)
			continue
		}

		ipLocationResultsArray = append(ipLocationResultsArray, &res)
	}

	return ipLocationResultsArray, nil
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

func GetIPLocation() (map[string]interface{}, error) {
	b, err := ioutil.ReadFile("goResources/visitorCounter/ipLocations.json") // just pass the file name
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	var dat map[string]interface{}
	err = json.Unmarshal(b, &dat)
	if err != nil {
		fmt.Println("Error unmarshaling IP Locations JSON", err)
	}

	return dat, nil
}

// Functions for our IP Locations

func WriteIPLocationToDB(ip string) (string, error) {
	lastInsertId := 0

	ipLocationReturn, err := GetIPLocationRequest(ip)
	if err != nil {
		fmt.Println("Error obtaining IP Location", err)
		return "", err
	}
	fmt.Println("IP LOCATION BEFORE: ", ipLocationReturn)

	ipLocationReturnJSON, err := json.Marshal(ipLocationReturn)
	if err != nil {
		fmt.Println("Error marshaling IP Location", err)
		return "", err
	}
	//ipLocationReturnJSONString := string(ipLocationReturnJSON)
	//fmt.Println("IP LOCATION AFTER: ", string(ipLocationReturnJSON))

	// Putting our IP Location information to a struct
	var ipLocationResult VisitorLocationResult
	unmarshalErr := json.Unmarshal(ipLocationReturnJSON, &ipLocationResult)
	if unmarshalErr != nil {
		fmt.Println("Error unmarshaling IP Location JSON", unmarshalErr)
		return "", unmarshalErr
	}

	// This inserts our quote and accompanying data into our table!
	dbErr := db.ConnPool.QueryRow(
		`INSERT INTO 
				ip_locations(
					as_asn,
					as_domain,
					as_name,
					as_route,
					ip,
					location_city,
					location_country,
					location_lat,
					location_lng,
					location_postalcode,
					location_region,
					location_timezone,
					isp) 
			VALUES(
				$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) 
			RETURNING 
				id`,
		ipLocationResult.As.Asn,
		ipLocationResult.As.Domain,
		ipLocationResult.As.Name,
		ipLocationResult.As.Route,
		ipLocationResult.Ip,
		ipLocationResult.Location.City,
		ipLocationResult.Location.Country,
		ipLocationResult.Location.Lat,
		ipLocationResult.Location.Lng,
		ipLocationResult.Location.PostalCode,
		ipLocationResult.Location.Region,
		ipLocationResult.Location.Timezone,
		ipLocationResult.Isp).Scan(&lastInsertId)
	if dbErr != nil {
		fmt.Println("Error saving IP to database: ", err)
		return "", err
	}


	return "Success writing IP location to DB!", nil
}
