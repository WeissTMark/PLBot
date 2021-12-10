/*
	Author: Mark Weiss
	Description: The Guilds struct to allow multiple guilds
*/
package plbot

type Guilds struct {
	id      string
	channel map[string]channels
}
