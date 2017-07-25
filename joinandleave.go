package main

func join(msg BotMessage) {
	if msg.Message == "join" {
		if len(msg.Args) > 2 {
			msg.SendMessage("Too many arguments")
			return
		}
		if HasPermissionUser(msg.Author.Name) {
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

			for _, vs := range g.VoiceStates {
				if vs.UserID == msg.Author.ID {
					msg.Author.session.ChannelVoiceJoin(g.ID, vs.ChannelID, false, true)
					msg.SendMessage("Joining your channel")
					return
				}
			}
		}
	}
}

func leave(msg BotMessage) {
	if msg.Message == "leave" {
		if len(msg.Args) > 0 {
			msg.SendMessage("Too many arguments")
			return
		}
		if HasPermissionUser(msg.Author.Name) {
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
			for _, vs := range g.VoiceStates {
				if vs.UserID == msg.Author.ID {
					vc, err := msg.Author.session.ChannelVoiceJoin(g.ID, vs.ChannelID, false, true)
					if err != nil{return}
					vc.Disconnect()
					msg.SendMessage("Leaving your channel")
					return
				}
			}
		}
	}
}
