package twitch

import (
	"fmt"
	"time"
)

type Channel struct {
	Mature                       bool      `json:"mature"`
	Status                       string    `json:"status"`
	BroadcasterLanguage          string    `json:"broadcaster_language"`
	BroadcasterType              string    `json:"broadcaster_type"`
	DisplayName                  string    `json:"display_name"`
	Game                         string    `json:"game"`
	Delay                        int64     `json:"delay"`
	Language                     string    `json:"language"`
	Id                           int64     `json:"_id"`
	Name                         string    `json:"name"`
	CreatedAt                    time.Time `json:"created_at"`
	UpdatedAt                    time.Time `json:"updated_at"`
	Logo                         string    `json:"logo"`
	Banner                       string    `json:"banner"`
	VideoBanner                  string    `json:"video_banner"`
	ProfileBanner                string    `json:"profile_banner"`
	ProfileBannerBackgroundColor string    `json:"profile_banner_background_color"`
	Partner                      bool      `json:"partner"`
	Views                        int64     `json:"views"`
	Followers                    int64     `json:"followers"`
}

// GetChannel requests channel information from Twitch.
// It returns a Channel struct if successful and any error encountered.
func (client *TwitchClient) GetChannel(channel string) (*Channel, error) {
	res := Channel{}

	err := client.getRequest(fmt.Sprintf("/channels/%s", channel), nil, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Versioned Channel Responses
// ===========================

type CommonChannelFields struct {
	Mature                       bool      `json:"mature"`
	Status                       string    `json:"status"`
	BroadcasterLanguage          string    `json:"broadcaster_language"`
	BroadcasterType              string    `json:"broadcaster_type"`
	DisplayName                  string    `json:"display_name"`
	Game                         string    `json:"game"`
	Delay                        int64     `json:"delay"`
	Language                     string    `json:"language"`
	Id                           int64     `json:"_id"`
	Name                         string    `json:"name"`
	CreatedAt                    time.Time `json:"created_at"`
	UpdatedAt                    time.Time `json:"updated_at"`
	Logo                         string    `json:"logo"`
	Banner                       string    `json:"banner"`
	VideoBanner                  string    `json:"video_banner"`
	ProfileBanner                string    `json:"profile_banner"`
	ProfileBannerBackgroundColor string    `json:"profile_banner_background_color"`
	Partner                      bool      `json:"partner"`
	Views                        int64     `json:"views"`
	Followers                    int64     `json:"followers"`
}

type TwitchChannel interface{
	// GetBroadcasterType can be one of [partner|affiliate|'']
	GetBroadcasterType() string
}

// Twtich API Channel v3 response
// ------------------------------
type v3Channel struct {
	CommonChannelFields
	Id int64 `json:"_id"`
}

// GetBroadcasterType returns 'partner' if the API response has the `Partner`
// field set, 'unknown' otherwise since we cannot distinguish between
// affiliates and regular streamers.
func (c *v3Channel) GetBroadcasterType() string {
	if c.Partner {
		return "partner"
	}
	return "unknown"
}

// Twtich API Channel v5 response
// ------------------------------
type v5Channel struct {
	Channel
	Id string `json:"_id"`
	BroadcasterType string `json:"broadcaster_type"`
}

// GetBroadcasterType returns the API `broadcaster_type` response.
func (c *v5Channel) GetBroadcasterType() string {
	return c.BroadcasterType
}

// Versioned API Functions
// -----------------------

// GetChannelForName requests channel information for the given channel name.
// It returns a v3Channel struct if successful and any error encountered.
func (client *TwitchClient) GetChannelForName(channelName string) (TwitchChannel, error) {
	res := v5Channel{}

	options := &RequestOptions{ Version: "3" }

	err := client.getRequest(fmt.Sprintf("/channels/%s", channelName), options, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// GetChannelForId requests channel information for the given channel ID.
// It returns a Channel struct if successful and any error encountered.
func (client *TwitchClient) GetChannelForId(channelId int64) (TwitchChannel, error) {
	res := v5Channel{}

	options := &RequestOptions{ Version: "5" }

	err := client.getRequest(fmt.Sprintf("/channels/%d", channelId), options, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}