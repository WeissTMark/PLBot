/*
	Author: Mark Weiss
	Description: The guild structure to be used to contain data about the guild
*/
package plbot

type Guilds struct {
	id      string
	channel map[string]channels
}
