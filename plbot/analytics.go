package plbot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func CreateAnalytics() map[string]guilds {
	return make(map[string]guilds)
}

func RunAnalytics(m *discordgo.MessageCreate, info map[string]guilds) {
	_, there := info[m.GuildID]
	if !there {
		info = map[string]guilds{m.GuildID: {}}
	}
	fmt.Print(info[m.GuildID])
	_, alsoThere := info[m.GuildID].channel[m.ChannelID]
	if !alsoThere {
		guild := info[m.GuildID]
		guild.channel = map[string]channels{m.GuildID: {}}
		info[m.GuildID] = guild
	}

	info[m.GuildID].channel[m.ChannelID].analytics.addWords(m.Message)
	//ana.addPunct(m.Message)
	//ana.addLetters(m.Message)
	//ana.addToD(m.Timestamp)

}
