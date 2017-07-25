package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"github.com/bwmarrin/discordgo"
	"strings"
	"log"
)

const (
	Version = "Version: 1.1"
	Prefix  = ">"
)

var (
	count  = 0
	Config BotConfig
	Perm   Permission
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
	if msg.Author.ID == botId { return }

	bingo(msg)

	if !strings.HasPrefix(msg.Message, Prefix) { return }

	//msg.Message = msg.Message[len(Prefix):]

	args := strings.Fields(msg.Message[len(Prefix):])

	command, args := args[0], args[1:]

	msg.Message = command

	msg.Args = args

	registerCommands(msg)

	log.Println(msg.Author.Name, ": ", msg.Message)
}

// Utils

func util(msg BotMessage) {
	if msg.Message == "source" {
		msg.SendMessage("I'm on GitHub! https://github.com/Noy/DiscordBotGo")
		return
	}
	if msg.Message == "ping" {
		msg.SendMessage("It works!")
		return
	}

	// testing
	if msg.Message == "usertest" {
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
		// loop through the members
		for _, member := range g.Members {
			// mention the user (if you want to just say their username, it's User.Username
			msg.SendMessage(member.User.Mention() + " is their name")
			fmt.Println(member.User.Email)
		}
	}
}

func registerCommands(msg BotMessage) {
	randomResponses(msg)
	help(msg)
	join(msg)
	kickCommand(msg)
	addPermissionFor(msg)
	removePermissionFor(msg)
	util(msg)
	cat(msg)
	search(msg)
	leave(msg)
	//announcer(msg)
}

//TODO this
func kickCommand(msg BotMessage) {
	if HasPermissionUser(msg.Author.Name) {
		if msg.Message == "kick" {
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