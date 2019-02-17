package google

import "encoding/json"

// Input represents the input data payload.
type Input struct {
	RawInputs []*RawInput
	Intent    string
	Arguments []*Argument
}

// RawInput transcription from each turn of conversation.
type RawInput struct {
	InputType string //Todo make enum type
	Query     string
	URL       string
}

// Argument list of provided argument values for the input requested by the Action.
type Argument struct {
	Name          string
	RawText       string
	TextValue     string
	Status        json.RawMessage
	IntValue      string
	FloatValue    float64
	BoolValue     bool
	DatetimeValue struct {
		Date struct {
			Year  int
			Month int
			Day   int
		}
		Time struct {
			Hours   int
			Minutes int
			Seconds int
			Nanos   int
		}
	}
	PlaceValue      *Location
	Extension       json.RawMessage
	StructuredValue json.RawMessage
}

// ExpectedInput the Action expects.
type ExpectedInput struct {
	InputPrompt        *InputPrompt      `json:"inputPrompt,omitempty"`
	PossibleIntents    []*ExpectedIntent `json:"possibleIntents,omitempty"`
	SpeechBiasingHints []string          `json:"speechBiasingHints,omitempty"`
}

// InputPrompt used for assistant to guide user to provide an input for the app's question.
type InputPrompt struct {
	RichInitialPrompt *RichResponse     `json:"richInitialPrompt,omitempty"`
	NoInputPrompts    []*SimpleResponse `json:"noInputPrompts,omitempty"`
}
