package identity

import (
	"net/url"
	"testing"

	"golang.org/x/oauth2"
)

func TestParseAuthorizationRequestURL(t *testing.T) {
	rawurl := "https://myservice.example.com/auth?client_id=GOOGLE_CLIENT_ID&redirect_uri=REDIRECT_URI&state=STATE_STRING&response_type=token"
	u, err := url.Parse(rawurl)
	if err != nil {
		t.Error(err)
	}
	param, err := ParseAuthorizationRequestURL(u.RawQuery)
	if err != nil {
		t.Error(err)
	}

	var want, got string

	want = "GOOGLE_CLIENT_ID"
	got = param.ClientID
	if want != got {
		t.Errorf("got: %s, want: %s", got, want)
	}

	want = "REDIRECT_URI"
	got = param.RedirectURI
	if want != got {
		t.Errorf("got: %s, want: %s", got, want)
	}

	want = "STATE_STRING"
	got = param.State
	if want != got {
		t.Errorf("got: %s, want: %s", got, want)
	}

	want = "token"
	got = param.ResponseType
	if want != got {
		t.Errorf("got: %s, want: %s", got, want)
	}

}

func TestRedirectURL(t *testing.T) {
	config := &oauth2.Config{
		RedirectURL: "REDIRECT_URL",
		Scopes:      []string{"email,profile,calendar"},
	}
	param := &RedirectParameters{
		Offline:   true,
		Force:     true,
		State:     "STATE_STRING",
		ProjectID: "PROJECT_ID",
	}

	result := RedirectURL(config, param)
	t.Error(result)
}
