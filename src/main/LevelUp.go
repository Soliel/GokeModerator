package main

import(
	"github.com/Soliel/CommandingDiscord"
	"fmt"
)

func LevelUp(context CommandingDiscord.Context) {
	fmt.Print("Levelup Called")
	context.Session.ChannelMessageSend(context.Channel.ID, "Hi~")
}
