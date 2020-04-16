package stockTracker

import (
    "fmt"
    "regexp"
    //"strings"
    //"io"
    //"os"
)

type Items struct {
    AllItems []Item
}

type Item struct {
    Name string
	Price string
	Availability string
}

// We take the Best Buy HTML, as a string, strip everything we don't
// need away.
func StripBestBuyHtml(data string) (string, error) {
    // We're trying to grab everything inside of the sku-item-list list
    re := regexp.MustCompile(`(?s)sku-item-list.*\</ol>`)
    match := re.FindString(data)

    // Now we need to break the match into submatches, one for each line item
    re = regexp.MustCompile(`(?s)li class="sku-item"(.*?)\</li>`)
	submatchall := re.FindAllString(match, -1)

	// Now we'll loop through the items and get the price and availability
	fmt.Println("Length: ", len(submatchall))
	for index, element := range submatchall {
	    fmt.Println("Index: ", index)
	    // Grabs the name
        re = regexp.MustCompile(`(?s)sku-title(.*?)\</h4>`)
        nameResult := re.FindString(string(element))
        fmt.Println("Name: ", nameResult)

	    // Grabs the price
	    re = regexp.MustCompile(`(?s)Your price for this item(.*?)\</span>`)
	    priceBlockResult := re.FindString(string(element))
	    re = regexp.MustCompile(`(?s)>(.*?)\</span>`)
	    price := re.FindString(priceBlockResult)
	    re = regexp.MustCompile(`(?s)>(.*?)\<`)
	    price = re.FindString(price)
	    fmt.Println("Price: ", price)

        // Grabs the stock
	    re = regexp.MustCompile(`(?s)fulfillment-add-to-cart-button(.*)\</button></div></div>`)
	    stockResult := re.FindString(string(element))
	    re = regexp.MustCompile(`>(Sold.*?|Add.*?)<`)
	    stockFinal := re.FindString(stockResult)
	    fmt.Println("Stock: ", stockFinal)
	}
    return "", nil
}