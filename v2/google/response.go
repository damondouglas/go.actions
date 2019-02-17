package google

// Response is the response sent by the fulfillment to Google Assistant.
type Response struct {
	ConversationToken  string           `json:"conversationToken,omitempty"`
	UserStorage        string           `json:"userStorage,omitempty"`
	ResetUserStorage   bool             `json:"resetUserStorage,omitempty"`
	ExpectUserResponse bool             `json:"expectUserResponse,omitempty"`
	ExpectedInputs     []*ExpectedInput `json:"expectedInputs,omitempty"`
	RichResponse       *RichResponse    `json:"richResponse,omitempty"`
	CustomPushMessage  *PushMessage     `json:"customPushMessage,omitempty"`
	IsInSandbox        bool             `json:"isInSandbox,omitempty"`
	SystemIntent       *SystemIntent    `json:"systemIntent,omitempty"`
}

// RichResponse that can include audio, text, cards, suggestions and structured data.
type RichResponse struct {
	Items             []*Item            `json:"items,omitempty"`
	Suggestions       []*Suggestion      `json:"suggestions,omitempty"`
	LinkOutSuggestion *LinkOutSuggestion `json:"linkOutSuggestion,omitempty"`
}

// Suggestion represents suggestion payload.
type Suggestion struct {
	Title string `json:"title,omitempty"`
}

// LinkOutSuggestion provides url action from payload.
type LinkOutSuggestion struct {
	DestinationName string         `json:"destinationName,omitempty"`
	OpenURLAction   *OpenURLAction `json:"openUrlAction,omitempty"`
}

// OpenURLAction of the App or Site to open when the user taps the suggestion chip.
type OpenURLAction struct {
	URL         string      `json:"url,omitempty"`
	AndroidApp  *AndroidApp `json:"androidApp,omitempty"`
	URLTypeHint string      `json:"urlTypeHint,omitempty"`
}

// AndroidApp specifies android app launch action.
type AndroidApp struct {
	PackageName string    `json:"packageName,omitempty"`
	Versions    *Versions `json:"versions,omitempty"`
}

// Versions specifies version of android app to launch.
type Versions struct {
	MinVersion int `json:"minVersion,omitempty"`
	MaxVersion int `json:"maxVersion,omitempty"`
}
