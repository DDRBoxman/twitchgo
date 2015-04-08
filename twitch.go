package twitch

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const baseUrl = "https://api.twitch.tv/kraken/"

type TwitchClient struct {
	httpClient *http.Client
}

func NewTwitchClient(httpClient *http.Client) TwitchClient {
	return TwitchClient{
		httpClient: httpClient,
	}
}

func (client *TwitchClient) getRequest(endpoint string, v interface{}) error {
	url := baseUrl + endpoint
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
