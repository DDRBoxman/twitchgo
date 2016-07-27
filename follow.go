package twitch

import (
	"fmt"
	"time"
)

type Follow struct {
	CreatedAt     time.Time `json:"created_at"`
	Id            string    `json:"_id"`
	User          User      `json:"user"`
	Notifications bool      `json:"notifications"`
}

type FollowResponse struct {
	Total   int64    `json:"_total"`
	Follows []Follow `json:"follows"`
}

func (client *TwitchClient) GetChannelFollows(channel string, options *RequestOptions) FollowResponse {
	res := FollowResponse{}
	client.getRequest(fmt.Sprintf("/channels/%s/follows", channel), options, &res)
	return res
}
