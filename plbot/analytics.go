package plbot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func CreateAnalytics() map[string]Guilds {
	return make(map[string]Guilds)
}

func RunAnalytics(m *discordgo.MessageCreate, info map[string]Guilds) map[string]Guilds {
	_, there := info[m.GuildID]
	if !there {
		fmt.Println("Made it inside the create guild statement")
		c := map[string]channels{m.ChannelID: {}}
		g := Guilds{m.GuildID, c}
		info = map[string]Guilds{m.GuildID: g}
	}

	_, alsoThere := info[m.GuildID].channel[m.ChannelID]
	//fmt.Println("alsoThere: ", alsoThere)
	if !alsoThere {
		fmt.Println("Made it inside the create channel statement")
		var chans channels
		chans.id = m.ChannelID
		chans.analytics = data{}
		info[m.GuildID].channel[m.ChannelID] = chans
	}
	fmt.Println(info[m.GuildID].channel[m.ChannelID].analytics)

	info[m.GuildID].channel[m.ChannelID] = info[m.GuildID].channel[m.ChannelID].analytics.addWords(m.Message)
	//ana.addPunct(m.Message)
	//ana.addLetters(m.Message)
	//ana.addToD(m.Timestamp)
	return info
}
