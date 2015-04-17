package twitch

import (
	"fmt"
	"time"
)

type User struct {
	DisplayName string    `json:"display_name"`
	Id          int64     `json:"_id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Bio         string    `json:"bio"`
	Logo        string    `json:"logo"`
	Email       string    `json:"email"`
	Partnered   bool      `json:"partnered"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (client *TwitchClient) GetUser(username string) (*User, error) {
	twitchUser := &User{}

	err := client.getRequest(fmt.Sprintf("/user/%s", username), nil, twitchUser)
	if err != nil {
		return nil, err
	}

	return twitchUser, nil
}
