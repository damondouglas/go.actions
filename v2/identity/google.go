package identity

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/damondouglas/go.actions/v2/dialogflow"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	oauth2api "google.golang.org/api/oauth2/v2"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

const (
	redirectHost     = "oauth-redirect.googleusercontent.com"
	redirectBasePath = "/r"
	clientIDKey      = "client_id"
	redirectURIKey   = "redirect_uri"
	stateKey         = "state"
	responseTypeKey  = "response_type"
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
	State       string
	Offline     bool
	Force       bool
	ProjectID   string
	Scopes      []string
	RedirectURL string
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

type identityBase struct {
	authParam     *AuthorizationParameters
	redirectParam *RedirectParameters
	config        *oauth2.Config
}

// New Identity.
func New(config *oauth2.Config, redirectParam *RedirectParameters) Identity {
	i := &identityBase{
		config:        config,
		redirectParam: redirectParam,
	}

	return i
}

func (i *identityBase) authRedirectURL() string {
	redirectURL := &url.URL{
		Scheme: "https",
		Host:   redirectHost,
		Path:   redirectBasePath + "/" + i.redirectParam.ProjectID,
	}
	return redirectURL.String()
}

func (i *identityBase) parseAuthorizationRequestURL(rawQuery string) (err error) {
	qryParams, err := url.ParseQuery(rawQuery)
	if err != nil {
		return err
	}
	qryMap := map[string]string{}
	for _, key := range queryKeys {
		qryMap[key] = qryParams.Get(key)
	}

	data, err := json.Marshal(qryMap)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &i.authParam)
	if err != nil {
		return err
	}

	i.authParam.RedirectURI, err = url.PathUnescape(i.authParam.RedirectURI)
	if err != nil {
		return err
	}

	return nil
}

func (i *identityBase) validateAuthorizationRequest() error {
	redirectURL, err := url.ParseRequestURI(i.authParam.RedirectURI)
	if err != nil {
		return err
	}
	redirectURLNotValid := errors.New("redirect URI is not valid")
	clientIDNotValid := errors.New("clientID not valid")
	if redirectURL.Host != redirectHost {
		return redirectURLNotValid
	}

	if redirectURL.Path != redirectBasePath+"/"+i.redirectParam.ProjectID {
		return redirectURLNotValid
	}

	if i.authParam.ClientID != i.config.ClientID {
		return clientIDNotValid
	}

	return nil
}

func (i *identityBase) authURL() string {
	i.config.RedirectURL = i.redirectParam.RedirectURL
	i.config.Scopes = i.redirectParam.Scopes
	offlineOpt := oauth2.AccessTypeOnline
	if i.redirectParam.Offline {
		offlineOpt = oauth2.AccessTypeOffline
	}

	result := i.config.AuthCodeURL(i.redirectParam.State, offlineOpt)

	if i.redirectParam.Force {
		result = i.config.AuthCodeURL(i.redirectParam.State, oauth2.ApprovalForce, offlineOpt)
	}

	return result
}

// AuthorizationHandler of initial OAuth2 flow.
func (i *identityBase) AuthorizationHandler(r *http.Request) (http.HandlerFunc, error) {
	err := i.parseAuthorizationRequestURL(r.URL.RawQuery)
	if err != nil {
		return nil, err
	}

	ctx := appengine.NewContext(r)
	if err = i.validateAuthorizationRequest(); err != nil {
		log.Infof(ctx, "%+v\n%+v\n%+v", i.authParam, i.redirectParam, i.config.ClientID)
		return nil, err
	}

	i.redirectParam.State = i.authParam.State

	authURL := i.authURL()

	handler := func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, authURL, http.StatusFound)
	}
	return handler, nil
}

func (i *identityBase) ExchangeHandler(r *http.Request) (http.HandlerFunc, error) {
	return nil, nil
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
