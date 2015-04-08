package twitch

import (
	"fmt"
)

type User struct {
	DisplayName string `json:"display_name"`
	Id          int64  `json:"_id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Bio         string `json:"bio"`
	Logo        string `json:"logo"`
	Email       string `json:"email"`
	Partnered   bool   `json:"partnered"`
}

func (client *TwitchClient) GetUser(username string) User {
	twitchUser := User{}
	client.getRequest(fmt.Sprintf("/user/%s", username), &twitchUser)
	return twitchUser
}
