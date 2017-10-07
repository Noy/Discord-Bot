package main

import (
	"github.com/BurntSushi/toml"
	"log"
	"github.com/Noy/DiscordBotGo/utils"
	"os"
	"bufio"
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
	//reloadConfig(user)
	Perm.Users = append(Perm.Users, user)
}

//TODO this function
func reloadConfig(user string) (err error) {
	file, err := os.Create("permission.toml")
	if err != nil {
		return err
	}
	defer file.Close()

	//need to append to the config

	w := bufio.NewWriter(file)

	//w.Write(newUser)

	//for i := 0; i < len(Perm.Users); i++ {
	//	fmt.Println("Users = [\"" ,Perm.Users[i],"\",\"" + user + "\"]")
	//}
	//
	//fmt.Println(strings.Trim(fmt.Sprint(Perm.Users), "[]"))
	//fmt.Println("Users = [\"" + "2\",\"" + user + "\"]")

	w.Flush()
	file.Close()
	return nil
}

func RmPerm(user string) {
	if _, err := toml.DecodeFile("permission.toml", &Perm); err != nil {
		log.Fatal(err)
	}
	Perm.Users = remove(Perm.Users, user)
}

func HasPermissionUser(user string) (bool) {
	for _, u := range Perm.Users {
		return utils.CaseInsensitiveContains(u, user)
	}
	return true
}

func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}