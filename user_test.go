package twitch_test

import (
	"net/url"
	"net/http"
	"path"
	"testing"
	"github.com/DDRBoxman/twitchgo"
	"net/http/httptest"
	"fmt"
	"log"
)

// RewriteTransport is an http.RoundTripper that rewrites requests
// using the provided URL's Scheme and Host, and its Path as a prefix.
// The Opaque field is untouched.
// If Transport is nil, http.DefaultTransport is used
type RewriteTransport struct {
	Transport http.RoundTripper
	URL       *url.URL
}

func (t RewriteTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// note that url.URL.ResolveReference doesn't work here
	// since t.u is an absolute url
	req.URL.Scheme = t.URL.Scheme
	req.URL.Host = t.URL.Host
	req.URL.Path = path.Join(t.URL.Path, req.URL.Path)
	rt := t.Transport
	if rt == nil {
		rt = http.DefaultTransport
	}
	return rt.RoundTrip(req)
}

var usersResponse = `
{"data":[{
   "id":"44322889",
   "login":"dallas",
   "display_name":"dallas",
   "type":"staff",
   "broadcaster_type":"",
   "description":"Just a gamer playing games and chatting. :)",
   "profile_image_url":"https://static-cdn.jtvnw.net/jtv_user_pictures/dallas-profile_image-1a2c906ee2c35f12-300x300.png",
   "offline_image_url":"https://static-cdn.jtvnw.net/jtv_user_pictures/dallas-channel_offline_image-1a2c906ee2c35f12-1920x1080.png",
   "view_count":191836881,
   "email":"login@provider.com"
}]}
`

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", usersResponse)
}

func TestHelixUsers(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(Handler))

	testServerURL, err := url.Parse(testServer.URL)
	if err != nil {
		log.Fatalln("failed to parse httptest.Server URL:", err)
	}

	transport := RewriteTransport{
		URL: testServerURL,
	}

	client := &http.Client{
		Transport: transport,
	}

	twitchClient := twitch.NewTwitchClientWithHTTPClient("blah", client)

	users, err := twitchClient.GetUsers(nil, []string{"dallas"})

	if len(*users) != 1 {
		t.Errorf("Len of users was %d expected 1", len(*users))
	}

	if (*users)[0].Login != "dallas" {
		t.Errorf("Len of users was %s expected dallas", (*users)[0].Login)
	}
}