package twitch

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/google/go-querystring/query"
)

const baseUrl = "https://api.twitch.tv/kraken/"

type TwitchClient struct {
	httpClient *http.Client
}

type RequestOptions struct {
	Limit     int64  `url:"limit"`
	Offset    int64  `url:"offset"`
	Direction string `url:"direction"`
}

func NewTwitchClient(httpClient *http.Client) TwitchClient {
	return TwitchClient{
		httpClient: httpClient,
	}
}

func (client *TwitchClient) getRequest(endpoint string, options *RequestOptions, v interface{}) error {
	url := baseUrl + endpoint

	if options != nil {
		v, err := query.Values(options)
		if err != nil {
			return err
		}
		url += "?" + v.Encode()
	}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/vnd.twitchtv.v3+json")
	res, err := client.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}

	return nil
}
