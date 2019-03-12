package twitch

import (
	"fmt"
	"net/url"
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

type helixFollowResponse struct {
	Total   int64         `json:"total"`
	Follows []helixFollow `json:"data"`
}

type helixFollow struct {
	FromID     string    `json:"from_id"`
	FromName   string    `json:"from_name"`
	ToID       string    `json:"to_id"`
	ToName     string    `json:"to_name"`
	FollowedAt time.Time `json:"followed_at"`
}

// GetFollowersForID requests follower information for a user/channel ID.
func (client *TwitchClient) GetFollowersForID(userID string, options *RequestOptions) (helixFollowResponse, error) {
	options.Version = "helix"

	if options.Extra == nil {
		options.Extra = &url.Values{}
	}

	options.Extra.Add("to_id", userID)

	res := helixFollowResponse{}
	err := client.getRequest("/users/follows", options, &res)
	return res, err
}
