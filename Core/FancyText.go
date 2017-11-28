package Core

import (
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

/*
	Placeholders: (idea definitely not stolen from nadeko, kek)
	%user% - mention to sending user
	%guildName% - server's name
	%chanName% - channel's name
	%randUser% - pick a random user
*/
//type FTCore struct {
//	session *discordgo.Session
//	logger  *Logger
//}

func FancifyText(t string, s *discordgo.Session, m *discordgo.MessageCreate) string {
	rand.Seed(time.Now().Unix())
	retStr := t
	guildname := "invalid guild"
	channame := "invalid channel"
	channel, cerr := s.Channel(m.ChannelID)
	var guild *discordgo.Guild
	var gerr error
	if cerr == nil {
		channame = channel.Name
		guild, gerr = s.Guild(channel.GuildID)
		if gerr == nil {
			guildname = guild.Name
		}
	}
	retStr = strings.Replace(retStr, "$user", m.Author.Mention(), -1)
	retStr = strings.Replace(retStr, "$guildName", guildname, -1)
	retStr = strings.Replace(retStr, "$chanName", "#"+channame, -1)
	for has := strings.Contains(retStr, "$randUser"); has; has = strings.Contains(retStr, "$randUser") {
		user := guild.Members[rand.Intn(guild.MemberCount-1)].User //will loop forever if a bot executes this in a 1 person guild, if possible
		if !user.Bot {
			retStr = strings.Replace(retStr, "$randUser", user.Mention(), 1)
		}
	}
	return retStr
}
