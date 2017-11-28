package main

import (
	Core "MrMcBaker/Core"
	"fmt"
	"os"
	"strconv"
	"github.com/bwmarrin/discordgo"
)

//For commands

func registerCommands(p *Core.Parser) {
	echoCmd := Core.Command{
		Name:              "echo",
		ArgumentCount:     1,
		HelpMsg:           "A command to *echo* what you say",
		UsageMsg:          "echo [message]",
		IsDisplayedOnHelp: true,
		PermLevel:         3,
		Category:          "Administrative",
		FancifyInput:      true,
		Command:           echo}
	p.Register(&echoCmd)

	pingCmd := Core.Command{
		Name:              "Ping!",
		ArgumentCount:     0,
		HelpMsg:           "Ping the bot! Or maybe a website in future...",
		UsageMsg:          "Ping!",
		IsDisplayedOnHelp: true,
		PermLevel:         0,
		Category:          "General",
		FancifyInput:      true,
		Command:           ping}
	p.Register(&pingCmd)

	succCmd := Core.Command{
		Name:              "succ",
		ArgumentCount:     0,
		HelpMsg:           "succ someone",
		UsageMsg:          "succ (opt.)[user mention]",
		IsDisplayedOnHelp: true,
		PermLevel:         0,
		Category:          "Memes",
		FancifyInput:      true,
		Command:           succ}
	p.Register(&succCmd)

	fuccCmd := Core.Command{
		Name:              "fucc",
		ArgumentCount:     0,
		HelpMsg:           "fucc someone",
		UsageMsg:          "fucc (opt.)[user mention]",
		IsDisplayedOnHelp: true,
		PermLevel:         0,
		Category:          "Memes",
		FancifyInput:      true,
		Command:           fucc}
	p.Register(&fuccCmd)

	whoamiCmd := Core.Command{
		Name:              "whoami",
		ArgumentCount:     0,
		HelpMsg:           "A command to get your user info",
		UsageMsg:          "whoami",
		IsDisplayedOnHelp: true,
		PermLevel:         0,
		Category:          "Information",
		FancifyInput:      true,
		Command:           whoami}
	p.Register(&whoamiCmd)

	seenCmd := Core.Command{
		Name:              "whois",
		ArgumentCount:     1,
		HelpMsg:           "A command to see the info of a user mentioned",
		UsageMsg:          "whois [user mention]",
		IsDisplayedOnHelp: true,
		PermLevel:         0,
		Category:          "Information",
		FancifyInput:      true,
		Command:           seen}
	p.Register(&seenCmd)

	shutdownCmd := Core.Command{
		Name:              "shutdown",
		ArgumentCount:     0,
		HelpMsg:           "Shuts down the bot",
		UsageMsg:          "shutdown",
		IsDisplayedOnHelp: true,
		PermLevel:         3,
		Category:          "Administrative",
		FancifyInput:      true,
		Command:           shutdown}
	p.Register(&shutdownCmd)

	setPointsCmd := Core.Command{
		Name:              "setPoints",
		ArgumentCount:     2,
		HelpMsg:           "Sets a user's points",
		UsageMsg:          "setPoints <mention> <points>",
		IsDisplayedOnHelp: true,
		PermLevel:         3,
		Category:          "Administrative",
		FancifyInput:      true,
		Command:           setPoints}
	p.Register(&setPointsCmd)

	warnCmd := Core.Command{
		Name:              "warnUser",
		ArgumentCount:     2,
		HelpMsg:           "Warns User",
		UsageMsg:          "warn <meintion> <warn points>",
		IsDisplayedOnHelp: true,
		PermLevel:         1,
		Category:          "Administrative",
		FancifyInput:      true,
		Command:           warnUser}
	p.Register(&warnCmd)

	setPermCmd := Core.Command{
		Name:              "setPerm",
		ArgumentCount:     2,
		HelpMsg:           "Sets a user's permission level",
		UsageMsg:          "setPerm <mention> <permlevel>",
		IsDisplayedOnHelp: true,
		PermLevel:         3,
		Category:          "Administrative",
		FancifyInput:      true,
		Command:           setPerm}
	p.Register(&setPermCmd)

	setPrefixCmd := Core.Command{
		Name:              "setPrefix",
		ArgumentCount:     1,
		HelpMsg:           "Set the bot's command prefix",
		UsageMsg:          "setPrefix <prefix>",
		IsDisplayedOnHelp: true,
		PermLevel:         3,
		Category:          "Administrative",
		FancifyInput:      true,
		Command:           setprefix}
	p.Register(&setPrefixCmd)

	pointsCmd := Core.Command{
		Name:              "points",
		ArgumentCount:     3,
		HelpMsg:           "Main points command",
		UsageMsg:          "setPrefix <give/take/set/gift> <mention> <value>",
		IsDisplayedOnHelp: true,
		PermLevel:         0,
		Category:          "Miscellaneous",
		FancifyInput:      true,
		Command:           points}
	p.Register(&pointsCmd)

	creditsCmd := Core.Command{
		Name:              "credits",
		ArgumentCount:     0,
		HelpMsg:           "Shows credits",
		UsageMsg:          "credits",
		IsDisplayedOnHelp: true,
		PermLevel:         0,
		Category:          "Miscellaneous",
		FancifyInput:      true,
		Command:           credits}		
	p.Register(&creditsCmd)
}

