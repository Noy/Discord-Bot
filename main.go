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
	Version = "Version: 1.0"
	Prefix = ">"
)

var (
	count = 0
	Config BotConfig
)

// GOOS=linux GOARCH=amd64 go build

func main() {
	var err error
	Config, err = loadConfig()
	if err != nil {
		log.Println("could not find config! please create one at ocnfig.toml")
		return
	}

	fmt.Printf(`
	__________.__                     __________        __
	\______   \__| ____    ____   ____\______   \ _____/  |_
 	|    |  _/  |/    \  / ___\ /  _ \|    |  _//  _ \   __\
 	|    |   \  |   |  \/ /_/  >  <_> )    |   (  <_> )  |
 	|______  /__|___|  /\___  / \____/|______  /\____/|__|
        	\/        \//_____/               \/              %-16s`+"\n\n", Version)

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " +  Config.Token)

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
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

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	dg.UpdateStatus(1, "Bingo.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(msg BotMessage, botId string) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if msg.Author.ID == botId {return}
	Bingo(msg)
	if HasPermissionUser(msg.Author) {
		// If it doesn't have the prefix, ignore the message
		if !strings.HasPrefix(msg.Message, Prefix) {return}

		msg.Message = msg.Message[len(Prefix):]

		registerCommands(msg)

		// Register all commands

		// I want to see the bot's response too.
		log.Println(msg.Author, ": ", msg.Message)
	} else {
		msg.SendMessage("Sorry, you do not have permission!")
	}
}

// Utils

func registerCommands(msg BotMessage) {
	RandomResponses(msg)
	Help(msg)
	airHorn(msg)
	join(msg)
	kickCommand(msg)
}

//TODO fix
func join(msg BotMessage) {
	if msg.Message == "join" {
		msg.Author.session.ChannelVoiceJoin("207558416132997122", "315027061028945921", false, true)
		msg.SendMessage("Joining.")
	}
}

//TODO fix
func airHorn(msg BotMessage)  {
	if msg.Message == "airhorn" {
		msg.SendMessage("Airhorn activated! Type !airhorn")
		//airhorn.Go()
	}
}

//TODO this
func kickCommand(msg BotMessage) {
	if HasPermissionUser(msg.Author) {
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

//TODO this
func permissionCommand(msg BotMessage) {
	if msg.Author.Name != "Noy" {msg.SendMessage("You do not have permission!")}

}