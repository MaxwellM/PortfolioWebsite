package visitorCounter

import (
	"PortfolioWebsite/goResources/db"
	"fmt"
	"time"
)

type IpResult struct {
	Id        int       `json:"id"`
	Ip        string    `json:'ip"`
	Timestamp time.Time `json:"timestamp"`
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
			return "", err
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

	fmt.Println("LAST INSERT ID: ", lastInsertId)

	return "NEW IP, ADDED TO DB!", nil
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
