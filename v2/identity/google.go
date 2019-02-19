package identity

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/damondouglas/go.actions/v2/dialogflow"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	oauth2api "google.golang.org/api/oauth2/v2"
)

const (
	redirectBase    = "https://oauth-redirect.googleusercontent.com/r/"
	clientIDKey     = "client_id"
	redirectURIKey  = "redirect_uri"
	stateKey        = "state"
	responseTypeKey = "response_type"
)

var (
	queryKeys = []string{
		clientIDKey,
		redirectURIKey,
		stateKey,
		responseTypeKey,
	}

	// BaseScopes defines basic authorization scope.
	BaseScopes = []string{"email"}
)

// AuthorizationParameters of the authorization request.
type AuthorizationParameters struct {
	ClientID     string `json:"client_id"`
	RedirectURI  string `json:"redirect_uri"`
	State        string `json:"state"`
	ResponseType string `json:"response_type"`
}

// RedirectParameters configures redirect url.
type RedirectParameters struct {
	State     string
	Offline   bool
	Force     bool
	ProjectID string
	Scopes    []string
}

// TokenInfo acquires Oauth2 token info from dialogflow.Request User IDToken.
func TokenInfo(client *http.Client, req *dialogflow.Request) (*oauth2api.Tokeninfo, error) {
	idToken := req.OriginalDetectIntentRequest.Payload.User.IDToken
	oauth2apiService, err := oauth2api.New(client)
	if err != nil {
		return nil, err
	}
	tokenInfoCall := oauth2apiService.Tokeninfo()
	tokenInfoCall.IdToken(idToken)
	return tokenInfoCall.Do()
}

// ParseAuthorizationRequestURL parses authorization requests from OAuth2 implicit flow.
func ParseAuthorizationRequestURL(rawQuery string) (param *AuthorizationParameters, err error) {
	qryParams, err := url.ParseQuery(rawQuery)
	if err != nil {
		return nil, err
	}
	qryMap := map[string]string{}
	for _, key := range queryKeys {
		qryMap[key] = qryParams.Get(key)
	}

	data, err := json.Marshal(qryMap)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &param)
	if err != nil {
		return nil, err
	}

	param.RedirectURI, err = url.PathUnescape(param.RedirectURI)
	if err != nil {
		return nil, err
	}

	return param, nil
}

// IsAuthorizationRequestValid verifies authorization request originated from Google Actions platform.
func IsAuthorizationRequestValid(param *AuthorizationParameters, projectID string, clientID string) bool {
	return param.RedirectURI == redirectBase+projectID && param.ClientID == clientID
}

// RedirectURL builds redirectURL after user authorization.
func RedirectURL(config *oauth2.Config, param *RedirectParameters) string {
	config.RedirectURL = redirectBase + param.ProjectID
	config.Scopes = param.Scopes
	offlineOpt := oauth2.AccessTypeOnline
	if param.Offline {
		offlineOpt = oauth2.AccessTypeOffline
	}

	result := config.AuthCodeURL(param.State, offlineOpt)

	if param.Force {
		result = config.AuthCodeURL(param.State, oauth2.ApprovalForce, offlineOpt)
	}

	return result
}

// AuthorizationHandler of initial OAuth2 flow.
func AuthorizationHandler(w http.ResponseWriter, r *http.Request, config *oauth2.Config, param *RedirectParameters) (http.HandlerFunc, error) {
	authParam, err := ParseAuthorizationRequestURL(r.URL.RawQuery)
	if err != nil {
		return nil, err
	}

	if !IsAuthorizationRequestValid(authParam, param.ProjectID, config.ClientID) {
		return nil, fmt.Errorf("authorization is invalid want: %s, got: %s and want: %s, got: %s", authParam.RedirectURI, r.URL.RawQuery, config.ClientID, authParam.ClientID)
	}

	param.State = authParam.State

	redirectURL := RedirectURL(config, param)

	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, redirectURL, http.StatusFound)
	}, nil
}

type clientSecretConfig struct {
	Web struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
	}
}

// ConfigFromPath loads oauth2.Config from path.
func ConfigFromPath(pathToSecret string) (*oauth2.Config, error) {
	var rawConfig *clientSecretConfig
	data, err := ioutil.ReadFile(pathToSecret)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &rawConfig)
	if err != nil {
		return nil, err
	}
	return &oauth2.Config{
		ClientID:     rawConfig.Web.ClientID,
		ClientSecret: rawConfig.Web.ClientSecret,
		Endpoint:     google.Endpoint,
	}, nil
}
