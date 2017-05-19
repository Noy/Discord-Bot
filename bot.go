package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"strings"
	"log"
	"github.com/Noy/DiscordBotGo/airhorn"
	"google.golang.org/api/youtube/v3"
)

const (
	Version = "Version: 1.0"
	Prefix  = ">"
)

var (
	count  = 0
	Config BotConfig
	Perm   Permission
	YTService *youtube.Service
)

// GOOS=linux GOARCH=amd64 go build

func main() {
	var err error
	Config, err = loadConfig()
	Perm, err = loadPermFile()
	fmt.Println("These Users currently have permission:", Perm.Users)

	if err != nil {
		log.Println("Could not find config! please create one at ocnfig.toml")
		return
	}

	fmt.Printf(`
	__________.__                     __________        __
	\______   \__| ____    ____   ____\______   \ _____/  |_
 	|    |  _/  |/    \  / ___\ /  _ \|    |  _//  _ \   __\
 	|    |   \  |   |  \/ /_/  >  <_> )    |   (  <_> )  |
 	|______  /__|___|  /\___  / \____/|______  /\____/|__|
        	\/        \//_____/               \/              %-16s`+ "\n\n", Version)

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Config.Token)

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	messagesIn := GetMessages(dg)
	go func() {
		for {
			messageCreate(<-messagesIn, dg.State.User.ID)
		}
	}()

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	dg.UpdateStatus(0, "Bingo.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func messageCreate(msg BotMessage, botId string) {
	if msg.Author.ID == botId {return}

	Bingo(msg)

	if !strings.HasPrefix(msg.Message, Prefix) {return}

	msg.Message = msg.Message[len(Prefix):]

	registerCommands(msg)

	log.Println(msg.Author.Name, ": ", msg.Message)
}

// Utils

func source(msg BotMessage) {
	if msg.Message == "source"{
		msg.SendMessage("I'm on GitHub! https://github.com/Noy/DiscordBotGo")
	}
}

func registerCommands(msg BotMessage) {
	RandomResponses(msg)
	Help(msg)
	airHorn(msg)
	join(msg)
	kickCommand(msg)
	AddPermissionFor(msg)
	source(msg)
}

// TODO

//TODO fix
func join(msg BotMessage) {
	if msg.Message == "join" {
		msg.Author.session.ChannelVoiceJoin("207558416132997122", "315027061028945921", false, true)
		msg.SendMessage("Joining.")
	}
}

//TODO fix
func airHorn(msg BotMessage) {
	if msg.Message == "airhorn" {
		msg.SendMessage("Airhorn activated! Type !airhorn")
		airhorn.Go()
	}
}

//TODO this
func kickCommand(msg BotMessage) {
	if HasPermissionUser(msg.Author.Name) {
		if msg.Message == "nig" {
			sesh := msg.Author.session
			c, err := sesh.State.Channel(string(msg.ChannelID)) // may not work
			if err != nil {
				// could not find channel
				return
			}
			g, err := sesh.State.Guild(c.GuildID)
			if err != nil {
				// could not find guild
				return
			}
			fmt.Println(g.Name)
		}
	}
}
