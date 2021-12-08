package plbot

import (
	"errors"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Token string `json: "Token"`
	BotPrefix string `json: "BotPrefix"`
}

func loadConfig(fileName string) (c *Config, error) {
	file, err := ioutil.ReadFile(fileName)
	
	if err != nil {
		return c, err
	}

	err = json.Unmarshal(file, &c)

	if err != nil {
		return c, err
	}

	return c, nil
}