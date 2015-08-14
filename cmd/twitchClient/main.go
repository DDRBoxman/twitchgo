/*

twitchClient is a simple CLI app that prints a user/channel's data.

Usage:

	twitchClient <channelName>

*/
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ddrboxman/twitchgo"
	"github.com/kr/pretty"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %v <channelName>\n", os.Args[0])
		return
	}

	client := twitch.NewTwitchClient(&http.Client{})

	channel, _ := client.GetChannel(os.Args[1])

	fmt.Printf("%# v\n", pretty.Formatter(channel))
}
