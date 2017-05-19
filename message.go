package main

import (
	"github.com/bwmarrin/discordgo"
	"strconv"
	"errors"
	"fmt"
)

type Target interface {
	SendMessage(message interface{}) error
}

type MessageDestination string

type User struct {
	ID string
	Name string
	ChannelID MessageDestination
	session *discordgo.Session
}

type BotMessage struct {
	ID        string
	Message   string
	ChannelID MessageDestination
	Author    *User
}

type MessageReaction struct {
	Message BotMessage
	Reactor *User
}

func GetMessages(session *discordgo.Session) (msgsOut <-chan BotMessage) {
	messages := make(chan BotMessage, 12)
	msgsOut = messages

	session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		messages <- BotMessage{
			ID: m.ID,
			Message: m.Message.Content,
			ChannelID: MessageDestination(m.ChannelID),
			Author: &User{
				ID: m.Author.ID,
				Name: m.Author.Username,
				session: s,
				ChannelID: MessageDestination(m.Author.ID)}}
	})
	return
}

func (msg BotMessage) SendMessage(message interface{}) error {
	return msg.ChannelID.SendMessage(message, msg.Author.session)
}

func (msg BotMessage) SendMessagef(format string, args ...interface{}) error {
	return msg.SendMessage(fmt.Sprintf(format, args...))
}

func (user *User) SendMessage(message interface{}) error {
	return user.ChannelID.SendMessage(message, user.session)
}

func (dest MessageDestination) SendMessage(message interface{}, session *discordgo.Session) error {
	var msgToSend string
	switch val := message.(type) {
	case string:
		msgToSend = val
	case int:
		msgToSend = strconv.Itoa(val)
	default:
		return errors.New("bad type passed")
	}

	_, err := session.ChannelMessageSend(string(dest), msgToSend)
	return err
}