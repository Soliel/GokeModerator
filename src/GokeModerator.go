package main

import (
	"github.com/bwmarrin/discordgo"
	"fmt"
)

var (
	BotID string
)

func main() {
	dg, err := discordgo.New("Bot" + "MjEwNTU3NjQ5ODE5OTI2NTMx.DXiWWQ.AGhPFJFLxZsB_nF18Z8IS9lQzXM")
	if err != nil {
		fmt.Println("unable to authenticate with discord, ", err)
		panic(err)
	}

	u, err := dg.User("@me")
	if err != nil {
		fmt.Println("Error obtaining account details, ", err)
	}

	BotID = u.ID

	err = dg.Open()
	if err != nil {
		fmt.Println("Error establishing connection with guild(s), ", err)
	}

	fmt.Println("Bot is now running as user: ", u.Username)

	<- make(chan struct{})
	return
}