func echo(args Core.Arguments, s *discordgo.Session, m *discordgo.MessageCreate) string {
	var retString string
	for i := 1; len(args.Args) > i; i++ {
		retString = fmt.Sprintln(retString, args.Args[i])
	}
	return retString
}


func ping(args Core.Arguments, s *discordgo.Session, m *discordgo.MessageCreate) string {
	return "Pong!"
}

func succ(args Core.Arguments, s *discordgo.Session, m *discordgo.MessageCreate) string {
	//s.ChannelMessageDelete(m.ChannelID, m.Message.ID)
	if args.Count >= 1 {
		return fmt.Sprintf("***%s succs %s***", m.Author.Mention(), args.Args[1])
	} else {
		return fmt.Sprintf("***%s succs %s***", s.State.User.Mention(), m.Author.Mention())
	}
}

func fucc(args Core.Arguments, s *discordgo.Session, m *discordgo.MessageCreate) string {
	//s.ChannelMessageDelete(m.ChannelID, m.Message.ID)
	if args.Count >= 1 {
		return fmt.Sprintf("***%s fuccs %s***", m.Author.Mention(), args.Args[1])
	} else {
		return fmt.Sprintf("***%s fuccs %s***", s.State.User.Mention(), m.Author.Mention())
	}
}

func whoami(args Core.Arguments, s *discordgo.Session, m *discordgo.MessageCreate) string {
	s.ChannelMessageDelete(m.ChannelID, m.Message.ID)
	ch, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Can't create PM Channel for user. ERROR: %s", err))
		return ""
	}
	author := discordgo.MessageEmbedAuthor{
		Name:    fmt.Sprintln("User info of: ", m.Author.String()),
		IconURL: m.Author.AvatarURL("")}

	retEmbed := discordgo.MessageEmbed{
		Author: &author,
		Color:  0x43c605}

	user, _ := Logger.GetInfo(m.Author.ID)
	retEmbed.Fields = append(retEmbed.Fields, &discordgo.MessageEmbedField{Name: "Is bot?", Value: fmt.Sprintln(m.Author.Bot), Inline: false})
	retEmbed.Fields = append(retEmbed.Fields, &discordgo.MessageEmbedField{Name: "Username", Value: m.Author.String(), Inline: false})
	retEmbed.Fields = append(retEmbed.Fields, &discordgo.MessageEmbedField{Name: "Mention", Value: m.Author.Mention(), Inline: false})
	retEmbed.Fields = append(retEmbed.Fields, &discordgo.MessageEmbedField{Name: "ID", Value: m.Author.ID, Inline: false})
	retEmbed.Fields = append(retEmbed.Fields, &discordgo.MessageEmbedField{Name: "Permission level", Value: fmt.Sprintf("%v", user.PermLevel), Inline: false})
	retEmbed.Fields = append(retEmbed.Fields, &discordgo.MessageEmbedField{Name: "Fancy points", Value: fmt.Sprintf("%v", user.FancyPoints), Inline: false})
	retEmbed.Fields = append(retEmbed.Fields, &discordgo.MessageEmbedField{Name: "Warn points", Value: fmt.Sprintf("%v", user.Warns), Inline: false})
	s.ChannelMessageSendEmbed(ch.ID, &retEmbed)
	return ""//fmt.Sprintln("Here you go ", m.Author.Mention(), "!")
}

