package identity

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	oauth2 "google.golang.org/api/oauth2/v2"

	"github.com/damondouglas/go.actions/v2/dialogflow"
	"github.com/damondouglas/go.actions/v2/helper"
)

// User from Google Account Linkage.
type User struct {
	Email       string
	ID          string
	AccessToken string
}

// GoogleIdentity handles Google Account linking.
type GoogleIdentity struct {
	Logger         helper.Logger
	ContextFactory helper.ContextFactory
	ClientFactory  helper.ClientFactory
	ClientID       string
	Storer         helper.Storer
}

// Context from http.Request.
func (g *GoogleIdentity) Context(r *http.Request) (ctx context.Context) {
	if g.ContextFactory != nil {
		ctx = g.ContextFactory.NewContext(r)
	}
	return ctx
}

// Client from context.Context.
func (g *GoogleIdentity) Client(ctx context.Context) (client *http.Client) {
	if ctx != nil {
		client = g.ClientFactory.NewClient(ctx)
	}

	return client
}

// LinkUserHandler handles Google Signin Account linking.
func (g *GoogleIdentity) LinkUserHandler(w http.ResponseWriter, r *http.Request) {
	var ctx context.Context = g.Context(r)
	var req *dialogflow.Request
	req, err := dialogflow.Decode(r.Body)
	g.logError(ctx, err)
	info, err := g.TokenInfo(req, r)
	g.logError(ctx, err)
	if !g.IsValidTokenInfo(info, g.ClientID) {
		g.logError(ctx, fmt.Errorf("token is not valid: want: %s, got %s", g.ClientID, info.Audience))
	} else {
		user := &User{
			Email:       info.Email,
			ID:          req.OriginalDetectIntentRequest.Payload.User.UserID,
			AccessToken: req.OriginalDetectIntentRequest.Payload.User.AccessToken,
		}
		err = g.Storer.Store(user)
		g.logError(ctx, err)
	}
}

func (g *GoogleIdentity) logError(ctx context.Context, err error) {
	if ctx != nil && g.Logger != nil && err != nil {
		g.Logger.Errorf(ctx, "%+v", err)
	}
}

// TokenInfo acquires Oauth2 token info from dialogflow.Request User IDToken.
func (g *GoogleIdentity) TokenInfo(req *dialogflow.Request, r *http.Request) (*oauth2.Tokeninfo, error) {
	ctx := g.Context(r)
	if ctx == nil {
		return nil, errors.New("Context is nil")
	}
	idToken := req.OriginalDetectIntentRequest.Payload.User.IDToken
	client := g.Client(ctx)
	oauth2Service, err := oauth2.New(client)
	if err != nil {
		return nil, err
	}
	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfoCall.IdToken(idToken)
	return tokenInfoCall.Do()
}

// IsValidTokenInfo validates token matched against clientID.
func (g *GoogleIdentity) IsValidTokenInfo(info *oauth2.Tokeninfo, clientID string) bool {
	return info.Audience == clientID
}
