package visitorCounter

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"PortfolioWebsite/src/go/common"
	"PortfolioWebsite/src/go/db"
)

type IpResult struct {
	Id        int       `json:"id"`
	Ip        string    `json:"ip"`
	Timestamp time.Time `json:"timestamp"`
}

type VisitorResult struct {
	Id           int       `json:"id"`
	Month        string    `json:"month"`
	Count        int       `json:"count"`
	Year         int       `json:"year"`
	PageCount    int       `json:"pageCount"`
	DateStamp    time.Time `json:"date_stamp"`
	AvgCount     float64   `json:"avgCount"`
	AvgPageCount float64   `json:"avgPageCount"`
}

type WhoIsAPI struct {
	Result WhoIsAPIResult `json:"result"`
}

type WhoIsAPIResult struct {
	Address        string   `json:"address"`
	City           string   `json:"city"`
	Country        string   `json:"country"`
	CreationDate   string   `json:"creation_date"`
	DNSSec         string   `json:"dnssec"`
	DomainName     string   `json:"domain_name"`
	Emails         string   `json:"emails"`
	ExpirationDate string   `json:"expiration_date"`
	Name           string   `json:"name"`
	NameServers    []string `json:"name_servers"`
	Org            string   `json:"org"`
	ReferralURL    string   `json:"referral_url"`
	Registrar      string   `json:"registrar"`
	State          string   `json:"state"`
	Status         string   `json:"status"`
	UpdatedDate    string   `json:"updated_date"`
	WhoIsServer    string   `json:"whois_server"`
	ZipCode        string   `json:"zipcode"`
}

type VisitorLocationResult struct {
	As        VisitorLocationAs       `json:"as"`
	Ip        string                  `json:"ip"`
	Isp       string                  `json:"isp"`
	Location  VisitorLocationLocation `json:"location"`
	Timestamp time.Time               `json:"timestamp"`
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
		emptyIPlocationsReturn, err := EmptyIPLocations()
		if err != nil {
			fmt.Println("Error clearing the ip_locations table!", err)
		} else {
			fmt.Println("Successfully emptied IP Locations in DB", emptyIPlocationsReturn)
		}
	}
}

