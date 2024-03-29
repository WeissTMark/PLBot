/*
	Author: Mark Weiss
	Description: All functions needed to add data to the data structure, as well as the "ToString" function that prints the data in a nice fashion
*/

package plbot

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/bwmarrin/discordgo"
)

type data struct {
	words   map[string]int
	letters map[rune]int
	punct   map[rune]int
	tod     map[int]int
}

/*
	Author: Mark Weiss
	Description: Adds all words in the message to the data structure
*/
func (d *data) addWords(m *discordgo.Message) {

	f := func(c rune) bool {
		ret := false
		if unicode.IsPunct(c) {
			ret = true
			if c == '\'' || c == '-' {
				ret = false
			}
		} else if unicode.IsSpace(c) {
			ret = true
		} else if unicode.IsNumber(c) {
			ret = true
		}
		return ret
	}

	cont := strings.FieldsFunc(m.Content, f)

	for _, word := range cont {
		if len(d.words) <= 0 {
			d.words = map[string]int{word: 1}
		} else {
			val, exists := d.words[word]
			if exists {
				d.words[word] = val + 1
			} else {
				d.words[word] = 1
			}
		}
	}
}

/*
	Author: Mark Weiss
	Description: Adds all letters in the given message to the data structure
*/
func (d *data) addLetters(m *discordgo.Message) {

	f := func(c rune) bool {
		ret := false
		if unicode.IsPunct(c) {
			ret = true
		} else if unicode.IsSpace(c) {
			ret = true
		}
		return ret
	}

	var ms []string
	for _, word := range strings.FieldsFunc(m.Content, f) {
		ms = append(ms, word)
	}
	msg := strings.Join(ms, "")

	for _, run := range msg {
		if len(d.letters) <= 0 {
			d.letters = map[rune]int{run: 1}
		} else {

			val, exists := d.letters[run]
			if exists {
				d.letters[run] = val + 1
			} else {
				d.letters[run] = 1
			}
		}
	}
}

/*
	Author: Mark Weiss
	Description: Adds all punctuation in the given message to the data structure
*/
func (d *data) addpunct(m *discordgo.Message) {

	f := func(c rune) bool {
		ret := true
		if unicode.IsPunct(c) {
			ret = false
		} else if unicode.IsSpace(c) {
			ret = false
		}
		return ret
	}

	var ms []string
	for _, word := range strings.FieldsFunc(m.Content, f) {
		ms = append(ms, word)
	}
	msg := strings.Join(ms, "")

	for _, run := range msg {
		if len(d.punct) <= 0 {
			d.punct = map[rune]int{run: 1}
		} else {

			val, exists := d.punct[run]
			if exists {
				d.punct[run] = val + 1
			} else {
				d.punct[run] = 1
			}
		}
	}
}

/*
	Author: Mark Weiss
	Description: Adds the hour of the given message to the data structure
*/
func (d *data) addToD(m discordgo.Timestamp) {
	hour := strings.Split(string(m), "T")
	hour = strings.Split(hour[1], ":")
	to, _ := strconv.Atoi(hour[0])
	to = (to - 7.0)
	if to < 0 {
		to = 24 + to
	}

	if len(d.tod) <= 0 {
		d.tod = map[int]int{to: 1}
	} else {

		val, exists := d.tod[to]
		if exists {
			d.tod[to] = val + 1
		} else {
			d.tod[to] = 1
		}
	}
}

/*
	Author: Mark Weiss
	Description: turns the given data structure into a nice looking discord formated string
*/
func (d data) toString() string {
	msg := "```Words: "
	for key, value := range d.words {
		msg += key + ":" + fmt.Sprint(value) + " "
	}
	msg += "\nLetters: "
	for key, value := range d.letters {
		msg += string(key) + ":" + fmt.Sprint(value) + " "
	}
	msg += "\nPunctuation: "
	for key, value := range d.punct {
		msg += string(key) + ":" + fmt.Sprint(value) + " "
	}
	msg += "\nTime of Day: "
	for key, value := range d.tod {
		msg += fmt.Sprint(key) + ":" + fmt.Sprint(value) + " "
	}
	msg += "\n```"
	return msg
}
