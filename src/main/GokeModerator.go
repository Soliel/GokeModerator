package main

import (
	"github.com/Soliel/CommandingDiscord"
	"github.com/bwmarrin/discordgo"
	"fmt"

	"io/ioutil"
	"encoding/json"

	"os"
	"os/signal"
	"syscall"
)

var (
	BotID          string
	CommandHandler *CommandingDiscord.CommandHandler
)

type config struct {
	BotToken   string `json:"bot_token"`
	BotPrefix  string `json:bot_prefix`
	Moderators []string `json:mod_list`
}

func main() {

	conf := loadConfig("config.json")

	dg, err := discordgo.New("Bot " + conf.BotToken)
	if err != nil {
		panic(err)
	}

	u, err := dg.User("@me")
	if err != nil {
		panic(err)
	}

	err, CommandHandler = CommandingDiscord.NewCommandHandler()
	if err != nil {
		panic("Unable to make command handler, bot operation ceased.")
	}
	registerCommands()

	BotID = u.ID

	dg.AddHandlerOnce(ready)
	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		panic(err)
	}

	//go CommandHandler.StartCooldownTicker()

	fmt.Println("Bot is now running as user: ", u.Username)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println("Message seen.")
	CommandingDiscord.HandleMessages(s, m, BotID, CommandHandler)
}

func loadConfig(filename string) *config {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error loading config, ", err)
		defaultconf := config{"DefaultToken", "!", []string{"mod1","mod2"}}
		jsonconf, err := json.Marshal(&defaultconf)
		if err != nil {
			panic(err)
		} else {
			ioutil.WriteFile(filename,[]byte(jsonconf), 0644)
			fmt.Println("Default config created. please add bot token.")
			return nil
		}
	}

	var confData config
	err = json.Unmarshal(body, &confData)
	if err != nil {
		fmt.Println("Error parsing JSON data, ", err)
		return nil
	}
	return &confData
}

func registerCommands() {
	CommandHandler.Register("Levelup", LevelUp, 86400)
	fmt.Println("All commands registered.")
}

func ready(s *discordgo.Session, r *discordgo.Ready) {
	fmt.Println("ready event fired.")
}