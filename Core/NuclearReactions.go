package Core

import (
	"github.com/bwmarrin/discordgo"
)

type Reaction struct {
	Trigger  string
	Reaction string
	TMode    int
	/*1: contains trigger string
	2: is trigger string*/
}
type ReactionDB struct {
	servers map[string]map[int]Reaction
	//         serv. id  reac id
}

func (r *ReactionDB) Process(m discordgo.MessageCreate) {

}
