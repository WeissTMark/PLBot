/*
	Author: Mark Weiss
	Description: Functions to handle stats requests
*/

package plbot

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func PrintStats(s *discordgo.Session, m *discordgo.MessageCreate, analytics map[string]Guilds) {
	msg := m.Content
	args := strings.Split(msg, " ")

	if len(args) >= 3 {
		chanel := args[2]
		getStats(s, m, chanel, analytics)
	} else {
		getStats(s, m, m.ChannelID, analytics)
	}

}

func getStats(s *discordgo.Session, m *discordgo.MessageCreate, chID string, analytics map[string]Guilds) {
	guil, er := analytics[m.GuildID]
	if !er {
		fmt.Println(guil)
		s.ChannelMessageSend(m.ChannelID, "Sorry, your server doesn't exist")
		return
	} else {
		cha, err := guil.channel[chID]
		if !err {
			s.ChannelMessageSend(m.ChannelID, "Sorry, that channel doesn't exist")
			return
		} else {
			s.ChannelMessageSend(m.ChannelID, cha.analytics.toString())
		}
	}
}
