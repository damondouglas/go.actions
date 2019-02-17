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

// Encode JSON from Response.
func (r *Response) Encode(w io.Writer) (err error) {
	return json.NewEncoder(w).Encode(r)
}
