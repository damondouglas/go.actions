package v2

import (
	"io"

	"github.com/damondouglas/go.actions/v2/dialogflow"
	"github.com/damondouglas/go.actions/v2/google"
)

var (
	// AudioType sets media type for audio.
	AudioType = &mediaTypeBase{"AUDIO"}

	// VideoType sets media type for video.
	VideoType = &mediaTypeBase{"VIDEO"}

	// EnglishUS sets language code to en-US.
	EnglishUS = &languageCodeBase{"en-US"}
)

// Event represents custom event response.
type Event struct {
	Name         string
	Parameters   map[string]string
	LanguageCode LanguageCode
}

// Encode custom event response.
func (e *Event) Encode(w io.Writer) error {
	resp := dialogflow.Event{
		FollowupEventInput: &dialogflow.FollowupEventInput{
			Name:         e.Name,
			Parameters:   e.Parameters,
			LanguageCode: e.LanguageCode.String(),
		},
	}
	return resp.Encode(w)
}

type baseResponse struct {
	Suggestions        []string
	underlyingResponse *dialogflow.Response
}

func (r *baseResponse) suggestions() {
	if r.Suggestions != nil && len(r.Suggestions) > 0 {
		suggestions := []*google.Suggestion{}
		for _, k := range r.Suggestions {
			suggestions = append(suggestions, &google.Suggestion{
				Title: k,
			})
		}
		r.underlyingResponse.Payload.Google.RichResponse.Suggestions = suggestions
	}
}

// Simple is a simple speech and text response to the user.
type Simple struct {
	baseResponse
	Say         string
	Display     string
	Suggestions []string
}

// Encode Simple response.
func (r *Simple) Encode(w io.Writer) error {
	r.underlyingResponse = base()
	r.baseResponse.Suggestions = r.Suggestions
	r.baseResponse.suggestions()
	r.underlyingResponse.Payload.Google.RichResponse.Items[0].SimpleResponse = &google.SimpleResponse{
		DisplayText:  r.Display,
		TextToSpeech: r.Say,
	}
	return r.underlyingResponse.Encode(w)
}

// Card is a card response to the client.
type Card struct {
	baseResponse
	RequiredResponse string
	FormattedText    string
	Title            string
	Subtitle         string
	Image            *Image
	Button           *Button
	Suggestions      []string
}

// Encode Basic Card response.
func (r *Card) Encode(w io.Writer) error {
	r.underlyingResponse = base()
	r.baseResponse.Suggestions = r.Suggestions
	r.baseResponse.suggestions()
	basicCard := &google.BasicCard{
		Title:        r.Title,
		Subtitle:     r.Subtitle,
		FormatedText: r.FormattedText,
	}
	if r.Image != nil {
		basicCard.Image = &google.Image{
			URL:               r.Image.URL,
			AccessibilityText: r.Image.AccessibilityText,
		}
	}
	if r.Button != nil {
		basicCard.Buttons = []*google.Button{
			&google.Button{
				Title: r.Button.Title,
				OpenURLAction: &google.OpenURLAction{
					URL: r.Button.URL,
				},
			},
		}
	}
	items := r.underlyingResponse.Payload.Google.RichResponse.Items
	items[0].SimpleResponse = &google.SimpleResponse{
		TextToSpeech: r.RequiredResponse,
	}
	items = append(items, &google.Item{
		BasicCard: basicCard,
	})
	r.underlyingResponse.Payload.Google.RichResponse.Items = items
	return r.underlyingResponse.Encode(w)
}

// CarouselBrowse configures a carousel to the client.
type CarouselBrowse struct {
	baseResponse
	RequiredResponse string
	Items            []*CarouselBrowseItem
	Suggestions      []string
}

// CarouselBrowseItem encodes the carousel browse item.
type CarouselBrowseItem struct {
	Title       string
	URL         string
	Description string
	Footer      string
	Image       *Image
}