func credits(args Core.Arguments, s *discordgo.Session, m *discordgo.MessageCreate) string {
	s.ChannelMessageDelete(m.ChannelID, m.Message.ID)
	ch, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Can't create PM Channel for user. ERROR: %s", err))
		return ""
	}
	embed := &discordgo.MessageEmbed{
	    Author:      &discordgo.MessageEmbedAuthor{},
	    Color:       0x00ff00, // Green
	    Description: "Mr.McBaker - Credits",
	    Fields: []*discordgo.MessageEmbedField{
	        &discordgo.MessageEmbedField{
	            Name:   "Author",
	            Value:  "Amagida",
	            Inline: true,
	        },
	        &discordgo.MessageEmbedField{
	            Name:   "My Owner Project",
	            Value:  Config.ProjectName,
	            Inline: true,
	        },
	    },
	    Image: &discordgo.MessageEmbedImage{
	        URL: "https://cdn.discordapp.com/app-icons/369935120137977877/22f44170060cea2413e6809c11f29e61.png",
	    },
	    Thumbnail: &discordgo.MessageEmbedThumbnail{
	        URL: "https://cdn.discordapp.com/app-icons/369935120137977877/22f44170060cea2413e6809c11f29e61.png",
	    },
	    Title:     "Hello, I'm Mr.McBaker!",
	}
	s.ChannelMessageSendEmbed(ch.ID, embed)
	return ""//fmt.Sprintln("Here you go ", m.Author.Mention(), "!")
}

func seen(args Core.Arguments, s *discordgo.Session, m *discordgo.MessageCreate) string {
	s.ChannelMessageDelete(m.ChannelID, m.Message.ID)
	ch, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Can't create PM Channel for user. ERROR: %s", err))
		return ""
	}
	if len(m.Mentions) > 0 {
		if Logger.EntryExists(m.Mentions[0].ID) {
			user, _ := Logger.GetInfo(m.Mentions[0].ID)
			author := discordgo.MessageEmbedAuthor{
				Name:    fmt.Sprintln("Last seen info of ", m.Mentions[0].String()),
				IconURL: m.Mentions[0].AvatarURL("")}
			retEmbed := discordgo.MessageEmbed{
				Author: &author,
				Color:  0x05c699}
			retEmbed.Fields = append(retEmbed.Fields, &discordgo.MessageEmbedField{Name: "Last seen on", Value: fmt.Sprintln(user.LastSeen)})

			LastChanName := "Invalid Channel"
			LastGuildName := "Invalid Guild"

			LastChan, cerr := s.Channel(user.LastChannel)

			if cerr != nil {
				//Logger.DeleteEntry(m.Mentions[0].ID)
				fmt.Println("Invalid channel detected on ", m.Mentions[0].ID)
			} else {
				LastGuild, _ := s.Guild(LastChan.GuildID)
				LastChanName = LastChan.Name
				LastGuildName = LastGuild.Name
			}
			retEmbed.Fields = append(retEmbed.Fields, &discordgo.MessageEmbedField{Name: "ID", Value: m.Mentions[0].ID, Inline: false})
			retEmbed.Fields = append(retEmbed.Fields, &discordgo.MessageEmbedField{Name: "Last messaged server & channel", Value: fmt.Sprintln(LastGuildName, ", ", LastChanName)})
			retEmbed.Fields = append(retEmbed.Fields, &discordgo.MessageEmbedField{Name: "Last message", Value: user.LastMessage})
			retEmbed.Fields = append(retEmbed.Fields, &discordgo.MessageEmbedField{Name: "Last played game", Value: user.LastGame})
			retEmbed.Fields = append(retEmbed.Fields, &discordgo.MessageEmbedField{Name: "Permission level", Value: fmt.Sprintf("%v", user.PermLevel), Inline: false})
			retEmbed.Fields = append(retEmbed.Fields, &discordgo.MessageEmbedField{Name: "Fancy points", Value: fmt.Sprintf("%v", user.FancyPoints), Inline: false})
			retEmbed.Fields = append(retEmbed.Fields, &discordgo.MessageEmbedField{Name: "Warn points", Value: fmt.Sprintf("%v", user.Warns), Inline: false})
			_, err := s.ChannelMessageSendEmbed(ch.ID, &retEmbed)
			if err != nil {
				fmt.Println(err)
			}
			return ""
		} else {
			s.ChannelMessageSend(ch.ID, "User is not registered!")
			return ""
		}
	} else {
		s.ChannelMessageSend(ch.ID, "Invalid mention!")
		return ""
	}
	return "error?"
}

