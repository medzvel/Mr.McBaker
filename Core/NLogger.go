package Core

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"time"
)

type UserInfo struct {
	LastSeen    time.Time
	LastMessage string
	LastChannel string
	LastGame    string
	PermLevel   int
	FancyPoints int
	Warns   	int
	Muted		int
}

type Logger struct {
	users map[string]*UserInfo
}

func MakeLogger() Logger {
	return Logger{make(map[string]*UserInfo)}
}
func (l *Logger) MakeUser(id string) {
	_, exists := l.users[id]
	if !exists {
		l.users[id] = &UserInfo{
			LastSeen:    time.Now(),
			LastMessage: "Last message not recorded",
			LastGame:    "Last played game not recorded",
			PermLevel:   0,
			Warns:		 0,
			Muted:		 0}
	}
}

func (l *Logger) GetInfo(id string) (*UserInfo, int) {
	user, exists := l.users[id]
	if exists {
		return user, 0
	} else {
		return &UserInfo{time.Now(), "invalid user", "chan", "invalid game", -1, -1, -1, -1}, 1
	}
}

func (l *Logger) SetPerm(id string, lv int) {
	if l.EntryExists(id) {
		l.users[id].PermLevel = lv
	}
}

func (l *Logger) SetPoints(id string, p int) {
	if l.EntryExists(id) {
		l.users[id].FancyPoints = p
	}
}

func (l *Logger) MuteUser(id string, p int) {
	if l.EntryExists(id) {
		l.users[id].Muted = p
	}
}

func (l *Logger) WarnUser(id string, p int) {
	if l.EntryExists(id) {
		l.users[id].Warns = p
	}
}

func (l *Logger) OutToFile(f string) {
	j, jerr := json.MarshalIndent(l.users, "", "\t")
	if jerr != nil {
		fmt.Println("error while generating json: \n", jerr)
		return
	}
	ioerr := ioutil.WriteFile(f, j, 0664)
	if ioerr != nil {
		fmt.Println("error while writing json to file: \n", ioerr)
		return
	}
	fmt.Println("Done writing to file ", f)
}

func (l *Logger) ReadFromFile(f string) {
	j, ioerr := ioutil.ReadFile(f)
	if ioerr != nil {
		fmt.Println("error while reading from file: \n", ioerr)
		return
	}
	jerr := json.Unmarshal(j, &l.users)
	if jerr != nil {
		fmt.Println("error while parsing json: \n", jerr)
		return
	}
	fmt.Println("Done reading from file ", f)
}

func (l *Logger) EntryExists(id string) bool {
	_, exists := l.users[id]
	return exists
}

func (l *Logger) UpdateEntryMsg(id string, m *discordgo.MessageCreate) {
	if l.EntryExists(id) {
		l.users[id].LastSeen = time.Now()
		l.users[id].LastMessage = m.Content
		l.users[id].LastChannel = m.ChannelID
	} else {
		l.MakeUser(id)
		l.UpdateEntryMsg(id, m)
	}
}

func (l *Logger) UpdateEntryPresence(id string, p *discordgo.PresenceUpdate) {
	if l.EntryExists(id) {
		l.users[id].LastSeen = time.Now()
		/// FIXME
		//l.users[id].LastGame = p.Presence.Game.Name //causes crash
	} else {
		l.MakeUser(id)
		l.UpdateEntryPresence(id, p)
	}
}

func (l *Logger) DeleteEntry(id string) {
	if l.EntryExists(id) {
		delete(l.users, id)
	}
}
