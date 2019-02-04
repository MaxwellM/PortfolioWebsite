package starWarsCharacterTableEample

import (
	"PortfolioWebsite/goResources/db"
	"fmt"
)

type StarWarsCharacterResult struct {
	Id          int                    `json:"id"`
	Name        string                 `json:"name"`
	Homeworld   string                 `json:"homeworld"`
	Born        string                 `json:"born"`
	Died        string                 `json:"died"`
	Species     string                 `json:"species"`
	Gender      string                 `json:"gender"`
	Affiliation []string `json:"totalPrice"`
	Associated  []string `json:"associated"`
	Masters     []string `json:"memo"`
	Apprentices []string `json:"apprentices"`
}

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

func RetreiveCharacter(id int)([]*StarWarsCharacterResult, error) {
	fmt.Println("ID AFTER: ", id)
	rows, err := db.ConnPool.Query(
		`SELECT
				id,
				name,
				home_world,
				born,
				died,
				species,
				gender,
				affiliation,
				associated,
				masters,
				apprentices
			FROM
				star_wars_characters
			WHERE (id = $1)`,
		id)

	if err != nil {
		fmt.Println("There was an error reading the quote builder table from the database:", err)
		return nil, fmt.Errorf("There was an error reading the quote builder table from the database %s: %s", id, err)
	}

	starWarsCharacterResultsArray := []*StarWarsCharacterResult{}

	defer rows.Close()

	for rows.Next() {

		var res StarWarsCharacterResult

		err = rows.Scan(
			&res.Id,
			&res.Name,
			&res.Homeworld,
			&res.Born,
			&res.Died,
			&res.Species,
			&res.Gender,
			&res.Affiliation,
			&res.Associated,
			&res.Masters,
			&res.Apprentices)

		if err != nil {
			fmt.Println("There was an error querying that database for the Star Wars Character:", err)
			continue
		}

		starWarsCharacterResultsArray = append(starWarsCharacterResultsArray, &res)
	}

	return starWarsCharacterResultsArray, nil
}
