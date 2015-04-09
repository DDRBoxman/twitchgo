package twitch

import (
	"fmt"
	"time"
)

type Stream struct {
	CreatedAt time.Time `json:"created_at"`
	Id        string    `json:"_id"`
	Viewers   int64     `json:"viewers"`
	Game      string    `json:"game"`
	Channel   Channel   `json:"channel"`
}

type StreamResponse struct {
	Total  int64  `json:"_total"`
	Stream Stream `json:"stream"`
}

func (client *TwitchClient) GetChannelStream(channel string, options *RequestOptions) StreamResponse {
	res := StreamResponse{}
	client.getRequest(fmt.Sprintf("/streams/%s", channel), options, &res)
	return res
}
