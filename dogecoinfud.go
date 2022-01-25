package main
import (
	"log"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"math/rand"
	"time"
)


func main() {
	consumerKey := "<replace>"
	consumerSecret := "<replace>"
	accessToken := "<replace>"
	accessSecret := "<replace>"

	// string to Tweet out	
	var botstring string
	var dogeFUD [5]string

	//generate random numbers weee
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := len(dogeFUD) - 1

	dogeFUD[0] = "Random FUD about #dogecoin: Dogecoin is abandoned or otherwise unmaintained.  FALSE!  It is actively maintained and you can check it out for yourself: https://github.com/dogecoin/dogecoin"
	dogeFUD[1] = "Random FUD about #dogecoin: The top 20 wallets hold more than 50% of all dogecoin.  FALSE! Most of those wallets (and the biggest) are exchanges where customer coins are pooled."
	dogeFUD[2] = "Random FUD about #dogecoin: Dogecoin is useless and has no value.  FALSE!  You can actually spend it at many online shops, without having to convert it to something else first!"
	dogeFUD[3] = "Random FUD about #dogecoin: Since dogecoin has no cap it has infinite supply and wont ever be worth anything.  FALSE!  It has limited yearly issuance (much like the dollar) and has already grown massively in value."
	dogeFUD[4] = "Random FUD about #dogecoin: Shibetoshi Nakamoto (or Elon Musk or <insert name here>) is the CEO of Dogecoin.  FALSE!  Dogecoin is not a company of any sort, and has no CEO."

	var randNum = rand.Intn(max - min + 1) + min
	botstring = dogeFUD[randNum]


	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	
	tweet, resp, err := client.Statuses.Update(botstring, nil)
	if err != nil {
	    log.Println(err)
	}
	log.Printf("%+v\n", resp)
	log.Printf("%+v\n", tweet)
}
