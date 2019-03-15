package twitch

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

type Stream struct {
	CreatedAt time.Time `json:"created_at"`
	Id        int64     `json:"_id"`
	Viewers   int64     `json:"viewers"`
	Game      string    `json:"game"`
	Channel   Channel   `json:"channel"`

	// StreamType is only avaliable when querying with the V5 endpoint.
	StreamType string `json:"stream_type"`
}

type StreamResponse struct {
	Total  int64   `json:"_total"`
	Stream *Stream `json:"stream"`
}

type StreamsResponse struct {
	Total   int64    `json:"_total"`
	Streams []Stream `json:"streams"`
}

func (client *TwitchClient) GetChannelStream(channel string, options *RequestOptions) (StreamResponse, error) {
	res := StreamResponse{}
	err := client.getRequest(fmt.Sprintf("/streams/%s", channel), options, &res)
	return res, err
}

func (client *TwitchClient) GetChannelsStream(channels ...string) (StreamsResponse, error) {
	res := StreamsResponse{}

	channelsString := strings.Join(channels, ",")
	err := client.getRequest(fmt.Sprintf("/streams?limit=%d&channel=%s", len(channels), channelsString), nil, &res)
	return res, err
}

func (client *TwitchClient) GetChannelsStreamV5(channelIDs ...string) (StreamsResponse, error) {
	res := StreamsResponse{}
	options := RequestOptions{
		Version: "5",
	}

	channelsString := strings.Join(channelIDs, ",")
	err := client.getRequest(fmt.Sprintf("/streams?limit=%d&channel=%s", len(channelIDs), channelsString), &options, &res)
	return res, err
}

type helixStream struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	UserName     string    `json:"user_name"`
	GameID       string    `json:"game_id"`
	CommunityIDs []string  `json:"community_ids"`
	Type         string    `json:"type"`
	Title        string    `json:"title"`
	ViewerCount  int       `json:"viewer_count"`
	StartedAt    time.Time `json:"started_at"`
	Language     string    `json:"language"`
	ThumbnailURL string    `json:"thumbnail_url"`
}

type helixStreamsResponse struct {
	Streams    []helixStream     `json:"data"`
	Pagination map[string]string `json:"pagination"`
}

/*
GetStreamsForIDs requests stream info for one or more channels.

https://dev.twitch.tv/docs/api/reference/#get-streams

A pagination cursor may be available in yourResponse.Pagination["cursor"]. This can be supplied in the options.Extra struct with the key "after" on subsequent calls.
*/
func (client *TwitchClient) GetStreamsForIDs(options *RequestOptions, channelIDs ...string) (res helixStreamsResponse, err error) {
	if options == nil {
		options = &RequestOptions{}
	}

	options.Version = "helix"

	if options.Extra == nil {
		options.Extra = &url.Values{}
	}

	for _, channel := range channelIDs {
		options.Extra.Add("user_id", channel)
	}

	err = client.getRequest(fmt.Sprintf("/streams"), options, &res)
	return res, err
}
