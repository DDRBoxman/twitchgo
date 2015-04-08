package twitch

import (
	"fmt"
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

func (client *TwitchClient) GetChannelSubscriptions(channel string) SubscriptionResponse {
	res := SubscriptionResponse{}
	client.getRequest(fmt.Sprintf("/channels/%s/subscriptions", channel), &res)
	return res
}
