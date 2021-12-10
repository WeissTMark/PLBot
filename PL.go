package main

import (
	"fmt"
	"math"
	"os"
	"os/signal"
	"plbot"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var TOKEN = "OTAwNzk5NDg3MjcyOTQ3NzEy.YXGk5Q.Rali3Qb2DkaH5xOlJ6zZRv6kVsY"
var ANALYTICS map[string]plbot.Guilds

func main() {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + TOKEN)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()

}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	//Check if guild/channel exist, and/or create them
	ANALYTICS = plbot.RunAnalytics(m, ANALYTICS)

	if !strings.HasPrefix(m.Content, "!") {
		return
	}

	command := strings.Split(m.Content, " ")
	command[0] = strings.ToLower(command[0])
	command[0] = strings.Replace(command[0], "!", "", 1)

	switch command[0] {
	case "stats":

		plbot.PrintStats(s, m, ANALYTICS)
	}
	fmt.Println(quadraticEquation(1.0, 2.0, 3.0))
}

func quadraticEquation(a float64, b float64, c float64) float64 {
	var x float64
	x = -b + math.Sqrt((b*b)-4*a*c)
	x = x / (2 * a)
	return x
}
