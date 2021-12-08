package main

import (
	"plbot"
)


func main() {
	err := plbot.LoadConfig("config.json")

	if err != nil {
		return
	}

	err = plbot.ConnectDiscord()

}