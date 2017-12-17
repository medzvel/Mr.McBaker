package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"github.com/medzvel/Mr.McBaker/Core"
)

var (
	Token   string
	CfgFile string
	LangFile string
	Config  Core.Config
	Parser  Core.Parser
	Logger  Core.Logger
	//Functions Core.Functions
	bot     *discordgo.Session
	err     error
)

var badwords = []string{
	`\bfuck\b`, 
	`\bFUCK\b`, 
	`\bbitch\b`, 
	`\bBITCH\b`, 
	`\bFuck\b`, 
	`\bBitch\b`,
}
var hellowords = []string{
    `\bhi\b`,
    `\bHI\b`,
    `\bhello\b`,
    `\bHELLO\b`,
    `\bHi\b`,
    `\bHello\b`,
}

var helloreactemojis []string = []string{"🖐", "🇭", "🇮"}

//var botstatus int = 0

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.StringVar(&CfgFile, "c", "", "Config file")
	flag.Parse()

	Parser, Logger = Config.Init("config")
	registerCommands(&Parser)
	Parser.LinkLogger(&Logger)
}

func main() {

	bot, err = discordgo.New("Bot " + "MzkyMDk0NTczNjUzOTE3NzA2.DRiOJw.UULvYvaFgxlF5EN03DpTnm9Nhrs")
	if err != nil {
		fmt.Println("Error creating Discord session:\n\t", err)
		return
	}

	bot.AddHandler(onMessage)
	bot.AddHandler(onStatusUpdate)
	bot.AddHandler(onJoin)
	err = bot.Open()
	if err != nil {
		fmt.Println("Error opening connection:\n\t", err)
		return
	}

	fmt.Println("Bot is up and running!")
	bot.UpdateStatus(0, Config.Playing)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	bot.Close()
	Config.End(CfgFile, &Parser, &Logger)
}

func onJoin(s *discordgo.Session, event *discordgo.GuildMemberAdd) {
	ch, err := s.UserChannelCreate(event.Member.User.ID)
	if err != nil {
		return
	}
	s.ChannelMessageSend(ch.ID, "Welcome to the Discord-Channel of the Tropical Island Roleplay SA-MP community.\nBy joining our Discord-Channel we would friendly request you to change your channel-username to your in-game name.Also we would appreciate it if you register a forum-account on www.tikibaye.com to get a better touch with our community.And read the community rules in the `annoncements` sub-section here in our Discord-Channel or on our forums.Enjoy your stay in the community!")
}

func onMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	user, _ := Logger.GetInfo(m.Author.ID)
	ch, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Can't create PM Channel for user. ERROR: %s", err))
		return
	}
	if m.Author.ID == s.State.User.ID {
		Logger.UpdateEntryMsg(m.Author.ID, m)
		return
	}

	s.ChannelMessageSend(m.ChannelID, Parser.Execute(s, m))
	if strings.Contains(m.Content, "🅱") {
		s.MessageReactionAdd(m.ChannelID, m.ID, "🅱")
	}
	/* BAD WORD DETECTOR */
	if isBadWord(m.Message.Content) {
		s.ChannelMessageDelete(m.ChannelID, m.Message.ID)
		s.ChannelMessageSend(ch.ID, "You can't use bad words!")
	}

	if isHello(m.Message.Content) {
		AddHelloReaction(s, m)
	}
	//
	/* HELLO WORD DETECTOR */ 

	if user.Muted == 1 {
		s.ChannelMessageDelete(m.ChannelID, m.Message.ID)
	}
	if len(m.Content) < 2 {
		s.ChannelMessageDelete(m.ChannelID, m.Message.ID)
	}
	Logger.UpdateEntryMsg(m.Author.ID, m)
	//SendDeveloperMessage(s, "HELLO DEVELOPER!")
	//s.ChannelMessageSend(ch.ID, fmt.Sprintf("YOUR ID IS %s", m.Author.ID))
}

func onStatusUpdate(s *discordgo.Session, p *discordgo.PresenceUpdate) {
	Logger.UpdateEntryPresence(p.Presence.User.ID, p)
}

func AddHelloReaction(s *discordgo.Session, m *discordgo.MessageCreate) {

	for i := 0; i < len(helloreactemojis); i++ {
		err := s.MessageReactionAdd(m.ChannelID, m.ID, helloreactemojis[i])
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("ERROR HAPPENED WHILE ADDING REACTION TO MESSAGE. ERROR: `%s`! `EMOJI ID: %s`", err, helloreactemojis[i]))
		}
	}

}