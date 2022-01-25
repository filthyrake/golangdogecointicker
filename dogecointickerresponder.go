package main
import (
	"fmt"
	"log"
	"strconv"
	"math/rand"
	"time"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)


func main() {
	consumerKey := "<replace>"
	consumerSecret := "<replace>"
	accessToken := "<replace>"
	accessSecret := "<replace>"

	// Array of strings!  This will host our fun facts
	var dogeFacts [6]string


	//generate random numbers weee
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := len(dogeFacts) - 1

	// Array of Sender IDs so I know who to reply to
	var senders []int

	// Let's populate those facts
	dogeFacts[0] = "Dogecoin was first unleashed upon the world on December 6th, 2013."
	dogeFacts[1] = "Dogecoin was created by Billy Markus and Jackson Palmer."
	dogeFacts[2] = "The almighty Doge has been featured multiple times in NASCAR races on cars."
	dogeFacts[3] = "The official font of Dogecoin is the greatest font in the world: Comic Sans."
	dogeFacts[4] = "Dogecoin has no cap.  Except in pictures.  Those are adorable."
	dogeFacts[5] = "The face of Dogecoin is the super cute Shiba Inu dog."


	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)


	// List most recent 100 Direct Messages.  100 is arbitrary, just a bigish number that seemed safe to cover all messages.
	messages, _, err := client.DirectMessages.EventsList(
		&twitter.DirectMessageEventsListParams{Count: 100},
	)
	//fmt.Println("User's DIRECT MESSAGES:")
	if err != nil {
		log.Fatal(err)
	}
	// Iterate through the messages, reply, and destroy
	for _, event := range messages.Events {
		//fmt.Printf("%+v\n", event)
		//fmt.Printf("  %+v\n", event.Message)
		//fmt.Printf("  %+v\n", event.Message.Data)
		i, err := strconv.Atoi(event.Message.SenderID)
		if err != nil {
			log.Fatal(err)
		}
		senders = append(senders, i)
		//fmt.Printf("  %+v\n", senders)
		_, err = client.DirectMessages.EventsDestroy(event.ID)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("DM Events Delete:\n err: %v\n", err)
	}

	// Now let's send fun facts!
	for _, sender := range senders {
		//fmt.Printf("Senders are %v\n", sender)
		// Pick a dogeFact
		var randNum = rand.Intn(max - min + 1) + min
		var dogeFact = dogeFacts[randNum]
		fmt.Printf(dogeFact)
		// Create Direct Message event
			var event, _, err = client.DirectMessages.EventsNew(&twitter.DirectMessageEventsNewParams{
				Event: &twitter.DirectMessageEvent{
					Type: "message_create",
					Message: &twitter.DirectMessageEventMessage{
						Target: &twitter.DirectMessageTarget{
							RecipientID: strconv.Itoa(sender),
						},
						Data: &twitter.DirectMessageData{
							Text: dogeFact,
						},
					},
				},
			})
			fmt.Printf("DM Event New:\n%+v, %v\n", event, err)
			if err != nil {
				log.Fatal(err)
			}
	}

	


}
