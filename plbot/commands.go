package plbot

import (
	"github.com/bwmarrin/discordgo"
	"time"
	"strings"
	"strconv"
	"fmt"
)

/*
	Author: Charles Shook
	Description: A help command that displays all the commands that the bot can do.
*/
func HelpCommand(s * discordgo.Session, m *discordgo.MessageCreate) {
	helpMessage := "I am here to help! \n\n" +
		"My commands are: \n" +
		config.BotPrefix + " help -> Shows the help commands. \n" +
		config.BotPrefix + " date -> Shows the current date. \n" +
		config.BotPrefix + " add num1 num2 -> Adds two numbers. \n" +
		config.BotPrefix + " sub num1 num2 -> Subtracts two numbers. \n"

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

	intVar, err := strconv.ParseInt(command[1], 10, 64)

	if err != nil {
		return
	}

	for i := int64(0); i < intVar; i++ {
		chicken += "Chicken "
	}

	s.ChannelMessageSend(m.ChannelID, chicken)
}

func AddCommand(s * discordgo.Session, m *discordgo.MessageCreate) {
	content := strings.Replace(m.Content, config.BotPrefix, "", 1)
	commands := strings.Fields(content)
	solution := ""

	if len(commands) < 3 {
		return
	}

	solution = "Answer: " + commands[1] + " + " + commands[2] + " = "

	floatOne, err := strconv.ParseFloat(commands[1], 64)

	if err != nil {
		return
	}

	var floatTwo float64
	floatTwo, err = strconv.ParseFloat(commands[2], 64)

	solution += fmt.Sprintf("%f", floatOne + floatTwo)
	s.ChannelMessageSend(m.ChannelID, solution)
}

func SubCommand(s * discordgo.Session, m *discordgo.MessageCreate) {
	content := strings.Replace(m.Content, config.BotPrefix, "", 1)
	commands := strings.Fields(content)
	solution := ""

	if len(commands) < 3 {
		return
	}

	solution = "Answer: " + commands[1] + " - " + commands[2] + " = "

	floatOne, err := strconv.ParseFloat(commands[1], 64)

	if err != nil {
		return
	}

	var floatTwo float64
	floatTwo, err = strconv.ParseFloat(commands[2], 64)

	solution += fmt.Sprintf("%f", floatOne - floatTwo)
	s.ChannelMessageSend(m.ChannelID, solution)
}

func DivideCommand(s * discordgo.Session, m *discordgo.MessageCreate) {
	content := strings.Replace(m.Content, config.BotPrefix, "", 1)
	commands := strings.Fields(content)
	solution := ""

	if len(commands) < 3 {
		return
	}

	solution = "Answer: " + commands[1] + " / " + commands[2] + " = "

	floatOne, err := strconv.ParseFloat(commands[1], 64)

	if err != nil {
		return
	}

	var floatTwo float64
	floatTwo, err = strconv.ParseFloat(commands[2], 64)

	solution += fmt.Sprintf("%f", floatOne / floatTwo)
	s.ChannelMessageSend(m.ChannelID, solution)
}

func PlayCommand(s * discordgo.Session, m *discordgo.MessageCreate) {
	content := strings.Replace(m.Content, config.BotPrefix, "", 1)
	commands := strings.Fields(content)

	channel, _ := s.State.Channel(m.ChannelID)
	guilds, _ := s.State.Guild(channel.GuildID)

	for _, vs := range guilds.VoiceStates {
		if vs.UserID == m.Author.ID {
			go PlayURL(s, guilds.ID, vs.ChannelID, commands[1])
		}
	}
}

func StopCommand(s * discordgo.Session, m *discordgo.MessageCreate) {
	channel, _ := s.State.Channel(m.ChannelID)
	guilds, _ := s.State.Guild(channel.GuildID)

	for _, vs := range guilds.VoiceStates {
		if vs.UserID == m.Author.ID {
			Stop(s, guilds.ID, vs.ChannelID)
		}
	}
}

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

/*
	Author: Christopher Bare
	Description: A command to change the color of a role
*/
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

/*
	Author: Christopher Bare
	Description: Returns a list of the roles a member has
*/
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

/*
	Author: Christopher Bare
	Description: A function that tries to get a color from a hex code or color name
*/
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