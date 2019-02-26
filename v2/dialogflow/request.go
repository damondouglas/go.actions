package dialogflow

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/damondouglas/go.actions/v2/google"
)

// Request is the Fullfillment HTTP Request from Dialogflow
type Request struct {
	ResponseID  string
	QueryResult struct {
		QueryText                string
		Action                   string
		RawParameterData         map[string]json.RawMessage `json:"parameters"`
		Parameters               map[string]*Parameter
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

// DecodeParameters decodes parameter data.
func (r *Request) DecodeParameters() {
	r.QueryResult.Parameters = map[string]*Parameter{}
	for key, data := range r.QueryResult.RawParameterData {
		r.QueryResult.Parameters[key] = newParameter(data)
	}
}

// Decode io.Reader into Request.
func Decode(r io.Reader) (req *Request, err error) {
	err = json.NewDecoder(r).Decode(&req)
	req.DecodeParameters()
	return req, err
}

// Parameter argument in request.
type Parameter struct {
	strValue string
	mapValue map[string]string
}

func newParameter(data []byte) *Parameter {
	p := &Parameter{}
	var value map[string]string
	err := json.Unmarshal(data, &value)
	if err == nil && len(value) > 0 {
		p.mapValue = value
		p.strValue = ""
	} else {
		p.mapValue = map[string]string{}
		p.strValue = strings.Replace(string(data), "\"", "", -1)
	}

	return p
}

// String returns string value.
func (p *Parameter) String() string {
	return p.strValue
}

// MapValue returns parameter value that maps to key-value pair.
func (p *Parameter) MapValue() map[string]string {
	return p.mapValue
}
