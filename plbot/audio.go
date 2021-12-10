package plbot

import (
	"github.com/bwmarrin/discordgo"
	"time"
	"github.com/jonas747/dca"
)

var voiceChannel *discordgo.VoiceConnection
var isPlaying bool

/*
	Author: Charles Shook
	Description: Plays a audio file that is on the same device that the bot is running on.
*/
func PlayURL(s * discordgo.Session, serverID string, channelID string, url string) (error) {
	options := dca.StdEncodeOptions
	options.RawOutput = true
	options.Bitrate = 128
	options.Application = "lowdelay"

	voiceChannel, err := s.ChannelVoiceJoin(serverID, channelID, false, false)

	if err != nil {
		return err
	}

	time.Sleep(500 * time.Millisecond)

	voiceChannel.Speaking(true)

	encodeSession, err := dca.EncodeFile(url, options)

	if err != nil {
		return err
	}

	done := make(chan error)

	dca.NewStream(encodeSession, voiceChannel, done)

	err = <-done

	return nil
}

/*
	Author: Charles Shook
	Description: Stops the playing of audio.
*/
func Stop(s * discordgo.Session, serverID string, channelID string) (error) {
	vc, err := s.ChannelVoiceJoin(serverID, channelID, false, false)

	if err != nil {
		return err
	}

	vc.Speaking(false)

	time.Sleep(500 * time.Millisecond)

	vc.Disconnect()

	return nil
}