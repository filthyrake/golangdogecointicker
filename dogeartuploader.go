package main

import (
	"fmt"
	"encoding/base64"
	"strconv"
	"io/ioutil"
	"net/url"
	"log"

	"github.com/ChimeraCoder/anaconda"
)


func main() {

	anaconda.SetConsumerKey("<replace>")
	anaconda.SetConsumerSecret("<replace>")
	api := anaconda.NewTwitterApi("<replace>", "<replace>")

	data, err := ioutil.ReadFile("<replace with file path>")
	if err != nil {
		log.Fatal(err)
	}

	mediaResponse, err := api.UploadMedia(base64.StdEncoding.EncodeToString(data))
	if err != nil {
		log.Fatal(err)
	}

	v := url.Values{}
	v.Set("media_ids", strconv.FormatInt(mediaResponse.MediaID, 10))

	tweetString := fmt.Sprintf("Today's #dogecoin art, randomly generated fresh for you!")

	_, err = api.PostTweet(tweetString, v)
	if err != nil {
		log.Fatal(err)
	} else {
		// fmt.Println(result)
	}
}
