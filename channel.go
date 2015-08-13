package twitch

import (
	"fmt"
	"time"
)

type Channel struct {
	Mature                       bool      `json:"mature"`
	Status                       string    `json:"status"`
	BroadcasterLanguage          string    `json:"broadcaster_language"`
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
