package main

import (
	"strings"
	"log"
)

func AddPermissionFor(msg BotMessage) {
	if strings.HasPrefix(msg.Message, "addperm") {
		if HasPermissionUser(msg.Author.Name) {
			name := strings.TrimLeft(msg.Message, "addperm ")
			AddPerm(name)
			log.Println(Perm.Users)
			msg.SendMessage("Added " + name + " to the permissions list.")
		} else {
			msg.SendMessage("You don't have permission")
		}
	}
}
