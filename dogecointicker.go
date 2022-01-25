package main
import (
	"log"
	"github.com/Zauberstuhl/go-coinbase"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"time"
	"os"
	"io/ioutil"
	"strconv"
	"fmt"
)


func main() {
	consumerKey := "<replace>"
	consumerSecret := "<replace>"
	accessToken := "<replace>"
	accessSecret := "<replace>"

	// string to Tweet out	
	var botstring string

	// Basic string text
	botstring = "The current value of 1 DOGE in USD is: $"

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	c := coinbase.APIClient{
	  Key: "<replace>",
	  Secret: "<replace>",
	}

	// Let's get the values of DOGE!
	exchanges, err := c.GetExchangeRates("DOGE")
	if err != nil {
		log.Fatal(err)
	}

	// get the current time, and more importantly current hour
	t := time.Now()
	h := t.Hour()

	if h == 0 {
		f, err := os.Create("<replace>") //text file to store the "base" value for the day
		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		_, err2 := f.WriteString(string(exchanges.Data.Rates["USD"]))

		if err2 != nil {
			log.Fatal(err2)
		}
	}
	
	// Let's add the USD value of DOGE to our tweet string
	botstring += string(exchanges.Data.Rates["USD"])
	if h != 0 {
		content, err := ioutil.ReadFile("<replace>") //base value for the day to compare to
		if err != nil {
			log.Fatal(err)
		}

		basevalue, erragain := strconv.ParseFloat(string(content), 64)
		if erragain != nil {
			log.Fatal(erragain)
		}
		currvalue, errnoway := strconv.ParseFloat(string(exchanges.Data.Rates["USD"]), 64)
		if errnoway != nil {
			log.Fatal(errnoway)
		}

		if currvalue > basevalue {
			upby := currvalue - basevalue
			botstring += " (📈 up "
			botstring += fmt.Sprintf("%f", upby)
		}

		if currvalue < basevalue {
			downby := basevalue - currvalue
			botstring += " (📉 down "
			botstring += fmt.Sprintf("%f", downby)
		}

		botstring += " so far today)."
	}
	botstring += " #dogecoin"

	
	tweet, resp, err := client.Statuses.Update(botstring, nil)
	if err != nil {
	    log.Println(err)
	}
	log.Printf("%+v\n", resp)
	log.Printf("%+v\n", tweet)
}