func EmptyVisitors() (string, error) {
	commandTag, err := db.ConnPool.Exec(context.Background(),
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
	commandTag, err := db.ConnPool.Exec(context.Background(),
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
	y, m, _ := current.Date()

	// First lets check to make sure that the month we're looking to add isn't already in the table
	var exists bool
	err := db.ConnPool.QueryRow(context.Background(),
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
	// If the row doesn't exist, we're going to create it!
	if !exists {
		// This inserts our quote and accompanying data into our table!
		err = db.ConnPool.QueryRow(context.Background(),
			`INSERT INTO
				monthly_visitors(
					month,
					count,
					year,
					page_count)
			VALUES(
				$1, $2, $3, $4)
			RETURNING
				id`,
			m.String(), 0, y, 0).Scan(&lastInsertId)
		if err != nil {
			fmt.Println("Error saving new month to database: ", err)
			return "", err
		}

		return fmt.Sprintf("Added Month: %s to DB!!", m.String()), nil
	} else {
		return "Row already in table!", nil
	}
}

func IncrementMonthlyPageCount() (string, error) {
	current := time.Now().UTC()
	_, m, _ := current.Date()
	message, err := db.ConnPool.Exec(context.Background(),
		`UPDATE
				monthly_visitors
			SET
				page_count = page_count + 1
			WHERE
				month = $1`,
		m.String())
	if err != nil {
		fmt.Println("There was an error updating the monthly visitors table in the database 1:", err)
		return "", err
	} else {
		fmt.Println("Success in updating the page_count!", message, m.String())
		return "SUCCESS!", nil
	}
}

func IncrementMonthlyVisitors() (string, error) {
	current := time.Now().UTC()
	_, m, _ := current.Date()
	_, err := db.ConnPool.Exec(context.Background(),
		`UPDATE
				monthly_visitors
			SET
				count = count + 1, page_count = page_count + 1
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

func CheckIfIPExists(ip string, domain string) (string, error) {
	unique := true

	fmt.Println("IP: ", ip)
	fmt.Println("Domain: ", domain)

	rows, err := db.ConnPool.Query(context.Background(),
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

	counter := 0

	for rows.Next() {
		dbID := 0
		dbIP := ""
		dbTimestamp := time.Now()
		counter++
		err := rows.Scan(&dbID, &dbIP, &dbTimestamp)
		if err != nil {
			fmt.Println("Error scanning row: ", err)
			return "", err
		}
		if dbIP == ip {
			fmt.Println("IP NOT UNIQUE: ", dbIP, " ", ip)
			unique = false
		}
	}

	if unique {
		message, err := WriteIPToDatabase(ip, domain)
		if err != nil {
			fmt.Println("Error inserting IP to DB", message)
			return message, err
		}
		updateIPLocation, err := WriteIPLocationToDB(ip, domain)
		if err != nil {
			fmt.Println("Error updating the IP Location to DB", updateIPLocation)
		}
	} else if !unique {
		// Not unique, but we want to increment page_count!
		pageCountReturn, err := IncrementMonthlyPageCount()
		if err != nil {
			fmt.Println("Error inserting IP into DB", pageCountReturn)
			return pageCountReturn, err
		}
	}
	return "Not Unique", nil
}

func WriteIPToDatabase(ip string, domain string) (string, error) {
	lastInsertId := 0
	now := time.Now()
	// This inserts our quote and accompanying data into our table!
	err := db.ConnPool.QueryRow(context.Background(),
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
		fmt.Println("Error saving IP to database 1: ", err)
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
	// Getting a rolling AVG from this URL:
	// https://www.compose.com/articles/metrics-maven-calculating-a-moving-average-in-postgresql/
	// Should be calculating a Rolling AVG from the previous 2 entries and the current entry.
	rows, err := db.ConnPool.Query(context.Background(),
		`SELECT
				id,
				month,
				count,
				year,
				page_count,
				date_stamp,
				AVG(count)
				OVER(ORDER BY id ASC ROWS BETWEEN 2 PRECEDING AND CURRENT ROW) AS avg_count,
				AVG(page_count)
            	OVER(ORDER BY id ASC ROWS BETWEEN 2 PRECEDING AND CURRENT ROW) AS avg_page_count
			FROM
				monthly_visitors
			ORDER BY
			   id ASC`)

	if err != nil {
		fmt.Println("There was an error reading the monthly_visitors table from the database 2:", err)
		return nil, err
	}

	monthlyVisitorsResultsArray := []*VisitorResult{}

	defer rows.Close()

	for rows.Next() {

		var res VisitorResult

		err = rows.Scan(
			&res.Id,
			&res.Month,
			&res.Count,
			&res.Year,
			&res.PageCount,
			&res.DateStamp,
			&res.AvgCount,
			&res.AvgPageCount)

		if err != nil {
			fmt.Println("There was an error querying that database for the Monthly Visitors Results:", err)
			continue
		}

		monthlyVisitorsResultsArray = append(monthlyVisitorsResultsArray, &res)
	}

	return monthlyVisitorsResultsArray, nil
}

func ReadIPDB() ([]*IpResult, error) {

	rows, err := db.ConnPool.Query(context.Background(),
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

	rows, err := db.ConnPool.Query(context.Background(),
		`SELECT
				ip_locations.as_asn,
				ip_locations.as_domain,
				ip_locations.as_name,
				ip_locations.as_route,
				ip_locations.ip,
				ip_locations.location_city,
				ip_locations.location_country,
				ip_locations.location_lat,
				ip_locations.location_lng,
				ip_locations.location_postalcode,
				ip_locations.location_region,
				ip_locations.location_timezone,
				ip_locations.isp,
				ips.timestamp
			FROM
				ip_locations, ips
            WHERE
                ip_locations.ip = ips.ip`)

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
			&res.Isp,
			&res.Timestamp)

		if err != nil {
			fmt.Println("There was an error querying that database for the IP Location Results:", err)
			continue
		}

		ipLocationResultsArray = append(ipLocationResultsArray, &res)
	}

	return ipLocationResultsArray, nil
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

func WriteIPLocationToDB(ip string, domain string) (string, error) {
	lastInsertId := 0

	whoIsKeyInfo, err := common.ReadJsonFile("whoIsApiKey.json")
	if err != nil {
		fmt.Println("Error reading JSON file!")
		return "", err
	}
	//whoIsHost := whoIsKeyInfo["host"].(string)
	whoIsKey := whoIsKeyInfo["key"].(string)

	url := "https://api.apilayer.com/whois/query?domain=" + domain

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("apikey", whoIsKey)

	if err != nil {
		fmt.Println("WHOISAPI CALL ERROR: ", err)
	}
	res, err := client.Do(req)
	if res.Body != nil {
		defer res.Body.Close()
	}
	ipLocationReturn, err := ioutil.ReadAll(res.Body)

	// Putting our IP Location information to a struct
	var ipLocationResult WhoIsAPI
	unmarshalErr := json.Unmarshal([]byte(ipLocationReturn), &ipLocationResult)
	// Lets check if it is because we're on LocalHost. If so, lets create a blank entry.
	if unmarshalErr != nil && !strings.Contains(ip, "localhost") {
		fmt.Println("Error unmarshaling IP Location JSON", unmarshalErr)
		return "", unmarshalErr
	} else if strings.Contains(ip, "localhost") {
		// Lets create an empty struct
		ipLocationResult.Result.DomainName = "LocalHost"
		ipLocationResult.Result.Registrar = "LocalHost"
		ipLocationResult.Result.City = "West Haven"
		ipLocationResult.Result.Country = "USA"
		ipLocationResult.Result.State = "Utah"
		ipLocationResult.Result.ZipCode = "84401"
		ipLocationResult.Result.Org = "Century Link"
	}

	// This inserts our quote and accompanying data into our table!
	dbErr := db.ConnPool.QueryRow(context.Background(),
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
		1,
		ipLocationResult.Result.DomainName,
		ipLocationResult.Result.Registrar,
		ip,
		ip,
		ipLocationResult.Result.City,
		ipLocationResult.Result.Country,
		41.2705,
		72.9470,
		ipLocationResult.Result.ZipCode,
		ipLocationResult.Result.State,
		"",
		ipLocationResult.Result.Org).Scan(&lastInsertId)
	if dbErr != nil {
		fmt.Println("Error saving IP to database 2: ", err)
		return "", err
	}

	return "Success writing IP location to DB!", nil
}
