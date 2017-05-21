package main

import "strings"

// this is stupid, i know haha. i was lazy, plan to remake it at one point
func search(msg BotMessage){
	if msg.Message == "images" {
		google := "http://www.google.com/search?tbm=isch&q="
		param := "&source=web&sa=X&ved=0ahUKEwiEr5rZrMLSAhXM8YMKHRVzCJcQ_AUICCgB&biw=1440&bih=799#imgrc=XWXPqrX1RFJiaM:"
		query := strings.Join(msg.Args, "%20")
		msg.SendMessage(google + query + param)
	}
}
