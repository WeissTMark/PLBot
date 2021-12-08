/*
	Author: Charles Shook
	Description: Structure for loading in discord bot config.
*/

package plbot

import (
	"encoding/json"
	"io/ioutil"
)


type Config struct {
	Token string `json: "Token"`
	BotPrefix string `json: "BotPrefix"`
}

var config = &Config{}

func LoadConfig(fileName string) (error) {
	file, err := ioutil.ReadFile(fileName)

	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &config)

	if err != nil {
		return err
	}

	return nil
}