package plbot

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
)

/*
	Author: Christopher Bare
	Description: A command to list the roles a user has
*/
func ListRoles(s *discordgo.Session, m *discordgo.MessageCreate) {

	// get the guild id
	g, err := s.Guild(m.GuildID)
	if err != nil {
		return
	}

	// get the memeber's roles and the guild's roles
	memroles := m.Member.Roles
	guildroles := g.Roles

	roleList := ""

	// for each member role
	for i := 0; i < len(memroles); i++ {
		// loop through all guild roles
		for j := 0; j < len(guildroles); j++ {
			// add role name to list as well as an index
			if memroles[i] == guildroles[j].ID {
				roleList += strconv.Itoa(i) + ": "
				roleList += guildroles[j].Name + "\n"
			}
		}
	}

	// if no roles were found, say so
	if roleList == "" {
		s.ChannelMessageSend(m.ChannelID, "You don't appear to have any roles")
		return
	}

	// send message of roles
	s.ChannelMessageSend(m.ChannelID, roleList)
	return

}
