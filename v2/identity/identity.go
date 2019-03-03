package identity

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"google.golang.org/appengine"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	actionsRedirectBase = "https://oauth-redirect.googleusercontent.com/r/"
	projectIDKey        = "PROJECT_ID"
	secretPathKey       = "SECRET_PATH"
	testHostKey         = "TEST_HOST"
	scopesKey           = "SCOPES"

	accessTokenKey = "access_token"
	stateKey       = "state"
	redirectURIKey = "redirect_uri"
	codeKey        = "code"
	tokenTypeKey   = "token_type"
	bearerKey      = "bearer"

	// AuthPath is the path to the first authorization flow.
	AuthPath = "auth"
)

// StoreHandler handles token storage.
type StoreHandler func(context.Context, *oauth2.Token)

var (
	app         *appConfig
	oauthConfig *oauth2.Config
)

func init() {
	var err error
	app, err = loadYamlConfig()
	if err != nil {
		log.Fatal(err)
	}

	oauthConfig, err = loadOauthConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type appConfig struct {
	ProjectID  string
	SecretPath string
	TestHost   string
	Scopes     []string
}

func loadYamlConfig() (*appConfig, error) {
	yamlMap := map[string]string{
		projectIDKey:  os.Getenv(projectIDKey),
		secretPathKey: os.Getenv(secretPathKey),
		testHostKey:   os.Getenv(testHostKey),
		scopesKey:     os.Getenv(scopesKey),
	}

	for key, value := range yamlMap {
		if value == "" {
			return nil, fmt.Errorf("%s is not set in app.yaml", key)
		}
	}

	scopes := strings.Split(yamlMap[scopesKey], ",")
	for i := range scopes {
		scopes[i] = strings.Trim(scopes[i], "")
	}

	return &appConfig{
		ProjectID:  yamlMap[projectIDKey],
		SecretPath: yamlMap[secretPathKey],
		TestHost:   yamlMap[testHostKey],
		Scopes:     scopes,
	}, nil
}

func loadOauthConfig() (config *oauth2.Config, err error) {
	data, err := ioutil.ReadFile(app.SecretPath)
	if err != nil {
		return nil, err
	}
	return google.ConfigFromJSON(data, app.Scopes...)
}

func host(r *http.Request) string {
	ctx := appengine.NewContext(r)
	h := "https://" + appengine.DefaultVersionHostname(ctx)
	if appengine.IsDevAppServer() {
		h = app.TestHost
	}
	return h
}

// Handler handles exchange token requests.
type Handler struct {
	Store StoreHandler
	State string
}

// AuthHandler the first step in the OAuth2 flow.
func AuthHandler(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	state := v.Get(stateKey)
	redirectURI := host(r) + "/exch"
	oauthConfig.RedirectURL = redirectURI
	u := oauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	log.Println("URL: ", u)
	http.Redirect(w, r, u, http.StatusFound)
}

// TokenHandler handles token requests.
func (h *Handler) TokenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		h.codeRequest(w, r)
	}
	if r.Method == "POST" {
		h.codePost(w, r)
	}
}

func (h *Handler) codeRequest(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	v := r.URL.Query()
	code := v.Get(codeKey)
	state := v.Get(stateKey)
	h.State = state
	err := h.token(ctx, code, w, r)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *Handler) codePost(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	v, err := url.ParseQuery(string(data))
	if err != nil {
		log.Fatal(err)
	}
	code := v.Get(codeKey)
	err = h.token(ctx, code, w, r)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *Handler) token(ctx context.Context, code string, w http.ResponseWriter, r *http.Request) (err error) {
	tok, err := oauthConfig.Exchange(ctx, code)
	if err != nil {
		return err
	}
	if h.Store != nil {
		h.Store(ctx, tok)
	}
	h.redirect(tok, w, r)
	return nil
}

func (h *Handler) redirect(tok *oauth2.Token, w http.ResponseWriter, r *http.Request) {
	qry, _ := url.ParseQuery("")
	qry.Set(accessTokenKey, tok.AccessToken)
	qry.Set(tokenTypeKey, bearerKey)
	qry.Set(stateKey, h.State)
	u := actionsRedirectBase + app.ProjectID + "#" + qry.Encode()
	// u := "https://oauth-redirect.googleusercontent.com/r/" + app.ProjectID + "#" + "access_token=" + tok.AccessToken + "&token_type=bearer&state=" + h.State
	http.Redirect(w, r, u, http.StatusFound)
}
