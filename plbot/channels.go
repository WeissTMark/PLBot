package plbot

type channel interface {
	addWords()
	addLetters()
	addPunct()
	addToD()
}

type channels struct {
	id        string
	analytics data
}
