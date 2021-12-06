package main

import (
	"time"
)

type data struct {
	words   []rune
	letters []byte
	punct   []rune
	tod     []time.Time
}

//Maybe split this all up into multiple files? and add an interface in here
func (*data) addLetters(letter byte) {

}

func (*data) addWords(word rune) {

}

func (*data) addPunct(punct rune) {

}

func (*data) addToD(tod time.Time) {

}
