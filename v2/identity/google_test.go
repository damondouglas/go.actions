package identity

import (
	"net/url"
	"testing"
)

func TestParseAuthorizationRequestURL(t *testing.T) {
	clientID := "995126499590-279e7alr4e74mbkdn3j46vua0pa1k0h4.apps.googleusercontent.com"
	projectID := "gallery-67d5b"
	originURL := url.URL{
		Scheme: "https",
		Host:   redirectHost,
		Path:   redirectBasePath + "/" + projectID,
	}
	rawurl := "/auth?redirect_uri=https%3A%2F%2Foauth-redirect.googleusercontent.com%2Fr%2Fgallery-67d5b&client_id=995126499590-279e7alr4e74mbkdn3j46vua0pa1k0h4.apps.googleusercontent.com&response_type=code&state=AB8b_TPxYUElE2nnblH8mg64KGA66Mb2f_wExsef4ASPvKW3KByLF2uGoEWxBiptwNk7IYWeU-BJ2ZbrHkDVDPFr838_0W3kSuN2iIFYVs-lc7DbT81YABMAaAag0qvsJDUUMuxRNjjYlUhwKC_gz1CZMUGCf7kWEWhJGqXkjJDpI9k5XLiSrHJ7anNEzt2xn8diYR2TsyP4Q8tbrbiT7n0jHUoSydFHufK2aOrwYPcirH1QHGjbdYsdXzC1SMrbFe0yCfuhaICkFlY4hO7rVrsntW-M4TKfXPYWpiRy1TCZifOgQ8MG2haeEz1Nl04_l3r0ua4tjajf9c7XewtB5fY4Idzqa2Q1eyoSoAtCUSGkTUagyWTNKae5FFOYkLzovxTf8n3y28VVZMn7Omj1L8b3LTH8YS7poLgjdNM8OlwkP2aSE8r7Fw3ompCjjf7IzOU9xyUIz0OAWKoziAyDM5ZDlEk7No9UlW_BbpLIm-j2pIB-4GyGK4bwmo4oVhiPhxxuVMRFc1YTMPvGJRiBAQXQvJ6ZvGA_DVq5gcgA3FXxeSaaQjudf3vbNHKdCh3-mbSf6aJ1fYZz7SSUxTXn_gCPSWcZZ1fXKlWRbwTzGRDMYUnPwtuHkAyPP_g40ckpVl99hGUrTPTktj2WVDJybPucX5QgMUfA-NXOlzPI2ZnKpZ5n3EGIfLY&scope=email%20https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fcalendar.events.readonly"
	u, err := url.Parse(rawurl)
	if err != nil {
		t.Error(err)
	}
	param, err := ParseAuthorizationRequestURL(u.RawQuery)
	if err != nil {
		t.Error(err)
	}

	var want, got string

	want = clientID
	got = param.ClientID
	if want != got {
		t.Errorf("got: %s, want: %s", got, want)
	}

	want = originURL.String()
	got = param.RedirectURI
	if want != got {
		t.Errorf("got: %s, want: %s", got, want)
	}

	want = "AB8b_TPxYUElE2nnblH8mg64KGA66Mb2f_wExsef4ASPvKW3KByLF2uGoEWxBiptwNk7IYWeU-BJ2ZbrHkDVDPFr838_0W3kSuN2iIFYVs-lc7DbT81YABMAaAag0qvsJDUUMuxRNjjYlUhwKC_gz1CZMUGCf7kWEWhJGqXkjJDpI9k5XLiSrHJ7anNEzt2xn8diYR2TsyP4Q8tbrbiT7n0jHUoSydFHufK2aOrwYPcirH1QHGjbdYsdXzC1SMrbFe0yCfuhaICkFlY4hO7rVrsntW-M4TKfXPYWpiRy1TCZifOgQ8MG2haeEz1Nl04_l3r0ua4tjajf9c7XewtB5fY4Idzqa2Q1eyoSoAtCUSGkTUagyWTNKae5FFOYkLzovxTf8n3y28VVZMn7Omj1L8b3LTH8YS7poLgjdNM8OlwkP2aSE8r7Fw3ompCjjf7IzOU9xyUIz0OAWKoziAyDM5ZDlEk7No9UlW_BbpLIm-j2pIB-4GyGK4bwmo4oVhiPhxxuVMRFc1YTMPvGJRiBAQXQvJ6ZvGA_DVq5gcgA3FXxeSaaQjudf3vbNHKdCh3-mbSf6aJ1fYZz7SSUxTXn_gCPSWcZZ1fXKlWRbwTzGRDMYUnPwtuHkAyPP_g40ckpVl99hGUrTPTktj2WVDJybPucX5QgMUfA-NXOlzPI2ZnKpZ5n3EGIfLY"
	got = param.State
	if want != got {
		t.Errorf("got: %s, want: %s", got, want)
	}
}

func TestIsAuthorizationRequestValid(t *testing.T) {
	rawurl := "/auth?redirect_uri=https%3A%2F%2Foauth-redirect.googleusercontent.com%2Fr%2Fmyprojectid&client_id=someclientid.apps.googleusercontent.com&response_type=code&state=AB8b_TPxYUElE2nnblH8mg64KGA66Mb2f_wExsef4ASPvKW3KByLF2uGoEWxBiptwNk7IYWeU-BJ2ZbrHkDVDPFr838_0W3kSuN2iIFYVs-lc7DbT81YABMAaAag0qvsJDUUMuxRNjjYlUhwKC_gz1CZMUGCf7kWEWhJGqXkjJDpI9k5XLiSrHJ7anNEzt2xn8diYR2TsyP4Q8tbrbiT7n0jHUoSydFHufK2aOrwYPcirH1QHGjbdYsdXzC1SMrbFe0yCfuhaICkFlY4hO7rVrsntW-M4TKfXPYWpiRy1TCZifOgQ8MG2haeEz1Nl04_l3r0ua4tjajf9c7XewtB5fY4Idzqa2Q1eyoSoAtCUSGkTUagyWTNKae5FFOYkLzovxTf8n3y28VVZMn7Omj1L8b3LTH8YS7poLgjdNM8OlwkP2aSE8r7Fw3ompCjjf7IzOU9xyUIz0OAWKoziAyDM5ZDlEk7No9UlW_BbpLIm-j2pIB-4GyGK4bwmo4oVhiPhxxuVMRFc1YTMPvGJRiBAQXQvJ6ZvGA_DVq5gcgA3FXxeSaaQjudf3vbNHKdCh3-mbSf6aJ1fYZz7SSUxTXn_gCPSWcZZ1fXKlWRbwTzGRDMYUnPwtuHkAyPP_g40ckpVl99hGUrTPTktj2WVDJybPucX5QgMUfA-NXOlzPI2ZnKpZ5n3EGIfLY&scope=email%20https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fcalendar.events.readonly"
	u, err := url.Parse(rawurl)
	if err != nil {
		t.Error(err)
	}
	param, err := ParseAuthorizationRequestURL(u.RawQuery)
	if err != nil {
		t.Error(err)
	}
	got := ValidateAuthorizationRequest(param, "myprojectid", "someclientid.apps.googleusercontent.com")
	if got != nil {
		t.Errorf("got: %v\n%+v", got, param)
	}
}
