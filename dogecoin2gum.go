package main
import (
	"github.com/ChimeraCoder/anaconda"
	"log"
	"github.com/Zauberstuhl/go-coinbase"
	"strconv"
	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/common"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"image/color"
	"math/rand"
	"time"
	"os"
	"fmt"
	"io/ioutil"
	"encoding/base64"
	"net/url"
)

// color map function for Domain Wrap art
func cmap(r, m1, m2 float64) color.RGBA {
	rgb := color.RGBA{
		uint8(common.Constrain(m1*200*r, 0, 255)),
		uint8(common.Constrain(r*200, 0, 255)),
		uint8(common.Constrain(m2*255*r, 70, 255)),
		255,
	}
	return rgb
}


func main() {
	anaconda.SetConsumerKey("<replace>")
	anaconda.SetConsumerSecret("<replace>")
	api := anaconda.NewTwitterApi("<replace>", "<replace>")

	// current price of 1 pack of gum
	var gumprice float64
	gumprice = 1

	// string to Tweet out	
	var botstring string

	// Basic string text
	botstring = "If we assume a pack of gum costs $1, then 1 DOGE at the current price will buy you: "


	c := coinbase.APIClient{
	  Key: "<replace>",
	  Secret: "<replace>",
	}

	// Let's get the values of DOGE!
	exchanges, err := c.GetExchangeRates("DOGE")
	if err != nil {
		log.Fatal(err)
	}
	
	f := string(exchanges.Data.Rates["USD"])

	s, err2 := strconv.ParseFloat(f, 64)
	if err2 != nil {
		log.Fatal(err2)
	}

	var howmuch float64
	howmuch = s / gumprice

	var howmanypieces float64
	howmanypieces = howmuch * 15
	var piecesint int
	piecesint = int(howmanypieces)

	var widthimg int
	widthimg = 100 * piecesint

	// Create the canvas
	var ca = generativeart.NewCanva(widthimg, 300)

	var howmuchstring string

	howmuchstring = strconv.FormatFloat(howmuch, 'f', -1, 64)

	var howmanypiecesstring string

	howmanypiecesstring = strconv.FormatFloat(howmanypieces, 'f', 1, 64)

	// Make the initial blank canvas
	rand.Seed(time.Now().Unix())
 
	ca.SetBackground(common.White)

	ca.FillBackground()

	// Write it to a file
	ca.ToPNG("<replace>")

	// Let's get the canvas file back open and ready for gum
	image1,err := os.Open("<replace>")
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}

	first, err := png.Decode(image1)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer image1.Close()
 
	// Now for the pack of gum
	image2,err := os.Open("<replace>")
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}
	second, err := png.Decode(image2)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer image2.Close()


	offset := image.Pt(5, 5)
	b := first.Bounds()
	image3 := image.NewRGBA(b)
	draw.Draw(image3, b, first, image.ZP, draw.Src)

	// Write the gum onto the canvas
	for i := 0; i < piecesint; i++ {
		var addloffset int
		addloffset = 5 + (i*94)
		offset = image.Pt(addloffset, 5)
		draw.Draw(image3, second.Bounds().Add(offset), second, image.ZP, draw.Over)
	}

	// Save the final image
	third,err := os.Create("<replace>")
	if err != nil {
		log.Fatalf("failed to create: %s", err)
	}
	jpeg.Encode(third, image3, &jpeg.Options{jpeg.DefaultQuality})
	defer third.Close()

	// Read the image to get ready to tweet it
	data, err := ioutil.ReadFile("<replace>")
	if err != nil {
		log.Fatal(err)
	}

	mediaResponse, err := api.UploadMedia(base64.StdEncoding.EncodeToString(data))
	if err != nil {
		log.Fatal(err)
	}


	// Let's add the number of packs of gum to our tweet string
	botstring += howmuchstring
	botstring += " packs of gum.  This also works out to roughly "
	botstring += howmanypiecesstring
	botstring += " sticks of gum. #dogecoin"

	v := url.Values{}
	v.Set("media_ids", strconv.FormatInt(mediaResponse.MediaID, 10))

	tweetString := fmt.Sprintf(botstring)

	_, err = api.PostTweet(tweetString, v)
	if err != nil {
		log.Fatal(err)
	} else {
		// fmt.Println(result)
	}
}
