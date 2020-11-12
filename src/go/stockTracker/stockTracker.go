package stockTracker

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "regexp"
    "strconv"

    "PortfolioWebsite/src/go/common"
    "github.com/PuerkitoBio/goquery"
)

type ItemResult struct {
    Id           int    `json:id`
    Store        string `json:store`
    Name         string `json:name`
    Price        string `json:price`
    Availability string `json:availability`
    URL          string `json:url`
}

var client http.Client

func GetStockInfoFromApiSource(vendor string, item string) ([]*ItemResult, error) {
    // We'll use this to store the data!
    itemResultsArray := []*ItemResult{}
    switch vendor {
        case "BestBuy":
            // This is our url
            bestBuyUrlBase := BestBuyDevices["URL"][0]

            // Get our BestBuy API key
            bestBuyApiInfo, err := common.ReadJsonFile("bestBuyApiKey.json")
            if err != nil {
                fmt.Println("Error reading JSON file!")
                return nil, err
            }
            bestBuyApiKey := bestBuyApiInfo["key"].(string)

            for key, value := range BestBuyDevices {
                fmt.Println(key)
                if item == key {
                    for index, v := range value {
                        // Now we need to add our device to our URL
                        var re = regexp.MustCompile(`%`)
                        bestBuyUrl := re.ReplaceAllString(bestBuyUrlBase, v)
                        // Request the HTML page.
                        req, err := http.NewRequest("GET", bestBuyUrl+bestBuyApiKey, nil)
                        req = common.SetHeaders(req)
                        res, err := client.Do(req)
                        if err != nil {
                            fmt.Println("Error getting data from URL")
                            return nil, err
                        }
                        defer res.Body.Close()
                        if res.StatusCode != 200 {
                            fmt.Println("status code error: %d %s", res.StatusCode, res.Status)
                            return nil, err
                        }

                        bytes, err := ioutil.ReadAll(res.Body)
                        if err != nil {
                            fmt.Println("Couldn't convert RESP to []Byte for your Current Conditions Request: ", err)
                        }

                        var allInfoMap map[string]interface{}
                        err = json.Unmarshal(bytes, &allInfoMap)
                        if err != nil {
                            return nil, err
                        }
                        // For each item found, get the name, price, and inventory
                        name := allInfoMap["name"].(string)
                        price := strconv.FormatFloat(allInfoMap["salePrice"].(float64), 'f', -1, 64)
                        availability := strconv.FormatBool(allInfoMap["onlineAvailability"].(bool))
                        url := allInfoMap["url"].(string)
                        // Now lets fill out struct!
                        bestBuyResult := ItemResult{
                            Id: index,
                            Store: "Best Buy",
                            Name: name,
                            Price: "$"+price,
                            Availability: availability,
                            URL: url,
                        }

                        if len(name) > 0 && len(price) > 0 && len(availability) > 0 {
                            itemResultsArray = append(itemResultsArray, &bestBuyResult)
                        }
                    }
                }
            }
        case "Target":
            // This is our url
            targetUrlBase := TargetDevices["URL"][0]

            // Get our Target API key
            targetApiInfo, err := common.ReadJsonFile("targetApiKey.json")
            if err != nil {
                fmt.Println("Error reading JSON file!")
                return nil, err
            }
            targetApiKey := targetApiInfo["key"].(string)

            for key, value := range TargetDevices {
                fmt.Println(key)
                if item == key {
                    for index, v := range value {
                        // Now we need to add our device to our URL
                        var re = regexp.MustCompile(`%`)
                        targetUrl := re.ReplaceAllString(targetUrlBase, v)
                        fmt.Println("URL: ", targetUrl+targetApiKey)
                        // Request the HTML page.
                        req, err := http.NewRequest("GET", targetUrl+targetApiKey, nil)
                        req = common.SetHeaders(req)
                        res, err := client.Do(req)
                        if err != nil {
                            fmt.Println("Error getting data from URL")
                            return nil, err
                        }
                        defer res.Body.Close()
                        if res.StatusCode != 200 {
                            fmt.Println("status code error: %d %s", res.StatusCode, res.Status)
                            return nil, err
                        }

                        bytes, err := ioutil.ReadAll(res.Body)
                        if err != nil {
                            fmt.Println("Couldn't convert RESP to []Byte for your Current Conditions Request: ", err)
                        }

                        var allInfoMap map[string]interface{}
                        err = json.Unmarshal(bytes, &allInfoMap)
                        if err != nil {
                            return nil, err
                        }
                        // For each item found, get the name, price, and inventory
                        productMap := allInfoMap["product"].(map[string]interface{})
                        //priceMap := productMap["price"].(map[string]interface{})
                        //listPriceMap := priceMap["listPrice"].(map[string]interface{})
                        itemMap := productMap["item"].(map[string]interface{})
                        nameMap := itemMap["product_description"].(map[string]interface{})
                        availabilityMap := productMap["available_to_promise_network"].(map[string]interface{})

                        //test := d.(map[string]interface{})["data"].(map[string]interface{})["type"]
                        name := nameMap["title"].(string)
                        //price := strconv.FormatFloat(listPriceMap["price"].(float64), 'f', -1, 64)
                        price, err := getTargetPrice(v)
                        if err != nil {
                            fmt.Println("Error obtaining Target price! ", err)
                        }
                        availability := availabilityMap["availability"].(string)
                        url := itemMap["buy_url"].(string)
                        targetResult := ItemResult{
                            Id: index,
                            Store:"Target",
                            Name: name,
                            Price: "$"+price,
                            Availability: availability,
                            URL: url,
                        }

                        if len(name) > 0 && len(price) > 0 && len(availability) > 0 {
                            itemResultsArray = append(itemResultsArray, &targetResult)
                        }
                    }
                }
            }
        case "Walmart":
            for key, value := range WalmartDevices {
                fmt.Println(key)
                if item == key {
                    fmt.Println("URL: ", value)
                    // Request the HTML page.
                    req, err := http.NewRequest("GET", value, nil)
                    req = common.SetHeaders(req)
                    res, err := client.Do(req)
                    if err != nil {
                        fmt.Println("Error getting data from URL")
                        return nil, err
                    }
                    defer res.Body.Close()
                    if res.StatusCode != 200 {
                        fmt.Println("status code error: %d %s", res.StatusCode, res.Status)
                        return nil, err
                    }

                    // Load the HTML document
                    doc, err := goquery.NewDocumentFromReader(res.Body)
                    if err != nil {
                        fmt.Println("Error loading the HTML document: ", err)
                        return nil, err
                    }

                    index := 0
                    // Find the item list!
                    doc.Find(".search-result-gridview-item-wrapper").Each(func(i int, s *goquery.Selection) {
                        // For each item found, get the name, price, and inventory
                        name := s.Find("div > div > div > a > span").Text()
                        price := s.Find(".product-price-with-fulfillment .price-group").Text()
                        availability := s.Find("div > div > div > div > button > span").Text()
                        url := s.Find("div.search-result-gridview-item.clearfix.arrange-fill > div:nth-child(5) > div > a").AttrOr("href", "")
                        fmt.Printf("Review %d: %s - %s - %s\n", i, name, price, availability, url)
                        // Now lets fill out struct!
                        walmartResult := ItemResult{
                            Id: index,
                            Store: "Walmart",
                            Name: name,
                            Price: price,
                            Availability: availability,
                            URL: "https://www.walmart.com/"+url,
                        }
                        index++
                        if len(name) > 0 && len(price) > 0 && len(availability) > 0 {
                            itemResultsArray = append(itemResultsArray, &walmartResult)
                        }
                    })
                }
            }
            case "GameStop":
             for key, value := range GameStopDevices {
                 fmt.Println(key)
                 if item == key {
                   // Request the HTML page.
                    // GameStop doesn't load everything at once. We need to wait until .product-grid is there...
                    fmt.Println("URL: ", value)
                    req, err := http.NewRequest("GET", value, nil)
                    req = common.SetHeaders(req)
                    res, err := client.Do(req)
                    if err != nil {
                        fmt.Println("Error getting data from URL")
                        return nil, err
                    }
                    defer res.Body.Close()
                    if res.StatusCode != 200 {
                        fmt.Println("status code error: %d %s", res.StatusCode, res.Status)
                        return nil, err
                    }
                    // Load the HTML document
                    doc, err := goquery.NewDocumentFromReader(res.Body)
                    if err != nil {
                        fmt.Println("Error loading the HTML document: ", err)
                        return nil, err
                    }

                    fmt.Println("Doc :", doc)

                    // We'll use this to store the data!
                    gameStopResultsArray := []*ItemResult{}

                    // Find the item list!
                    doc.Find(".product-grid-tile-wrapper").Each(func(i int, s *goquery.Selection) {
                        // GameStop, thankfully, has a nice Json for each item! Lets grab that!
                        if allInfoJson, ok := s.Find(".pdp-link").Attr("data-gtmdata"); ok {
                                allInfoMap, err := common.GetMapFromData(allInfoJson)
                                if err != nil {
                                    fmt.Println("Error converting data into Json: ", err)
                                }
                                priceMap := allInfoMap["price"].(map[string]interface{})
                                productInfoMap := allInfoMap["productInfo"].(map[string]interface{})

                                // For each item found, get the name, price, and inventory
                                name := productInfoMap["name"].(string)
                                price := priceMap["sellingPrice"].(string)
                                availability := productInfoMap["availability"].(string)
                                url := s.Find(".pdp-link > a").AttrOr("href", "")
                                //fmt.Printf("Review %d: %s - %s - %s\n", i, name, price, availability)
                                // Now lets fill out struct!
                                gameStopResult := ItemResult{
                                    Id: i,
                                    Store:"GameStop",
                                    Name: name,
                                    Price: "$"+price,
                                    Availability: availability,
                                    URL: "https://www.gamestop.com/"+url,
                                }

                                if len(name) > 0 && len(price) > 0 && len(availability) > 0 {
                                    gameStopResultsArray = append(gameStopResultsArray, &gameStopResult)
                                }
                        }
                    })

    return gameStopResultsArray, nil
                 }
             }
        default:
            fmt.Println("Well, didn't recognize the vendor. Better add it!")
            return nil, fmt.Errorf("Well, didn't recognize the vendor. Better add it!")
    }
    return itemResultsArray, nil
}

