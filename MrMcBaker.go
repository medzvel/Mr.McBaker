package main

import (
	Core "MrMcBaker/Core"
	"flag"
	//"time"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	Token   string
	CfgFile string
	Config  Core.Config
	Parser  Core.Parser
	Logger  Core.Logger
	bot     *discordgo.Session
	err     error
)

var badwords []string = []string{"fuck", "FUCK", "bitch", "BITCH"}
var hellowords []string = []string{"hi", "HI", "hello", "HELLO"}
var helloreactemojis []string = []string{"383729614066810880", "h", "e", "l", "l", "o"}

//var botstatus int = 0

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.StringVar(&CfgFile, "c", "", "Config file")
	flag.Parse()

	Parser, Logger = Config.Init(CfgFile)
	registerCommands(&Parser)
	Parser.LinkLogger(&Logger)
}

func main() {

	bot, err = discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error creating Discord session:\n\t", err)
		return
	}

	bot.AddHandler(onMessage)
	bot.AddHandler(onStatusUpdate)

	err = bot.Open()
	if err != nil {
		fmt.Println("Error opening connection:\n\t", err)
		return
	}

	fmt.Println("Bot is up and running!")
	bot.UpdateStatus(0, Config.Playing)
	//BOT STATUS UPDATER
   /* statusupdater := time.NewTimer(time.Second * 5)
    go func() {
        <-statusupdater.C
        if botstatus == 0 {
        	bot.UpdateStatus(0, fmt.Sprintf("%shelp", Config.Prefix))
        	botstatus = 1
        	statusupdater.Reset(time.Second * 5)
        } else {
        	bot.UpdateStatus(0, Config.Playing)
        	botstatus = 0
        	statusupdater.Reset(time.Second * 5)
        }

    }()*/

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	bot.Close()
	Config.End(CfgFile, &Parser, &Logger)
}

func onMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	//discordgo.Channel().GuildID
	user, _ := Logger.GetInfo(m.Author.ID)
	if m.Author.ID == s.State.User.ID {
		Logger.UpdateEntryMsg(m.Author.ID, m)
		return
	}
	s.ChannelMessageSend(m.ChannelID, Parser.Execute(s, m))
	if strings.Contains(m.Content, "🅱") {
		s.MessageReactionAdd(m.ChannelID, m.ID, "🅱")
	}
	for i := 0; i < len(badwords); i++ {
		if strings.Contains(m.Content, badwords[i]) {
			s.ChannelMessageDelete(m.ChannelID, m.Message.ID)
			Logger.MuteUser(m.Author.ID, Config.MuteTime)
			s.ChannelMessageSend("369935120137977877", "TEST PM MESSAGE YLEO RATO IGINEBI HA?")
		}
	}
	for i := 0; i < len(hellowords); i++ {
		if strings.Contains(m.Content, hellowords[i]) {
			AddHelloReaction(s, m)
		}
	}
	if user.Muted == 1 {
		s.ChannelMessageDelete(m.ChannelID, m.Message.ID)
	}
	if len(m.Content) < 2 {
		s.ChannelMessageDelete(m.ChannelID, m.Message.ID)
	}
	Logger.UpdateEntryMsg(m.Author.ID, m)
}

func onStatusUpdate(s *discordgo.Session, p *discordgo.PresenceUpdate) {
	Logger.UpdateEntryPresence(p.Presence.User.ID, p)
}

func AddHelloReaction(s *discordgo.Session, m *discordgo.MessageCreate) {
	
	for i := 0; i < len(helloreactemojis); i++ {
		s.MessageReactionAdd(m.ChannelID, m.ID, helloreactemojis[i])
	}

}
