package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

func main() {
	var typeb string
	flag.StringVar(&typeb, "type", "do", "Cloud that will be used")
	flag.Parse()
	log.Debug().Msg("Args parsed")

	switch typeb {
	case "do":
		launchDo()
	default:
		fmt.Println("Type not supported:", typeb)
	}
}

type DigitalOceanToken struct {
	AccessToken string `yaml:"access-token"`
}

var configurationLocation = os.Getenv("HOME") + "/.config/doctl/config.yaml"

func launchDo() {
	token := DigitalOceanToken{}

	f, _ := os.Open(configurationLocation)

	content, _ := ioutil.ReadAll(f)

	yaml.Unmarshal(content, &token)

	fmt.Print(token)
}
