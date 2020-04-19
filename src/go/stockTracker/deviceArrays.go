package stockTracker

// Houses our BestBuy SKU numbers for the Nintendo Switch
var BestBuyNintendoSwitchSKUNumbers [6]string = [6]string{
    "6257139", "6364255", "6364253", "6257142", "6257135", "6257148"}

// Houses our Target Device Numbers for the Nintendo Switch
var TargetNintendoSwitchDeviceNumbers [4]string = [4]string{
    "77464001", "77419248", "77464002", "77419246"}


var BestBuyDevices map[string][]string = map[string][]string {
    "URL": []string{"https://api.bestbuy.com/v1/products/%.json?show=sku,name,salePrice,onlineAvailability,url&apiKey="},
    "Nintendo Switch": []string{"6257139", "6364255", "6364253", "6257142", "6257135", "6257148"},
    "Dyson V11 Vacuum": []string{"6401091", "6401113", "6331930", "6331929"},
    "Apple AirPods": []string{"6084400", "5706659", "6083595"},
}

var TargetDevices map[string][]string = map[string][]string {
    "URL": []string{"https://redsky.target.com/v2/pdp/tcin/%?excludes=taxonomy,promotion,bulk_ship,rating_and_review_reviews,rating_and_review_statistics,question_answer_statistics"},
    "Nintendo Switch": []string{"77464001", "77419248", "77464002", "77419246"},
    "Dyson V11 Vacuum": []string{"54529249", "54529271"},
    "Apple AirPods": []string{"54191101", "54191097", "54191099"},
}

var WalmartDevices map[string]string = map[string]string {
    "Nintendo Switch": "https://www.walmart.com/search/?cat_id=2636_4646529_2002476&facet=retailer%3AWalmart.com",
    "Dyson V11 Vacuum": "https://www.walmart.com/search/?cat_id=0&facet=retailer%3AWalmart.com&query=dyson+v11",
    "Apple AirPods": "https://www.walmart.com/browse/apple-airpods/0?_be_shelf_id=15781&cat_id=0&facet=shelf_id%3A15781%7C%7Cretailer%3AWalmart.com&max_price=&min_price=100&page=1&value=%24100+and+up",
}