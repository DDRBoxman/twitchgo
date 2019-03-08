package twitch

import (
	"fmt"
	"net/url"
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

type helixUsersResponse struct {
	Users []HelixUser `json:"data"`
}

type HelixUser struct {
	ID              string `json:"id"`
	Login           string `json:"login"`
	DisplayName     string `json:"display_name"`
	Type            string `json:"type"`
	BroadcasterType string `json:"broadcaster_type"`
	Description     string `json:"description"`
	ProfileImageURL string `json:"profile_image_url"`
	OfflineImageURL string `json:"offline_image_url"`
	ViewCount       int64  `json:"view_count"`
	Email           string `json:"email"`
}

func (user *HelixUser) IsPartnered() bool {
	return user.BroadcasterType == "partner"
}

func (client *TwitchClient) GetUser(username string) (*User, error) {
	twitchUser := &User{}

	err := client.getRequest(fmt.Sprintf("/user/%s", username), nil, twitchUser)
	if err != nil {
		return nil, err
	}

	return twitchUser, nil
}

func (client *TwitchClient) GetUsers(id, login []string) (*[]HelixUser, error) {
	usersResponse := &helixUsersResponse{}

	options := &RequestOptions{
		Version: "helix",
		Extra:   &url.Values{},
	}

	for _, userID := range id {
		options.Extra.Add("id", userID)
	}

	for _, userLogin := range login {
		options.Extra.Add("login", userLogin)
	}

	err := client.getRequest("/users", options, usersResponse)
	if err != nil {
		return nil, err
	}

	return &usersResponse.Users, err
}
