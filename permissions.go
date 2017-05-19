package main

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Permission struct {
	Users []string
}

func loadPermFile() (perm Permission, err error) {
	_, err = toml.DecodeFile("permission.toml", &perm)
	return
}


func AddPerm(user string) {
	if _, err := toml.DecodeFile("permission.toml", &Perm); err != nil {
		log.Fatal(err)
	}
	Perm.Users = append(Perm.Users, user)
}

func HasPermissionUser(user string) (bool) {
	return contains(Perm.Users, user)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {return true}
	}
	return false
}