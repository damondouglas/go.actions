package dialogflow

import (
	"encoding/json"
	"io"

	"github.com/damondouglas/go.actions/v2/google"
)

// Response from a dialogflow webhook.
type Response struct {
	Payload *GooglePayload `json:"payload"`
}

// GooglePayload represents response payload for Google Assistant.
type GooglePayload struct {
	Google *google.Response `json:"google"`
}

// Event represents a custom event payload from webhook.
type Event struct {
	FollowupEventInput *FollowupEventInput `json:"followupEventInput"`
}

// Encode JSON response.
func (e *Event) Encode(w io.Writer) (err error) {
	return json.NewEncoder(w).Encode(e)
}

// FollowupEventInput respresents custom event payload data.
type FollowupEventInput struct {
	Name         string            `json:"name"`
	Parameters   map[string]string `json:"parameters"`
	LanguageCode string            `json:"languageCode"`
}

// Encode JSON from Response.
func (r *Response) Encode(w io.Writer) (err error) {
	return json.NewEncoder(w).Encode(r)
}