func getTargetPrice(item string)(string, error) {
         // This is our url
        targetUrlBase := TargetDevices["PriceURL"][0]

        // Get our Target API key
        targetApiInfo, err := common.ReadJsonFile("targetApiKey.json")
        if err != nil {
            fmt.Println("Error reading JSON file!")
            return "", err
        }
        targetApiKey := targetApiInfo["key"].(string)

        // Now we need to add our device to our URL
        var re = regexp.MustCompile(`%`)
        targetUrl := re.ReplaceAllString(targetUrlBase, item)
        fmt.Println("URL: ", targetUrl+targetApiKey)
        // Request the HTML page.
        req, err := http.NewRequest("GET", targetUrl+targetApiKey, nil)
        req = common.SetHeaders(req)
        res, err := client.Do(req)
        if err != nil {
            fmt.Println("Error getting data from URL")
            return "", err
        }
        defer res.Body.Close()
        if res.StatusCode != 200 {
            fmt.Println("status code error: %d %s", res.StatusCode, res.Status)
            return "", err
        }

        bytes, err := ioutil.ReadAll(res.Body)
        if err != nil {
            fmt.Println("Couldn't convert RESP to []Byte for your Current Conditions Request: ", err)
        }

        var allInfoMap map[string]interface{}
        err = json.Unmarshal(bytes, &allInfoMap)
        if err != nil {
            return "", err
        }
        // For each item found, get the name, price, and inventory
        priceMap := allInfoMap["price"].(map[string]interface{})
        price := strconv.FormatFloat(priceMap["current_retail"].(float64), 'f', -1, 64)
        return price, nil
}

