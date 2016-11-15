package twitch

import (
	"fmt"
	"strings"
	"time"
)

type Stream struct {
	CreatedAt time.Time `json:"created_at"`
	Id        int64     `json:"_id"`
	Viewers   int64     `json:"viewers"`
	Game      string    `json:"game"`
	Channel   Channel   `json:"channel"`
}

type StreamResponse struct {
	Total  int64   `json:"_total"`
	Stream *Stream `json:"stream"`
}

type StreamsResponse struct {
	Total   int64    `json:"_total"`
	Streams []Stream `json:"streams"`
}

func (client *TwitchClient) GetChannelStream(channel string, options *RequestOptions) (StreamResponse, error) {
	res := StreamResponse{}
	err := client.getRequest(fmt.Sprintf("/streams/%s", channel), options, &res)
	return res, err
}

func (client *TwitchClient) GetChannelsStream(channels ...string) (StreamsResponse, error) {
	res := StreamsResponse{}

	channelsString := strings.Join(channels, ",")
	err := client.getRequest(fmt.Sprintf("/streams?limit=%d&channel=%s", len(channels), channelsString), nil, &res)
	return res, err
}
