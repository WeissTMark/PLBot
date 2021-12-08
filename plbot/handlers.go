/*
	Author: Charles Shook
	Description: Handlers for different type of messages.
*/

package plbot

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func HandlerMessageCreate(s * discordgo.Session, m *discordgo.MessageCreate) {
	// Do NOT respond to messages that the bot sends
	if m.Author.ID == s.State.User.ID {
        return
    }

	if !strings.HasPrefix(m.Content, config.BotPrefix) {
		return
	}

	// channel, _ := dg.Channel(m.ChannelID)
	// v := voiceInstances[channel.GuildID]

	content := strings.Replace(m.Content, config.BotPrefix, "", 1)
	command := strings.Fields(content)

	if len(command) == 0 {
		return
	}

	switch command[0] {
	case "help":
		HelpCommand(s, m)
	case "date":
		DateCommand(s, m)
	case "chicken":
		ChickenCommand(s, m)
	// case "radio":
	// 	RadioCommand(s, m, v)
	// case "join":
	// 	JoinCommand(s, m, v)
	default:
		return
	}
 }
