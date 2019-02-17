package dialogflow

import (
	"encoding/json"
	"io"

	"github.com/damondouglas/go.actions/v2/google"
)

// Request is the Fullfillment HTTP Request from Dialogflow
type Request struct {
	ResponseID  string
	QueryResult struct {
		QueryText                string
		Action                   string
		AllRequiredParamsPresent bool
		OutputContexts           []struct {
			Name string
		}
		Intent struct {
			Name        string
			DisplayName string
		}
		IntentDetectionConfidence float64
		DiagnosticInfo            struct {
		}
		LanguageCode string
	}
	OriginalDetectIntentRequest struct {
		Source  string
		Version string
		Payload *google.Request
	}
	Session string
}

// Decode io.Reader into Request.
func Decode(r io.Reader) (req *Request, err error) {
	err = json.NewDecoder(r).Decode(&req)
	return req, err
}
