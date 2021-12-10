package plbot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"github.com/bwmarrin/discordgo"
)

var dg *discordgo.Session

func ConnectDiscord() (error) {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		return err
	}
	
	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(HandlerMessageCreate)
	
	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		return err
	}
	
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	
	// Cleanly close down the Discord session.
	dg.Close()

	return nil
}