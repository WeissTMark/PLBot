module PL

go 1.17

require github.com/bwmarrin/discordgo v0.23.2

replace plbot => ./plbot

require (
	github.com/gorilla/websocket v1.4.0 // indirect
	github.com/jonas747/dca v0.0.0-20210930103944-155f5e5f0cc7 // indirect
	github.com/jonas747/ogg v0.0.0-20161220051205-b4f6f4cf3757 // indirect
	golang.org/x/crypto v0.0.0-20181030102418-4d3f4d9ffa16 // indirect
	plbot v0.0.0-00010101000000-000000000000 // indirect
)
