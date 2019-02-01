package goResources

import "fmt"

func AddCharacter(name string, born string, associated map[string]interface{}, gender string, affiliation map[string]interface{}, masters map[string]interface{})  (string, error) {

	fmt.Println("Name: ", name)
	fmt.Println("Born: ", born)
	fmt.Println("Associated: ", associated)
	fmt.Println("Gender: ", gender)
	fmt.Println("Affiliation: ", affiliation)
	fmt.Println("Masters: ", masters)


	return "", nil
}
