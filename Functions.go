package main

import (
	"regexp"
	
	"github.com/bwmarrin/discordgo"
)



func isHello(message string) bool {
    for i := 0; i < len(hellowords); i++ {
        re := regexp.MustCompile(hellowords[i])
        matches := re.FindAllString(message, -1)
        if len(matches) > 0 {
            return true
        }
    }
    return false
}

func isBadWord(message string) bool {
    for i := 0; i < len(badwords); i++ {
        re := regexp.MustCompile(badwords[i])
        matches := re.FindAllString(message, -1)
        if len(matches) > 0 {
            return true
        }
    }
    return false	
}

func SendDeveloperMessage(s *discordgo.Session, message string) {
	ch, err := s.UserChannelCreate("217027600025911297")
	if err != nil {
		return
	}
	s.ChannelMessageSend(ch.ID,message)
}

func SendSuggestionsMessage(s *discordgo.Session, message string) {
	s.ChannelMessageSend("387135667710197762", message)
}
