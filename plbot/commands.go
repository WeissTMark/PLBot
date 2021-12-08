package plbot

import (
	"github.com/bwmarrin/discordgo"
	"time"
	"strings"
	"strconv"
)

/*
	Author: Charles Shook
	Description: A help command that displays all the commands that the bot can do.
*/
func HelpCommand(s * discordgo.Session, m *discordgo.MessageCreate) {
	helpMessage := "I am here to help! \n\n" +
		"My commands are: \n" +
		config.BotPrefix + " help -> Shows the help commands. \n" +
		config.BotPrefix + " date -> Shows the current date. \n"

	s.ChannelMessageSend(m.ChannelID, helpMessage)
}


/*
	Author: Charles Shook
	Description: Command that allows the bot to return the current date.
*/
func DateCommand(s * discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, time.Now().Format("01-02-2006"))
}

/*
	Author: Charles Shook
	Description:Command that impliments the Chicken language talked about in PL.
*/
func ChickenCommand(s * discordgo.Session, m *discordgo.MessageCreate) {
	content := strings.Replace(m.Content, config.BotPrefix, "", 1)
	command := strings.Fields(content)
	chicken := ""

	if len(command) < 2 {
		return
	}

	intVar, err := strconv.Atoi(command[1])

	if err != nil {
		return
	}

	for i := 0; i <intVar; i++ {
		chicken += "Chicken "
	}

	s.ChannelMessageSend(m.ChannelID, chicken)
}

// func RadioCommand(s * discordgo.Session, m *discordgo.MessageCreate, v *Voice) {
// 	if v == nil {
// 		return
// 	}

// 	radio := Radio{"", v}
// 	radio.data = strings.Fields(m.Content)[1]

// 	go func() {
// 		radioSignal <- radio
// 	}()
// }

// func JoinCommand(s * discordgo.Session, m *discordgo.MessageCreate, v *Voice) {
// 	voiceChannelID := ""

// 	for _, g := range dg.State.Guilds {
// 		for _, v := range g.VoiceStates {
// 			if v.UserID == m.Author.Username {
// 				voiceChannelID = v.ChannelID
// 			}
// 		}
// 	}

// 	if voiceChannelID == "" {
// 		return
// 	}

// 	if v != nil {
// 		return
// 	} else {
// 		channel, _ := dg.Channel(m.ChannelID)
// 		guildId := channel.GuildID

// 		v = new(Voice)
// 		voiceInstances[guildId] = v 
// 		v.session = s
// 	}

// }