// Encode a carousel browse to the client.
func (r *CarouselBrowse) Encode(w io.Writer) error {
	r.underlyingResponse = base()
	r.baseResponse.Suggestions = r.Suggestions
	r.baseResponse.suggestions()

	items := r.underlyingResponse.Payload.Google.RichResponse.Items
	items[0].SimpleResponse = &google.SimpleResponse{
		TextToSpeech: r.RequiredResponse,
	}

	carouselBrowse := &google.CarouselBrowse{
		Items: []*google.CarouselBrowseItem{},
	}

	carouselItems := carouselBrowse.Items

	for _, k := range r.Items {
		item := &google.CarouselBrowseItem{
			Title:       k.Title,
			Description: k.Description,
			Footer:      k.Footer,
		}

		if k.Image != nil {
			image := &google.Image{
				URL:               k.Image.URL,
				AccessibilityText: k.Image.AccessibilityText,
			}
			item.Image = image
		}

		if k.URL != "" {
			action := &google.OpenURLAction{
				URL: k.URL,
			}
			item.OpenURLAction = action
		}

		carouselItems = append(carouselItems, item)
	}

	carouselBrowse.Items = carouselItems

	items = append(items, &google.Item{
		CarouselBrowse: carouselBrowse,
	})
	r.underlyingResponse.Payload.Google.RichResponse.Items = items
	return r.underlyingResponse.Encode(w)
}

// Select presents options to the user.
type Select struct {
	baseResponse
	RequiredResponse string
	Items            []*SelectItem
	Suggestions      []string
}

// SelectItem is the item presented for the user to select.
type SelectItem struct {
	Key         string
	Synonyms    []string
	Description string
	Image       *Image
	Title       string
}

// Encode Select to the client.
func (r *Select) Encode(w io.Writer) error {
	r.underlyingResponse = base()
	r.baseResponse.Suggestions = r.Suggestions
	r.baseResponse.suggestions()

	items := r.underlyingResponse.Payload.Google.RichResponse.Items
	items[0].SimpleResponse = &google.SimpleResponse{
		TextToSpeech: r.RequiredResponse,
	}

	selectItems := []*google.SelectItem{}

	for _, k := range r.Items {
		item := &google.SelectItem{
			OptionInfo: &google.OptionInfo{
				Key:      k.Key,
				Synonyms: k.Synonyms,
			},
			Description: k.Description,
			Title:       k.Title,
		}
		if k.Image != nil {
			item.Image = &google.Image{
				URL:               k.Image.URL,
				AccessibilityText: k.Image.AccessibilityText,
			}
		}
		selectItems = append(selectItems, item)
	}

	carouselSelect := &google.Select{
		Items: selectItems,
	}

	intent := &google.SystemIntent{
		Intent: "actions.intent.OPTION",
		Data: &google.Data{
			Type:           "type.googleapis.com/google.actions.v2.OptionValueSpec",
			CarouselSelect: carouselSelect,
		},
	}

	r.underlyingResponse.Payload.Google.SystemIntent = intent

	return r.underlyingResponse.Encode(w)
}

// Confirmation presents a confirmation prompt to the user.
type Confirmation struct {
	baseResponse
	RequiredResponse string
	ConfirmationText string
	Suggestions      []string
}

// Encode Confirmation to the user.
func (r *Confirmation) Encode(w io.Writer) error {
	r.underlyingResponse = base()
	r.baseResponse.Suggestions = r.Suggestions
	r.baseResponse.suggestions()

	items := r.underlyingResponse.Payload.Google.RichResponse.Items
	items[0].SimpleResponse = &google.SimpleResponse{
		TextToSpeech: r.RequiredResponse,
	}

	intent := &google.SystemIntent{
		Intent: "actions.intent.CONFIRMATION",
		Data: &google.Data{
			Type: "type.googleapis.com/google.actions.v2.ConfirmationValueSpec",
			DialogSpec: &google.DialogSpec{
				RequestConfirmationText: r.ConfirmationText,
			},
		},
	}

	r.underlyingResponse.Payload.Google.SystemIntent = intent

	return r.underlyingResponse.Encode(w)
}

// Signin initiates oauth flow to user.
type Signin struct {
	baseResponse
	RequiredResponse string
}