func shutdown(args Core.Arguments, s *discordgo.Session, m *discordgo.MessageCreate) string {
	bot.Close()
	Config.End(CfgFile, &Parser, &Logger)
	os.Exit(0)
	return "this message shouldnt be seen"
}

func setPerm(args Core.Arguments, s *discordgo.Session, m *discordgo.MessageCreate) string {
	if len(m.Mentions) > 0 {
		if Logger.EntryExists(m.Mentions[0].ID) {
			val, err := strconv.Atoi(args.Args[2])
			if err == nil {
				Logger.SetPerm(m.Mentions[0].ID, val)
				return "Done!"
			} else {
				return "Not a number!"
			}
		} else {
			return "User is not registered"
		}
	} else {
		return "Invalid mention!"
	}
	return ""
}

func warnUser(args Core.Arguments, s *discordgo.Session, m *discordgo.MessageCreate) string {
	if len(m.Mentions) > 0 {
		if Logger.EntryExists(m.Mentions[0].ID) {
			val, err := strconv.Atoi(args.Args[2])
			if err == nil {
				Logger.WarnUser(m.Mentions[0].ID, val)
				var warnedmsg string
				warnedmsg = fmt.Sprintf("Administrator %s warned user %s.", m.Author.Mention(), m.Mentions[0].String())
				return warnedmsg
			} else {
				return "Not a number!"
			}
		} else {
			return "User is not registered"
		}
	} else {
		return "Invalid mention!"
	}
	return ""
}

func setPoints(args Core.Arguments, s *discordgo.Session, m *discordgo.MessageCreate) string {
	if len(m.Mentions) > 0 {
		if Logger.EntryExists(m.Mentions[0].ID) {
			val, err := strconv.Atoi(args.Args[2])
			if err == nil {
				Logger.SetPoints(m.Mentions[0].ID, val)
				return "Done!"
			} else {
				return "Not a number!"
			}
		} else {
			return "User is not registered"
		}
	} else {
		return "Invalid mention!"
	}
	return ""
}

func setprefix(args Core.Arguments, s *discordgo.Session, m *discordgo.MessageCreate) string {
	Parser.SetPrefix(args.Args[1])
	return "Done!"
}

func points(args Core.Arguments, s *discordgo.Session, m *discordgo.MessageCreate) string {
	//give, take, set
	val, err := strconv.Atoi(args.Args[3])
	if err != nil {
		return "3rd argument must be a number!"
	}
	if len(m.Mentions) > 0 {
		if !Logger.EntryExists(m.Mentions[0].ID) {
			return "The user mentioned isnt registered!"
		}
		usrT, _ := Logger.GetInfo(m.Mentions[0].ID) //target user
		usrS, _ := Logger.GetInfo(m.Author.ID)      //source user
		if args.Args[1] == "give" {
			if usrS.PermLevel >= 3 {
				if (usrT.FancyPoints + val) < 0 {
					Logger.SetPoints(m.Mentions[0].ID, 0)
				} else {
					Logger.SetPoints(m.Mentions[0].ID, usrT.FancyPoints+val)
				}
			} else {
				return "Insufficent permissions."
			}
		} else if args.Args[1] == "take" {
			if usrS.PermLevel >= 3 {
				if (usrT.FancyPoints - val) < 0 {
					Logger.SetPoints(m.Mentions[0].ID, 0)
				} else {
					Logger.SetPoints(m.Mentions[0].ID, usrT.FancyPoints-val)
				}
			} else {
				return "Insufficent permissions."
			}
		} else if args.Args[1] == "set" {
			if usrS.PermLevel >= 3 {
				Logger.SetPoints(m.Mentions[0].ID, val)
			} else {
				return "Insufficent permissions."
			}
		} else if args.Args[1] == "gift" {
			return "Command not yet ready"
		} else {
			return "Wrong subcommand, it should be give/take/set/gift!"
		}
	} else {
		return "Invalid mention!"
	}
	return "Done!"
}