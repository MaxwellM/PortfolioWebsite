package stockTracker

import (
    "fmt"
    //"log"
    "net/http"

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
    targetResultsArray := []*ItemResult{}

    // Find the item list!
    doc.Find(".productGridContainer").Each(func(i int, s *goquery.Selection) {
        // For each item found, get the name, price, and inventory
        // li > div > div > div > div > div > div > div > a
        name := s.Find(".h-display-flex > a").Text()
        // li > div > div > div > div > div > div > span
        price := s.Find(".h-text-bs").Text()
        availability := s.Find("li > div > div > div > div > div > div > div > div > div > div > div > div > button").Text()
        //fmt.Printf("Review %d: %s - %s - %s\n", i, name, price, availability)
        // Now lets fill out struct!
        targetResult := ItemResult{Id: i, Store:"Target", Name: name, Price: price, Availability: availability}

        if len(name) > 0 && len(price) > 0 && len(availability) > 0 {
            targetResultsArray = append(targetResultsArray, &targetResult)
        }
    })

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
        fmt.Println("Json: ", allInfoJson)
        allInfoMap, err := common.GetMapFromData(allInfoJson)
        if err != nil {
            fmt.Println("Error converting data into Json: ", err)
        }
        fmt.Println("After Json: ", allInfoMap)
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

