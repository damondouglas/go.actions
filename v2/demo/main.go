package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/damondouglas/go.actions/v2/identity"
	"golang.org/x/oauth2"

	"github.com/damondouglas/go.actions/v2"
	"github.com/damondouglas/go.actions/v2/dialogflow"
	"google.golang.org/appengine"
)

type requestHandler func(*dialogflow.Request, http.ResponseWriter, *http.Request)

const (
	pathToSecretKey = "SECRET_PATH"
	projectIDKey    = "PROJECT_ID"
	testHostKey     = "TEST_HOST"
)

var (
	suggestions = []string{"simple", "basic", "select", "media", "list"}
	button      = &v2.Button{
		URL:   "https://assistant.google.com/",
		Title: "This is a button",
	}
	image = &v2.Image{
		URL:               "https://picsum.photos/200",
		AccessibilityText: "This is accessibility text.",
	}
	keepConvoGoing = "  What else would you like to see?"

	carouselItems = []*v2.CarouselBrowseItem{
		{
			Title:       "Item 1 Title",
			URL:         "https://google.com",
			Description: "Item 1 Description",
			Footer:      "Item 1 footer",
			Image:       image,
		},
		{
			Title:       "Item 2 Title",
			URL:         "https://google.com",
			Description: "Item 2 Description",
			Footer:      "Item 2 footer",
			Image:       image,
		},
		{
			Title:       "Item 3 Title",
			URL:         "https://google.com",
			Description: "Item 3 Description",
			Footer:      "Item 3 footer",
			Image:       image,
		},
	}

	selectItems = []*v2.SelectItem{
		{
			Key:         "SELECTION_KEY_ONE",
			Synonyms:    []string{"synonym of title 1", "synonym of title 2", "synonym of title 3"},
			Description: "Item 1 Description",
			Image:       image,
			Title:       "Item 1 Title",
		},
		{
			Key:         "SELECTION_KEY_TWO",
			Synonyms:    []string{"Google Home Assistant", "Assistant on the Google Home"},
			Description: "Google Home is a voice-activated speaker powered by Google Assistant.",
			Image:       image,
			Title:       "Google Home",
		},
		{
			Key:         "SELECTION_KEY_THREE",
			Synonyms:    []string{"Google Pixel XL", "Pixel", "Pixel XL"},
			Description: "Pixel. Phone by Google.",
			Image:       image,
			Title:       "Google Pixel",
		},
	}

	responseMap = map[string]v2.Encoder{
		"simple": &v2.Simple{
			Display:     "This is an example of a simple response." + keepConvoGoing,
			Say:         "This is an example of a simple response." + keepConvoGoing,
			Suggestions: suggestions,
		},
		"basic": &v2.Card{
			RequiredResponse: "Simple Responses must be included" + keepConvoGoing,
			Title:            "Title this is a title",
			Subtitle:         "This is a subtitle",
			FormattedText:    "This is a basic card. Text in a basic card can include \"quotes\",\n__bold__, ***bold italic*** or ___strong emphasis___.",
			Button:           button,
			Image:            image,
			Suggestions:      suggestions,
		},
		"carousel": &v2.CarouselBrowse{
			RequiredResponse: "Simple responses must be included.",
			Items:            carouselItems,
			Suggestions:      suggestions,
		},
		"select": &v2.Select{
			RequiredResponse: "Simple responses must be included.",
			Items:            selectItems,
			Suggestions:      suggestions,
		},
		"confirmation": &v2.Confirmation{
			RequiredResponse: "Simple responses must be included.",
			ConfirmationText: "Can you confirm?",
			Suggestions:      suggestions,
		},
		"media": &v2.Media{
			RequiredResponse: "Simple responses must be included.",
			Type:             v2.AudioType,
			URL:              "http://storage.googleapis.com/automotive-media/Jazz_In_Paris.mp3",
			Description:      "A funky Jazz tune",
			Icon: &v2.Image{
				URL:               "http://storage.googleapis.com/automotive-media/album_art.jpg",
				AccessibilityText: "Information about jazz",
			},
			Title:       "Jazz in Paris",
			Suggestions: suggestions,
		},
		"list": &v2.List{
			RequiredResponse: "Simple responses must be included.",
			Items:            selectItems,
			Suggestions:      suggestions,
		},
	}

	intentMap = map[string]requestHandler{
		"signin":         signin,
		"profile":        profile,
		"fulfill_signin": profile,
	}
)

func main() {
	h := &identity.Handler{}
	h.Store = store
	http.HandleFunc("/auth", identity.AuthHandler)
	http.HandleFunc("/exch", h.TokenHandler)
	http.HandleFunc("/action", action)
	appengine.Main()
}

func store(ctx context.Context, token *oauth2.Token) {
	log.Println("TOKEN", token)
}

func action(w http.ResponseWriter, r *http.Request) {
	req, err := request(r)
	if err != nil {
		log.Fatalf("%v", err)
	}

	key := extractType(req)

	if encoder, ok := responseMap[key]; ok {
		err = encoder.Encode(w)
		if err != nil {
			log.Fatalf("%v", err)
		}
	} else {
		dispatch(req, w, r)
	}

}

func request(r *http.Request) (req *dialogflow.Request, err error) {

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func extractType(req *dialogflow.Request) string {
	if req == nil {
		return ""
	}
	inputs := req.OriginalDetectIntentRequest.Payload.Inputs
	if len(inputs) == 0 {
		return ""
	}
	args := inputs[0].Arguments
	if args == nil {
		return ""
	}
	if len(args) == 0 {
		return ""
	}
	return args[0].TextValue
}

func dispatch(req *dialogflow.Request, w http.ResponseWriter, r *http.Request) {
	intent := req.QueryResult.Intent.DisplayName
	log.Println("INTENT: ", intent)

	if handler, ok := intentMap[intent]; ok {
		handler(req, w, r)
	}

	evt := &v2.Event{
		Name: "goback",
	}
	err := evt.Encode(w)
	if err != nil {
		log.Fatal(err)
	}
}

func signin(req *dialogflow.Request, w http.ResponseWriter, r *http.Request) {
	evt := v2.Signin{
		RequiredResponse: "Welcome to the gallery.",
	}
	err := evt.Encode(w)
	if err != nil {
		log.Fatal(err)
	}

}

func profile(req *dialogflow.Request, w http.ResponseWriter, r *http.Request) {
	basic := v2.Simple{
		Display: req.OriginalDetectIntentRequest.Payload.User.AccessToken,
		Say:     "Hi",
	}
	err := basic.Encode(w)
	if err != nil {
		log.Fatal(err)
	}
}
