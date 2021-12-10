/*
	Author: Mark Weiss
	Description: The two overarching functions to allow the constantly running bot to interact with the data
*/

package plbot

import (
	"github.com/bwmarrin/discordgo"
)

func CreateAnalytics() map[string]Guilds {
	return make(map[string]Guilds)
}

func RunAnalytics(m *discordgo.MessageCreate, info map[string]Guilds) map[string]Guilds {
	_, there := info[m.GuildID]
	if !there {
		c := map[string]channels{m.ChannelID: {}}
		g := Guilds{m.GuildID, c}
		info = map[string]Guilds{m.GuildID: g}
	}

	_, alsoThere := info[m.GuildID].channel[m.ChannelID]
	//fmt.Println("alsoThere: ", alsoThere)
	if !alsoThere {
		var chans channels
		chans.id = m.ChannelID
		chans.analytics = data{}
		info[m.GuildID].channel[m.ChannelID] = chans
	}

	var chans channels
	chans.id = m.ChannelID

	dat := info[m.GuildID].channel[m.ChannelID].analytics
	dat.addWords(m.Message)
	dat.addLetters(m.Message)
	dat.addpunct(m.Message)
	dat.addToD(m.Timestamp)

	chans.analytics = dat
	info[m.GuildID].channel[m.ChannelID] = chans

	return info
}
