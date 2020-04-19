package stockTracker

import (
    "fmt"
    //"log"
    //"time"
    "net/http"
    "io/ioutil"
    "strconv"
    "encoding/json"

    "PortfolioWebsite/src/go/common"
    "github.com/PuerkitoBio/goquery"
)

type ItemResult struct {
    Id           int    `json:id`
    Store        string `json:store`
    Name         string `json:name`
    Price        string `json:price`
    Availability string `json:availability`
}

var client http.Client

// We take the Best Buy HTML, as a string, strip everything we don't
// need away.
func StripBestBuyHtml(url string) ([]*ItemResult, error) {
    // Request the HTML page.
    req, err := http.NewRequest("GET", url, nil)
    req.Header.Add("Accept", `text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8`)
    req.Header.Add("Accept-Charset", `ISO-8859-1,utf-8;q=0.7,*;q=0.3`)
    req.Header.Add("Accept-Encoding", `none`)
    req.Header.Add("Accept-Language", `en-US,en;q=0.8`)
    req.Header.Add("Connection", `keep-alive`)
    req.Header.Add("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11`)
    res, err := client.Do(req)
    //res, err := http.Get(url)
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

    // We'll use this to store the data!
    bestBuyResultsArray := []*ItemResult{}

    //*[@id="pricing-price-28323958"]/div/div/div[1]/div/div[2]/div/div/div

    // Find the item list!
    doc.Find(".sku-item-list .sku-item").Each(func(i int, s *goquery.Selection) {
        // For each item found, get the name, price, and inventory
        name := s.Find(".sku-title").Text()
        price := s.Find("div.pricing-lib-price-8-2013-6.price-view-pb.priceView-layout-medium > div > div > div > div > div > span:nth-child(1)").Text()
        availability := s.Find(".fulfillment-add-to-cart-button").Text()
        //fmt.Printf("Review %d: %s - %s - %s\n", i, name, price, availability)
        // Now lets fill out struct!
        bestBuyResult := ItemResult{Id: i, Store: "Best Buy", Name: name, Price: price, Availability: availability}

        if len(name) > 0 && len(price) > 0 && len(availability) > 0 {
            bestBuyResultsArray = append(bestBuyResultsArray, &bestBuyResult)
        }
    })

    return bestBuyResultsArray, nil
}

func StripWalmartHtml(url string) ([]*ItemResult, error) {
    // Request the HTML page.
    req, err := http.NewRequest("GET", url, nil)
    req.Header.Add("Accept", `text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8`)
    req.Header.Add("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11`)
    res, err := client.Do(req)
    //res, err := http.Get(url)
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

    // We'll use this to store the data!
    walmartResultsArray := []*ItemResult{}

    // Find the item list!
    doc.Find(".search-result-gridview-item-wrapper").Each(func(i int, s *goquery.Selection) {
        // For each item found, get the name, price, and inventory
        name := s.Find("div > div > div > a > span").Text()
        price := s.Find(".product-price-with-fulfillment .price-group").Text()
        availability := s.Find("div > div > div > div > button > span").Text()
        //fmt.Printf("Review %d: %s - %s - %s\n", i, name, price, availability)
        // Now lets fill out struct!
        walmartResult := ItemResult{Id: i, Store: "Walmart", Name: name, Price: price, Availability: availability}

        if len(name) > 0 && len(price) > 0 && len(availability) > 0 {
            walmartResultsArray = append(walmartResultsArray, &walmartResult)
        }
    })

    return walmartResultsArray, nil
}

func StripTargetHtml(url string) ([]*ItemResult, error) {
    // https://redsky.target.com/v2/pdp/tcin/77464001?excludes=taxonomy,promotion,bulk_ship,rating_and_review_reviews,rating_and_review_statistics,question_answer_statistics

    // For some reason redsky.target.com is an api (returns json) that is publicly accessible. To use it, just substitute
    // the productID...

    // Lets get our product list and loop over it!
    targetProductsArray := TargetNintendoSwitchDeviceNumbers

    // We'll use this to store the data!
    targetResultsArray := []*ItemResult{}

    for index, element := range targetProductsArray {
        // Request the HTML page.
        // Target doesn't load everything at once. We need to wait until .product-grid is there...
        req, err := http.NewRequest("GET", "https://redsky.target.com/v2/pdp/tcin/"+element+"?excludes=taxonomy,promotion,bulk_ship,rating_and_review_reviews,rating_and_review_statistics,question_answer_statistics", nil)
        req.Header.Add("Accept", `text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8`)
        req.Header.Add("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11`)
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

        productMap := allInfoMap["product"].(map[string]interface{})
        priceMap := productMap["price"].(map[string]interface{})
        listPriceMap := priceMap["listPrice"].(map[string]interface{})
        itemMap := productMap["item"].(map[string]interface{})
        nameMap := itemMap["product_description"].(map[string]interface{})
        availabilityMap := productMap["available_to_promise_network"].(map[string]interface{})

        //test := d.(map[string]interface{})["data"].(map[string]interface{})["type"]
        name := nameMap["title"].(string)
        price := strconv.FormatFloat(listPriceMap["price"].(float64), 'f', -1, 64)
        availability := availabilityMap["availability"].(string)
        targetResult := ItemResult{Id: index, Store:"Target", Name: name, Price: price, Availability: availability}

        if len(name) > 0 && len(price) > 0 && len(availability) > 0 {
            targetResultsArray = append(targetResultsArray, &targetResult)
        }

    }

    return targetResultsArray, nil
}

func StripGameStopHtml(url string) ([]*ItemResult, error) {
    // Request the HTML page.
    // Target doesn't load everything at once. We need to wait until .product-grid is there...
    req, err := http.NewRequest("GET", url, nil)
    req.Header.Add("Accept", `text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8`)
    req.Header.Add("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11`)
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

    // We'll use this to store the data!
    gameStopResultsArray := []*ItemResult{}

    // Find the item list!
    doc.Find(".product-grid-tile-wrapper").Each(func(i int, s *goquery.Selection) {
        // GameStop, thankfully, has a nice Json for each item! Lets grab that!
        allInfoJson := s.Find("div.product-tile-header > div.pdp-link > a").AttrOr("data-gtmdata", "{}")
        allInfoMap, err := common.GetMapFromData(allInfoJson)
        if err != nil {
            fmt.Println("Error converting data into Json: ", err)
        }
        priceMap := allInfoMap["price"].(map[string]interface{})
        productInfoMap := allInfoMap["productInfo"].(map[string]interface{})

        // For each item found, get the name, price, and inventory
        name := productInfoMap["name"].(string)
        // They don't have the $ symbol, so we'll add it.
        price := "$"+priceMap["sellingPrice"].(string)
        availability := productInfoMap["availability"].(string)
        //fmt.Printf("Review %d: %s - %s - %s\n", i, name, price, availability)
        // Now lets fill out struct!
        gameStopResult := ItemResult{Id: i, Store:"GameStop", Name: name, Price: price, Availability: availability}

        if len(name) > 0 && len(price) > 0 && len(availability) > 0 {
            gameStopResultsArray = append(gameStopResultsArray, &gameStopResult)
        }
    })

    return gameStopResultsArray, nil
}

