package starWarsCharacterTableEample

import (
	"fmt"

	"PortfolioWebsite/goResources/db"
)

func AddCharacter(name string, homeworld string, born string, died string, gender string, species string, affiliation []string , associated []string , masters []string, apprentices []string)  (int, error) {
	// Used to return the ID of the row we inserted into our DB.
	lastInsertId := 0

	//fmt.Println("Name: ", name)
	//fmt.Println("Born: ", born)
	//fmt.Println("Associated: ", associated)
	//fmt.Println("Gender: ", gender)
	//fmt.Println("Affiliation: ", affiliation)
	//fmt.Println("Masters: ", masters)

	// This inserts our quote and accompanying data into our table!
	err := db.ConnPool.QueryRow(
		"INSERT INTO star_wars_characters(name, home_world, born, died, species, gender, affiliation, associated, masters, apprentices) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id",
		name, homeworld, born, died, species, gender, affiliation, associated, masters, apprentices).Scan(&lastInsertId)
	if err != nil {
		fmt.Println("Error saving character to database: ", err)
		return 0, err
	}

	return lastInsertId, nil
}
