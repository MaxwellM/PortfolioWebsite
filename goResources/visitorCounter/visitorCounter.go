package visitorCounter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type knownIP struct {
	IP        string    `json:"ip"`
	TimeStamp time.Time `json:"timeStamp"`
}

var knownIPs []knownIP


func AppendToIPStruct(ip string)(*[]knownIP, error){

	knownIPs = append(knownIPs, knownIP{IP: ip, TimeStamp: time.Now()})

	updateIPReturn, err := UpdateVisitors(knownIPs)
	if err != nil {
		fmt.Println("Couldn't return UpdateVisitors: ", err)
	} else {
		fmt.Println("Update Visitors Return: ", updateIPReturn)
	}

	return &knownIPs, nil
}

//func UpdateVisitors(ipStruct []knownIP) (map[string]interface{}, error) {
//	updateVisitorsNewJSON, err := json.Marshal(ipStruct)
//	if err != nil {
//		fmt.Println("Error writing weather to file", err)
//		return nil, err
//	}
//
//	file, err := os.OpenFile("goResources/visitorCounter/visitors.json", os.O_APPEND|os.O_WRONLY, 0644)
//	if err != nil {
//		fmt.Println("Error opening the Visitors JSON file", err)
//		return nil, err
//	}
//	byteVisitorsFile, err := ioutil.ReadAll(file)
//	if err != nil {
//		fmt.Println("Error reading the Visitors JSON file", err)
//		return nil, err
//	}
//
//	var knownIPs []knownIP
//
//	err := json.Unmarshal(byteVisitorsFile, &knownIPs)
//
//	byteVisitorsFile = append(byteVisitorsFile, updateVisitorsNewJSON)
//
//}

func UpdateVisitors(ipStruct []knownIP) (map[string]interface{}, error) {

	updateVisitorsReturnJSON, err := json.Marshal(ipStruct)
	if err != nil {
		fmt.Println("Error writing weather to file", err)
		return nil, err
	}

	file, err := os.OpenFile("goResources/visitorCounter/visitors.json", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening the Visitors JSON file", err)
		return nil, err
	}

	defer file.Close()

	//visitorJSONByte := []byte(updateVisitorsReturnJSON)
	//visitorJSONReturn, err := file.Write(visitorJSONByte)
	//if err != nil {
	//	fmt.Println("Error writing Visitor JSON to Visitor JSON file")
	//	return nil, err
	//}

	visitorJSONReturn := ioutil.WriteFile("goResources/visitorCounter/visitors.json", updateVisitorsReturnJSON, 0644)
	fmt.Println(visitorJSONReturn)

	allVisitors := map[string]interface{}{
		"Visitors": visitorJSONReturn,
	}

	return allVisitors, nil
}

func ReadVisitors() ([]map[string]interface{}, error) {
	b, err := ioutil.ReadFile("goResources/visitorCounter/visitors.json") // just pass the file name
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	var dat []map[string]interface{}
	err = json.Unmarshal(b, &dat)
	if err != nil {
		fmt.Println("Error unmarshaling Visitor JSON", err)
	}

	return dat, nil
}
