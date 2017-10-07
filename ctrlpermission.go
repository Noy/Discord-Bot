package main

import (
	"log"
)

func addPermissionFor(msg BotMessage) {
	if msg.Message == "addperm" {
		if HasPermissionUser(msg.Author.Name) {
			name := msg.Args[0]
			AddPerm(name)
			log.Println(Perm.Users)
			msg.SendMessage("Added " + name + " to the permission list.")
		} else {
			msg.SendMessage("You do not have permission.")
		}
	}
}

func removePermissionFor(msg BotMessage) {
	if msg.Message == "rmperm" {
		if HasPermissionUser(msg.Author.Name) {
			name := msg.Args[0]
			RmPerm(name)
			log.Println(Perm.Users)
			msg.SendMessage("Removed " + name + " to the permissions list.")
		} else {
			msg.SendMessage("You don't have permission")
		}
	}
}
