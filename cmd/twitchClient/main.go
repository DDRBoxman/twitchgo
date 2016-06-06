/*

twitchClient is a simple CLI app that prints a user/channel's data.

Usage:

	twitchClient <channelName>

*/
package main

import (
	"fmt"
	"os"

	"github.com/ddrboxman/twitchgo"
	"github.com/kr/pretty"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %v <channelName>\n", os.Args[0])
		return
	}

	client := twitch.NewTwitchClient("test")
	channel, _ := client.GetChannel(os.Args[1])

	fmt.Printf("%# v\n", pretty.Formatter(channel))
}
