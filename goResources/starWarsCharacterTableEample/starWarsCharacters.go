package starWarsCharacterTableEample

import (
	"PortfolioWebsite/goResources/db"
	"encoding/json"
	"fmt"
	"reflect"
)

type CharacterResult struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Homeworld   string   `json:"homeworld"`
	Born        string   `json:"born"`
	Died        string   `json:"died"`
	Species     string   `json:"species"`
	Gender      string   `json:"gender"`
	Affiliation string `json:"affiliation"`
	Associated  string `json:"associated"`
	Masters     string `json:"masters"`
	Apprentices string `json:"apprentices"`
}

type StarWarsCharacterResult struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Homeworld   string   `json:"homeworld"`
	Born        string   `json:"born"`
	Died        string   `json:"died"`
	Species     string   `json:"species"`
	Gender      string   `json:"gender"`
	Affiliation string `json:"affiliation"`
	Associated  string `json:"associated"`
	Masters     string `json:"masters"`
	Apprentices string`json:"apprentices"`
}

func AddCharacter(name string, homeworld string, born string, died string, gender string, species string, affiliation string , associated string, masters string, apprentices string)  (int, error) {
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
		`INSERT INTO 
				star_wars_characters(
					name, 
					home_world, 
					born, 
					died, 
					species, 
					gender, 
					affiliation, 
					associated, 
					masters, 
					apprentices) 
			VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) 
			RETURNING id`,
		name, homeworld, born, died, species, gender, affiliation, associated, masters, apprentices).Scan(&lastInsertId)
	if err != nil {
		fmt.Println("Error saving character to database: ", err)
		return 0, err
	}

	return lastInsertId, nil
}

func LoadAllStarWarsCharacters(Name, Species string) ([]*CharacterResult, error) {
	fmt.Println("NAME: ", Name)
	fmt.Println("Species: ", Species)
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
			WHERE (name = $1 OR $1 = '') 
			AND   (species = $2 OR $2 = '')`,
		Name, Species)

	if err != nil {
		fmt.Println("There was an error reading the star_wars_characters table from the database:", err)
	}

	characterResultsArray := []*CharacterResult{}

	defer rows.Close()

	for rows.Next() {

		var characterResult CharacterResult

		err = rows.Scan(
			&characterResult.Id,
			&characterResult.Name,
			&characterResult.Homeworld,
			&characterResult.Born,
			&characterResult.Died,
			&characterResult.Species,
			&characterResult.Gender,
			&characterResult.Affiliation,
			&characterResult.Associated,
			&characterResult.Masters,
			&characterResult.Apprentices)

		if err != nil {
			fmt.Println("There was an error querying that database for all Star Wars Characters:", err)
			continue
		}

		characterResultsArray = append(characterResultsArray, &characterResult)
	}


	return characterResultsArray, nil
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

func ResubmitCharacter(character map[string]interface{}) (string, error) {
	//fmt.Println("Quote: ", quote)

	// We have to cast our ID to an int. This worked.
	// https://tanaikech.github.io/2017/06/02/changing-from-float64-to-int-for-values-did-unmarshal-using-mapstringinterface/
	id := int(character["id"].(float64))

	editedName := character["name"]
	editedHomeworld := character["homeworld"]
	editedBorn := character["born"]
	editedDied := character["died"]
	editedSpecies := character["species"]
	editedGender := character["gender"]
	editedAffiliation, _ := json.Marshal(character["affiliation"])
	editedAssociated, _ := json.Marshal(character["associated"])
	editedMasters, _ := json.Marshal(character["masters"])
	editedApprentices, _ := json.Marshal(character["apprentices"])


	fmt.Println("Affiliation Type: ", reflect.TypeOf(editedAffiliation))
	//totalPrice := quote["quote"].(map[string]interface{})["Final Price"]
	//memeo := quote["quote"].(map[string]interface{})["Memo"]

	_, err := db.ConnPool.Exec(
		`UPDATE
				star_wars_characters
			SET
				id = $1,
				name = $2,
				home_world = $3,
				born = $4,
				died = $5,
				species = $6,
				gender = $7,
				affiliation = $8,
				associated = $9,
				masters = $10,
				apprentices = $11
			WHERE 
				(id = $1)`,
		id,
		editedName,
		editedHomeworld,
		editedBorn,
		editedDied,
		editedSpecies,
		editedGender,
		editedAffiliation,
		editedAssociated,
		editedMasters,
		editedApprentices)
	if err != nil {
		return "Failed to save to the DB", err
	}
	return "Success!", nil
}
