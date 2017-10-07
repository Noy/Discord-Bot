package main

import (
	"strings"
)

func announce(msg BotMessage) {
	if msg.Message == "announce" {
		if HasPermissionUser(msg.Author.Name) {
			words := strings.Join(msg.Args, " ")
			c, err := msg.Author.session.State.Channel(string(msg.ChannelID))
			if err != nil {
				// Could not find channel.
				return
			}
			g, err := msg.Author.session.State.Guild(c.GuildID)
			if err != nil {
				// Could not find guild.
				return
			}
			msg.Author.session.ChannelMessageDelete(c.ID, msg.ID)
			//msg.SendMessage(words)
			for _, channels := range g.Channels {
				msg.Author.session.ChannelMessageSend(channels.ID, words)
			}
		} else {
			msg.SendMessage("no perms.")
		}
	}
}