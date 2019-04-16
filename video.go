package twitch

import (
	"fmt"
	"net/url"
	"time"
)

type Video struct {
	CreatedAt time.Time `json:"created_at"`
	Id        string    `json:"_id"`
	Views     int64     `json:"views"`
	Game      string    `json:"game"`
	Title     string    `json:"title"`
	Channel   Channel   `json:"channel"`
	URL       string    `json:"url"`
	Status    string    `json:"status"`
}

type VideosResponse struct {
	Total  int64   `json:"_total"`
	Videos []Video `json:"videos"`
}

func (client *TwitchClient) GetChannelVideos(channel string, broadcasts bool, limit int64) (VideosResponse, error) {
	res := VideosResponse{}

	err := client.getRequest(fmt.Sprintf("/channels/%s/videos?limit=%d&broadcasts=%t", channel, limit, broadcasts), nil, &res)
	return res, err
}

type HelixVideosResponse struct {
	Videos     []HelixVideo `json:"data"`
	Pagination map[string]string
}

type HelixVideo struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	UserName     string    `json:"user_name"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	PublishedAt  time.Time `json:"published_at"`
	URL          string    `json:"url"`
	ThumbnailURL string    `json:"thumbnail_url"`
	Viewable     string    `json:"viewable"`
	ViewCount    int64     `json:"view_count"`
	Language     string    `json:"language"`
	Type         string    `json:"type"`
	Duration     string    `json:"duration"`
}

/*
GetVideosForID requests video info for a user/channel ID.

https://dev.twitch.tv/docs/api/reference/#get-videos

A pagination cursor may be available in yourResponse.Pagination["cursor"]. This can be supplied in the options.Extra struct with the key "after" on subsequent calls.
*/
func (client *TwitchClient) GetVideosForID(userID string, options *RequestOptions) (res HelixVideosResponse, err error) {
	if options == nil {
		options = &RequestOptions{}
	}
	options.Version = "helix"

	if options.Extra == nil {
		options.Extra = &url.Values{}
	}

	options.Extra.Add("user_id", userID)

	err = client.getRequest("/videos", options, &res)
	return res, err
}
