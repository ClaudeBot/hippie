package main

import (
	"encoding/json"
	"flag"
	"github.com/claudebot/hipchat-go/hipchat"
	"github.com/claudebot/hippie/lambda"
	_ "github.com/claudebot/hippie/scripts/excuses"
	_ "github.com/claudebot/hippie/scripts/giphy"
	_ "github.com/claudebot/hippie/scripts/httpcodes"
	_ "github.com/claudebot/hippie/scripts/ping"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

var (
	client    *hipchat.Client
	roomID    int
	webhookID int
	botToken  = flag.String("token", "", "the HipChat API access token")
	botRoom   = flag.String("room", "", "the name of the room for the bot to be deployed in")
	botURL    = flag.String("url", "", "the publicly accessible URL to the bot (try `ngrok` for testing)")
	botKey    = flag.String("key", "hippie", "the unique identifier for the bot's webhook")
	botPort   = flag.Int("port", 8080, "the port to that the bot will use to accept webhooks")
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		log.Println("invalid request (empty body)")
		http.Error(w, "invalid request (empty body)", 400)
		return
	}
	defer r.Body.Close()

	var roomMsg RoomMessage
	if err := json.NewDecoder(r.Body).Decode(&roomMsg); err != nil {
		log.Printf("unable to decode room message: %+v", err)
		http.Error(w, "unable to decode room message", 500)
	}

	result, err := lambda.Run(roomMsg.Item.Message.Message)
	if err != nil || len(result) == 0 {
		log.Printf("error or no matching lambda: %+v", err)
		return
	}

	w.Header().Set("Content-Type", "application/javascript")

	roomNot := hipchat.NotificationRequest{
		Message:       result,
		MessageFormat: "text",
	}

	if err := json.NewEncoder(w).Encode(&roomNot); err != nil {
		log.Printf("unable to encode response message: %+v", err)
		http.Error(w, "unable to encode response message", 500)
	}
}

func init() {
	flag.Parse()
	if len(*botToken) == 0 || len(*botRoom) == 0 || len(*botURL) == 0 {
		log.Fatalln("usage: hippie -token=<hipchat token> -room=<room name> -url=<url to hosted bot>")
	}

	log.Printf("launching hippie with: '%s', '%s', '%s'\n", *botToken, *botRoom, *botURL)

	// TODO: this could probably be 'handled' a lot better ...
	client = hipchat.NewClient(*botToken)

	base, err := url.Parse(*botURL)
	if err != nil {
		log.Fatalln(err)
	}

	main, err := url.Parse("webhook")
	if err != nil {
		log.Fatalln(err)
	}

	u := base.ResolveReference(main)

	if err := webhookRegister(*botRoom, *botKey, u); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/webhook", webhookHandler)
	log.Printf("running hippie on port %d ...\n", *botPort)
	err := http.ListenAndServe(":"+strconv.Itoa(*botPort), nil)
	if err != nil {
		panic(err)
	}
}
