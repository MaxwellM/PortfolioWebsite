package goExamples

import (
	"encoding/json"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"io/ioutil"
)

type TwitterInfo struct {
	APIKey            string
	APISecretKey      string
	AccessToken       string
	AccessSecretToken string
}

func getTwitterInfo() (TwitterInfo, error) {
	file, err := ioutil.ReadFile("twitterKey.json")
	if err != nil {
		fmt.Println("Error reading JSON file: ", err)
	}

	//fmt.Println("File: ", file)

	data := TwitterInfo{}

	err = json.Unmarshal([]byte(file), &data)
	//fmt.Println("Error unmarshinlg JSON: ", err)

	//fmt.Println("DATA: ", data)

	return data, nil
}

func GetMostRecentTweet() (){

}

func SubmitTweet(tweetToSend string) (*twitter.Tweet, error) {
	fmt.Println("Submitting Tweet! ", tweetToSend)

	twitterInfo, err := getTwitterInfo()
	if err != nil {
		fmt.Println("Error obtaining our Twitter Info! ", err.Error())
		return nil, err
	}

	//fmt.Println("Twitter Info: ", twitterInfo)

	//config := oauth1.NewConfig("consumerKey", "consumerSecret")
	//token := oauth1.NewToken("accessToken", "accessSecret")
	config := oauth1.NewConfig(twitterInfo.APIKey, twitterInfo.APISecretKey)
	token := oauth1.NewToken(twitterInfo.AccessToken, twitterInfo.AccessSecretToken)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Home Timeline
	//tweets, resp, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
	//	Count: 2,
	//})

	//fmt.Println("Home: ", tweets, resp, err)

	// Send a Tweet
	tweet, resp, err := client.Statuses.Update(tweetToSend, nil)

	fmt.Println("Send Tweet: ", tweet)
	fmt.Println("RESP: ", resp)

	// Status Show
	//tweet, resp, err = client.Statuses.Show(585613041028431872, nil)

	//fmt.Println("Status: ", tweet, resp, err)

	// Search Tweets
	//search, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{
	//	Query: "I linked my Twitter feed directly to my portfolio website! Check it out",
	//})

	//fmt.Println("Search: ", search, resp, err)

	// User Show
	//user, resp, err := client.Users.Show(&twitter.UserShowParams{
	//	ScreenName: "MAXintosh2010",
	//})

	//fmt.Println("User Show: ", user, resp, err)

	// Followers
	//followers, resp, err := client.Followers.List(&twitter.FollowerListParams{})

	//fmt.Println("Followers: ", followers, resp, err)

	return tweet, nil
}