package main

import (
	"github.com/bwmarrin/discordgo"
)




func SendDeveloperMessage(s *discordgo.Session, message string) {
	ch, err := s.UserChannelCreate("217027600025911297")
	if err != nil {
		return
	}
	s.ChannelMessageSend(ch.ID,message)
}