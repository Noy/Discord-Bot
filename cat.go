package main

import (
	"net/http"
	"math/rand"
	"fmt"
)

func cat(msg BotMessage) {
	if msg.Message == "cat" {
		msgs := []string{"png", "gif"}
		resp, err := http.Get("http://thecatapi.com/api/images/get?format=src&type=" + msgs[rand.Intn(len(msgs))])
		if err != nil {
			fmt.Sprintf("error: %d", err)
			return
		}

		msg.SendMessage("http://" + resp.Request.URL.Host + resp.Request.URL.Path)
	}
}
