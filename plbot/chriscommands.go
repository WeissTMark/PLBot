package plbot

import (
	"strconv"
	"strings"

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

	roleList := getMemberRoles(m.Member, g)

	// if no roles were found, say so
	if len(roleList) == 0 {
		s.ChannelMessageSend(m.ChannelID, "You don't appear to have any roles")
		return
	}

	output := ""

	for index, val := range roleList {
		output += strconv.Itoa(index) + ": " + val.Name + "\n"
	}

	// send message of roles
	s.ChannelMessageSend(m.ChannelID, output)
	return

}

func ChangeRoleColor(s *discordgo.Session, m *discordgo.MessageCreate) {
	// [1] role to change
	// [2] color

	content := strings.Replace(m.Content, config.BotPrefix, "", 1)
	command := strings.Fields(content)

	// check correct number of arguments
	if len(command) < 3 {
		s.ChannelMessageSend(m.ChannelID, "Usage: <prefix> rolecolor <role # to change> <color (hex or name)>\nUse <prefix> rolelist to get a list of your roles.")
		return
	}

	// get the role number (index)
	roleNum, err := strconv.Atoi(command[1])
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Please input the role number you would like to change, to get that, use: <prefix> rolelist.")
		return
	}

	// get the color
	color := getColor(command[2])
	if color == -1 {
		s.ChannelMessageSend(m.ChannelID, "I'm unaware of that color, you could also try inputting a hex color code instead!")
	}

	// get the guild id
	g, err := s.Guild(m.GuildID)
	if err != nil {
		return
	}

	// get member roles
	roleList := getMemberRoles(m.Member, g)

	// make sure index is in range
	if roleNum >= len(roleList) {
		s.ChannelMessageSend(m.ChannelID, "You do not have a role #"+strconv.Itoa(roleNum)+"\nUse <prefix> listroles to get a list of roles and the numbers associated")
		return
	}

	// get actual role
	role := roleList[roleNum]

	// edit role's color
	s.GuildRoleEdit(g.ID, role.ID, role.Name, int(color), role.Hoist, role.Permissions, role.Mentionable)
	return
}

func getMemberRoles(m *discordgo.Member, g *discordgo.Guild) []*discordgo.Role {
	memroles := m.Roles
	guildroles := g.Roles

	roleList := []*discordgo.Role{}

	for i := 0; i < len(memroles); i++ {
		// loop through all guild roles
		for j := 0; j < len(guildroles); j++ {
			// add role name to list as well as an index
			if memroles[i] == guildroles[j].ID {
				roleList = append(roleList, guildroles[j])
			}
		}
	}

	return roleList
}

func getColor(input string) int {
	// list of colors
	colors := map[string]int{"red": 0xFF0000, "green": 0x00FF00, "blue": 0x0000FF,
		"white": 0xFFFFFF, "black": 0x00001, "yellow": 0xFFFF00, "pink": 0xFF00FF,
		"cyan": 0x00FFFF}

	// make input lowercase for finding color in map
	input = strings.ToLower(input)

	// get the color from map
	color, found := colors[input]

	// if it wasnt in the map
	if !found {
		// try to convert it from hex
		hexcolor, err := strconv.ParseInt(input, 16, 64)
		// if you can't convert, return -1 as an error
		if err != nil {
			return -1
		}
		// if conversion was successful, return the color as an integer
		return int(hexcolor)
	}

	// return the color from the map
	return color
}
