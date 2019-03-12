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
	Total      int64             `json:"total"`
	Follows    []helixFollow     `json:"data"`
	Pagination map[string]string `json:"pagination"`
}

type helixFollow struct {
	FromID     string    `json:"from_id"`
	FromName   string    `json:"from_name"`
	ToID       string    `json:"to_id"`
	ToName     string    `json:"to_name"`
	FollowedAt time.Time `json:"followed_at"`
}

/*
GetFollowersForID requests follower information for a user/channel ID.

https://dev.twitch.tv/docs/api/reference/#get-users-follows

Follower data is sorted by Twitch: most recent follower first.

A pagination cursor may be available in yourResponse.Pagination["cursor"]. This can be supplied in the options.Extra struct with the key "after" on subsequent calls.
*/
func (client *TwitchClient) GetFollowersForID(userID string, options *RequestOptions) (helixFollowResponse, error) {
	if options == nil {
		options = &RequestOptions{}
	}
	options.Version = "helix"

	if options.Extra == nil {
		options.Extra = &url.Values{}
	}

	options.Extra.Add("to_id", userID)

	res := helixFollowResponse{}
	err := client.getRequest("/users/follows", options, &res)
	return res, err
}
