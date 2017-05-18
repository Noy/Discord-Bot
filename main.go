package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"strings"
	"log"
	"time"
	"math/rand"
	"github.com/Noy/DiscordBotGo/airhorn"
)

const (
	Version = "Version: 1.0"
	Prefix = ">"
)

var (
	count = 0
	Config BotConfig
)

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

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
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
	if msg.Author.Name == "Noy" || msg.Author.Name == "Owen" {
		bingo(msg)

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
	randomResponses(msg)
	helpCommand(msg)
	airHorn(msg)
}

// Commands

func randomResponses(msg BotMessage) {
	msgs := []string{"Yes", "No", "Of course", "Ew, no", ":nauseated_face:", "YES", "Oh my god yes", "I wish", "Uh, no", "/me cringes", "/tableflip", "/shrug"}
	rand.Seed(time.Now().Unix()) // initialize random generator
	message := fmt.Sprint(msgs[rand.Intn(len(msgs))])
	if strings.HasPrefix(msg.Message, "wouldyou") || strings.HasPrefix(msg.Message, "haveyou") || strings.HasPrefix(msg.Message, "willyou") {
		msg.SendMessage(message)
	}
}

func bingo(msg BotMessage) {
	if CaseInsensitiveContains(msg.Message, "bingo") {
		count++
		msg.SendMessagef(`The word "bingo" has been said %d times`, count)
	}
}

func CaseInsensitiveContains(a, b string) bool {
	return strings.Contains(strings.ToLower(a), strings.ToLower(b))
}

func helpCommand(msg BotMessage) {
	if msg.Message == "help" {
		msg.SendMessage("---Introduction---")
		msg.SendMessage("Discord Bot made in Go by N.")
		msg.SendMessage("Every command starts with a '>'")
		msg.SendMessage("Some commands are !kick , !google , !topic , etc..")
		msg.SendMessage("For all commands, visit this Gist: https://gist.github.com/Noy/54c439a0f2d577f29ba6985d79bdb9be")
	}
}

func airHorn(msg BotMessage)  {
	if msg.Message == "airhorn" {
		msg.SendMessage("Airhorn activated! Type !airhorn")
		airhorn.Go()
	}
}