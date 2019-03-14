package twitch

import (
	"fmt"
	"net/url"
	"time"
)

type Subscription struct {
	CreatedAt time.Time `json:"created_at"`
	Id        string    `json:"_id"`
	User      User      `json:"user"`
}

type SubscriptionResponse struct {
	Total         int64          `json:"_total"`
	Subscriptions []Subscription `json:"subscriptions"`
}

func (client *TwitchClient) GetChannelSubscriptions(channel string, options *RequestOptions) SubscriptionResponse {
	res := SubscriptionResponse{}
	client.getRequest(fmt.Sprintf("/channels/%s/subscriptions", channel), options, &res)
	return res
}

type helixSubscription struct {
	BroadcasterID   string `json:"broadcaster_id"`
	BroadcasterName string `json:"broadcaster_name"`
	IsGift          bool   `json:'is_gift'`
	Tier            string `json:"tier"`
	PlanName        string `json:"plan_name"`
	SubscriberID    string `json:"user_id"`
	SubscriberName  string `json:"user_name"`
}

type helixSubscriptionResponse struct {
	Subscriptions []helixSubscription `json:"data"`
	Pagination    map[string]string
}

/*
GetSubscribersForID requests subscriber info for a user/channel ID.

https://dev.twitch.tv/docs/api/reference/#get-broadcaster-subscriptions

Note: Twitch requires an OAuth access token with the `channel:read:subscriptions` scope in order to access subscriber information for the matching user.

A pagination cursor may be available in yourResponse.Pagination["cursor"]. This can be supplied in the options.Extra struct with the key "after" on subsequent calls.
*/
func (client *TwitchClient) GetSubscribersForID(userID string, options *RequestOptions) (res helixSubscriptionResponse, err error) {
	if options == nil {
		options = &RequestOptions{}
	}
	options.Version = "helix"

	if options.Extra == nil {
		options.Extra = &url.Values{}
	}

	options.Extra.Add("broadcaster_id", userID)

	res = helixSubscriptionResponse{}
	err = client.getRequest("/subscriptions", options, &res)
	return res, err
}
