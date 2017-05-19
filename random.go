package main

import (
	"time"
	"fmt"
	"strings"
	"math/rand"
)

func RandomResponses(msg BotMessage) {
	msgs := []string{"Yes", "No", "Of course", "Ew, no", ":nauseated_face:", "YES", "Oh my god yes", "I wish", "Uh, no", "/me cringes", "/tableflip", "/shrug"}
	rand.Seed(time.Now().Unix()) // initialize random generator
	message := fmt.Sprint(msgs[rand.Intn(len(msgs))])
	if strings.HasPrefix(msg.Message, "wouldyou") || strings.HasPrefix(msg.Message, "haveyou") || strings.HasPrefix(msg.Message, "willyou") {
		msg.SendMessage(message)
	}
}