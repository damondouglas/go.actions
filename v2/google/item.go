package google

// SimpleResponse containing speech or text to show the user.
type SimpleResponse struct {
	TextToSpeech string `json:"textToSpeech,omitempty"`
	SSML         string `json:"ssml,omitempty"`
	DisplayText  string `json:"displayText,omitempty"`
}

// Item of the response.
type Item struct {
	SimpleResponse   *SimpleResponse   `json:"simpleResponse,omitempty"`
	BasicCard        *BasicCard        `json:"basicCard,omitempty"`
	StructedResponse *StructedResponse `json:"structuredResponse,omitempty"`
	MediaResponse    *MediaResponse    `json:"mediaResponse,omitempty"`
	CarouselBrowse   *CarouselBrowse   `json:"carouselBrowse,omitempty"`
	TableCard        *TableCard        `json:"tableCard,omitempty"`
}

// BasicCard for displaying some information, e.g. an image and/or text.
type BasicCard struct {
	Title               string    `json:"title,omitempty"`
	Subtitle            string    `json:"subtitle,omitempty"`
	FormatedText        string    `json:"formattedText,omitempty"`
	Image               *Image    `json:"image,omitempty"`
	Buttons             []*Button `json:"buttons,omitempty"`
	ImageDisplayOptions string    `json:"imageDisplayOptions,omitempty"`
}

// StructedResponse defined for app to respond with structured data.
type StructedResponse struct {
	// Todo: implement this later
}

// MediaResponse indicating a set of media to be played within the conversation.
type MediaResponse struct {
	MediaType    string         `json:"mediaType,omitempty"`
	MediaObjects []*MediaObject `json:"mediaObjects,omitempty"`
}

// CarouselBrowse presents a set of AMP documents as a carousel of large-tile items.
// Items may be selected to launch their associated AMP document in an AMP viewer.
type CarouselBrowse struct {
	Items               []*CarouselBrowseItem `json:"items,omitempty"`
	ImageDisplayOptions string                `json:"imageDisplayOptions,omitempty"`
}

// CarouselBrowseItem is the item presented in a CarouselBrowse response.
type CarouselBrowseItem struct {
	Title         string         `json:"title,omitempty"`
	Description   string         `json:"description,omitempty"`
	Footer        string         `json:"footer,omitempty"`
	Image         *Image         `json:"image,omitempty"`
	OpenURLAction *OpenURLAction `json:"openUrlAction,omitempty"`
}

// TableCard for displaying a table of text.
type TableCard struct {
	Title            string            `json:"title,omitempty"`
	Subtitle         string            `json:"subtitle,omitempty"`
	Image            *Image            `json:"image,omitempty"`
	ColumnProperties []*ColumnProperty `json:"columnProperties,omitempty"`
	Rows             []*Row            `json:"rows,omitempty"`
	Buttons          []*Button         `json:"buttons,omitempty"`
}

// ColumnProperty specifies column in table card.
type ColumnProperty struct {
	Header              string `json:"header,omitempty"`
	HorizontalAlignment string `json:"horizontalAlignment,omitempty"` // Todo: make enum type
}

// Row defines the row in a table card.
type Row struct {
	Cells       []*Cell `json:"cells,omitempty"`
	DivideAfter bool    `json:"divideAfter,omitempty"`
}

// Cell determines text content in cell of table card.
type Cell struct {
	Text string `json:"text,omitempty"`
}

// Image displays an image.displayed in the card.
type Image struct {
	URL               string `json:"url,omitempty"`
	AccessibilityText string `json:"accessibilityText,omitempty"`
	Height            int    `json:"height,omitempty"`
	Width             int    `json:"width,omitempty"`
}

// MediaObject represents one media object which is returned with MediaResponse. Contains information about the media, such as name, description, url, etc.
type MediaObject struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ContentURL  string `json:"contentUrl,omitempty"`
	LargeImage  *Image `json:"largeImage,omitempty"`
	Icon        *Image `json:"icon,omitempty"`
}

// Button that usually appears at the bottom of a card.
type Button struct {
	Title         string         `json:"title,omitempty"`
	OpenURLAction *OpenURLAction `json:"openUrlAction,omitempty"`
}
