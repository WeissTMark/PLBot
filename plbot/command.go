package plbot

import (
	"github.com/bwmarrin/discordgo"
	"errors"
)

type command interface {
	create() error
	send(s *discordgo.Session, m *discordgo.MessageCreate) error
}