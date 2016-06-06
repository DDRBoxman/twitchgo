package twitch

import (
	"fmt"
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
