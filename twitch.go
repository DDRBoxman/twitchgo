package twitch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const baseUrl = "https://api.twitch.tv/kraken"

type TwitchClient struct {
	HttpClient *http.Client
	ClientID   string
}

type RequestOptions struct {
	Limit     int64  `url:"limit"`
	Offset    int64  `url:"offset"`
	Direction string `url:"direction"`
	Nonce     int64  `url:"_"`
	Channel   string `url:"channel"`
	Version   string
}

func NewTwitchClient(clientID string) TwitchClient {
	return TwitchClient{
		HttpClient: &http.Client{},
		ClientID:   clientID,
	}
}

func NewTwitchClientWithHTTPClient(clientID string, httpClient *http.Client) TwitchClient {
	return TwitchClient{
		HttpClient: httpClient,
		ClientID:   clientID,
	}
}

func (client *TwitchClient) getRequest(endpoint string, options *RequestOptions, out interface{}) error {
	targetUrl := baseUrl + endpoint
	targetVersion := "3"

	if options != nil {
		v := url.Values{}

		if options.Direction != "" {
			v.Add("direction", options.Direction)
		}

		if options.Limit != 0 {
			v.Add("limit", fmt.Sprintf("%d", options.Limit))
		}

		if options.Offset != 0 {
			v.Add("offset", fmt.Sprintf("%d", options.Offset))
		}

		if options.Nonce != 0 {
			v.Add("_", fmt.Sprintf("%d", options.Nonce))
		}

		if options.Channel != "" {
			v.Add("channel", options.Channel)
		}

		if options.Version != "" {
			targetVersion = options.Version
		}

		if len(v) != 0 {
			targetUrl += "?" + v.Encode()
		}
	}

	req, _ := http.NewRequest("GET", targetUrl, nil)
	req.Header.Set("Accept", fmt.Sprintf("application/vnd.twitchtv.v%s+json", targetVersion))
	req.Header.Set("Client-ID", client.ClientID)
	res, err := client.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Request failed with status: %v", res.StatusCode)
	}

	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, out)
	if err != nil {
		return err
	}

	return nil
}
