package main

func Help(msg BotMessage) {
	if msg.Message == "help" {
		msg.SendMessage("---Introduction---")
		msg.SendMessage("Discord Bot made in Go by N.")
		msg.SendMessage("Every command starts with a '>'")
		msg.SendMessage("For all commands, visit this Gist: https://gist.github.com/Noy/54c439a0f2d577f29ba6985d79bdb9be")
	}
}