// Encode Sigin.
func (r *Signin) Encode(w io.Writer) error {
	r.underlyingResponse = base()
	r.baseResponse.Suggestions = r.Suggestions
	r.baseResponse.suggestions()

	items := r.underlyingResponse.Payload.Google.RichResponse.Items
	items[0].SimpleResponse = &google.SimpleResponse{
		TextToSpeech: r.RequiredResponse,
	}

	intent := &google.SystemIntent{
		Intent: "actions.intent.SIGN_IN",
		Data: &google.Data{
			Type: "type.googleapis.com/google.actions.v2.SignInValueSpec",
		},
	}

	r.underlyingResponse.Payload.Google.SystemIntent = intent

	return r.underlyingResponse.Encode(w)
}

type languageCodeBase struct {
	value string
}

func (l *languageCodeBase) String() string {
	return l.value
}

// LanguageCode specifies language code.
type LanguageCode interface {
	String() string
}

// MediaType specifies media type.
type MediaType interface {
	String() string
}

type mediaTypeBase struct {
	value string
}

func (t *mediaTypeBase) String() string {
	return t.value
}

// Media presented to the user.
type Media struct {
	baseResponse
	RequiredResponse string
	Type             MediaType
	URL              string
	Description      string
	Icon             *Image
	Title            string
	Suggestions      []string
}

// Encode Media response.
func (r *Media) Encode(w io.Writer) error {
	r.underlyingResponse = base()
	r.baseResponse.Suggestions = r.Suggestions
	r.baseResponse.suggestions()

	items := r.underlyingResponse.Payload.Google.RichResponse.Items
	items[0].SimpleResponse = &google.SimpleResponse{
		TextToSpeech: r.RequiredResponse,
	}

	media := &google.MediaResponse{
		MediaType: r.Type.String(),
		MediaObjects: []*google.MediaObject{
			{
				ContentURL:  r.URL,
				Description: r.Description,
				Icon: &google.Image{
					URL:               r.Icon.URL,
					AccessibilityText: r.Icon.AccessibilityText,
				},
				Name: r.Title,
			},
		},
	}

	items = append(items, &google.Item{
		MediaResponse: media,
	})

	r.underlyingResponse.Payload.Google.RichResponse.Items = items

	return r.underlyingResponse.Encode(w)
}

// List builds a list response.
type List struct {
	baseResponse
	RequiredResponse string
	Items            []*SelectItem
	Suggestions      []string
}

// Encode a list.
func (r *List) Encode(w io.Writer) error {
	r.underlyingResponse = base()
	r.baseResponse.Suggestions = r.Suggestions
	r.baseResponse.suggestions()

	items := r.underlyingResponse.Payload.Google.RichResponse.Items
	items[0].SimpleResponse = &google.SimpleResponse{
		TextToSpeech: r.RequiredResponse,
	}

	selectItems := []*google.SelectItem{}

	for _, k := range r.Items {
		item := &google.SelectItem{
			OptionInfo: &google.OptionInfo{
				Key:      k.Key,
				Synonyms: k.Synonyms,
			},
			Description: k.Description,
			Title:       k.Title,
		}
		if k.Image != nil {
			item.Image = &google.Image{
				URL:               k.Image.URL,
				AccessibilityText: k.Image.AccessibilityText,
			}
		}
		selectItems = append(selectItems, item)
	}

	intent := &google.SystemIntent{
		Intent: "actions.intent.OPTION",
		Data: &google.Data{
			Type: "type.googleapis.com/google.actions.v2.OptionValueSpec",
			ListSelect: &google.Select{
				Items: selectItems,
			},
		},
	}

	r.underlyingResponse.Payload.Google.SystemIntent = intent

	return r.underlyingResponse.Encode(w)
}

// Button specifies button details and user click response.
type Button struct {
	Title string
	URL   string
}

// Image specifies image detail.
type Image struct {
	URL               string
	AccessibilityText string
}

// Encoder encodes the JSON response to io.Writer.
type Encoder interface {
	Encode(w io.Writer) error
}

func base() *dialogflow.Response {
	return &dialogflow.Response{
		Payload: &dialogflow.GooglePayload{
			Google: &google.Response{
				ExpectUserResponse: true,
				RichResponse: &google.RichResponse{
					Items: []*google.Item{
						&google.Item{},
					},
				},
			},
		},
	}
}
