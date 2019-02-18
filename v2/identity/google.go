package identity

import (
	"net/http"

	oauth2 "google.golang.org/api/oauth2/v2"

	"github.com/damondouglas/go.actions/v2/dialogflow"
)

// TokenInfo acquires Oauth2 token info from dialogflow.Request User IDToken.
func TokenInfo(client *http.Client, req *dialogflow.Request) (*oauth2.Tokeninfo, error) {
	idToken := req.OriginalDetectIntentRequest.Payload.User.IDToken
	oauth2Service, err := oauth2.New(client)
	if err != nil {
		return nil, err
	}
	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfoCall.IdToken(idToken)
	return tokenInfoCall.Do()
}
