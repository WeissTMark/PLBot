/*
	Author: Charles Shook
	Description: An interface for bot commands.
*/

package plbot

import (
	"github.com/bwmarrin/discordgo"
	"errors"
)

type command interface {
	create() (discordgo.MessageCreate, error)
	send(s *discordgo.Session, m *discordgo.MessageCreate) error
}