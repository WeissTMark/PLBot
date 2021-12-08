package plbot

import (
	"strings"
	"unicode"

	"github.com/bwmarrin/discordgo"
)

type channel interface {
	addWords()
	addLetters()
	addPunct()
	addToD()
}

type data struct {
	words   map[string]int
	letters map[string]int
	punct   map[string]int
	tod     map[string]int
}

func (d data) addWords(m *discordgo.Message) {
	f := func(c rune) bool {
		return unicode.IsPunct(c) || unicode.IsSpace(c)
	}

	cont := strings.FieldsFunc(m.Content, f)

	for _, word := range cont {
		val, exists := d.words[word]
		//.Println(word)
		if exists {
			//fmt.Print(": Exists\n")
			d.words[word] = val + 1
		} else {
			//fmt.Print(": Doesn't exist\n")
			d.words = map[string]int{word: 1}
		}

	}
	//fmt.Println(d.words)
}

func (d data) addLetters(m *discordgo.Message) {
	val, exists := d.letters[m.Content]
	if exists {
		d.letters[m.Content] = val + 1
	} else {
		d.letters[m.Content] = 1
	}
}

func (d data) addpunct(m *discordgo.Message) {
	val, exists := d.punct[m.Content]
	if exists {
		d.punct[m.Content] = val + 1
	} else {
		d.punct[m.Content] = 1
	}
}
