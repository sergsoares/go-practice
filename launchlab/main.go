package main

import (
	"bufio"
	"context"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/digitalocean/godo"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

type Params struct {
	name string
}

func main() {
	var typeb string
	var name string
	flag.StringVar(&typeb, "type", "do", "Cloud that will be used")
	flag.StringVar(&name, "name", "launchlab", "Name that will be used in Cloud Instance")
	flag.Parse()
	log.Debug().Msg("Args parsed")

	param := Params{
		name: name
	}
	switch typeb {
	case "do":
		launchDo(name)
	default:
		fmt.Println("Type not supported:", typeb)
	}
}

type DigitalOceanToken struct {
	AccessToken string `yaml:"access-token"`
}

var configurationLocation = os.Getenv("HOME") + "/.config/doctl/config.yaml"

// var dockerComposeCloudInit = "templates/cloud-init.yml"
var elasticDockercompose = "examples/elasticsearch.yml"

func launchDo(param Params) {
	token := DigitalOceanToken{}
	f, _ := os.Open(configurationLocation)
	content, _ := ioutil.ReadAll(f)
	yaml.Unmarshal(content, &token)

	client := godo.NewFromToken(token.AccessToken)
	createRequest := &godo.DropletCreateRequest{
		Name:     param.name,
		Region:   "nyc3",
		Size:     "s-1vcpu-1gb",
		UserData: generateCloudInit(elasticDockercompose),
		SSHKeys: []godo.DropletCreateSSHKey{
			{0, "43:7d:f6:a5:2e:15:78:4e:58:8a:f8:1a:ae:47:bf:5f"},
		},
		Image: godo.DropletCreateImage{
			Slug: "ubuntu-20-04-x64",
		},
	}
	ctx := context.TODO()

	newDroplet, _, err := client.Droplets.Create(ctx, createRequest)

	if err != nil {
		fmt.Print("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Droplet Created!!!")
	fmt.Println(newDroplet)
}

func encodeFile(path string) string {
	f, _ := os.Open(path)

	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	return string(content)
}

var baseyaml string = `#cloud-config
groups:
  - docker
users:
  - default
  # the docker service account
  - name: docker-service
    groups: docker
package_upgrade: true
packages:
  - apt-transport-https
  - ca-certificates
  - curl
  - gnupg-agent
  - software-properties-common`

var runcmd string = `# power_state:
#   mode: reboot
#   message: Restarting after installing docker & docker-compose
runcmd:
  # install docker following the guide: https://docs.docker.com/install/linux/docker-ce/ubuntu/
  - curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
  - sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
  - sudo apt-get -y update
  - sudo apt-get -y install docker-ce docker-ce-cli containerd.io
  - sudo systemctl enable docker
  # install docker-compose following the guide: https://docs.docker.com/compose/install/
  - sudo curl -L "https://github.com/docker/compose/releases/download/1.25.4/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
  - sudo chmod +x /usr/local/bin/docker-compose`

func generateCloudInit(templatePath string) string {
	// elasticConfig := "examples/elasticsearch.yml"
	f2, _ := os.Open(templatePath)
	content2, _ := ioutil.ReadAll(f2)

	str := base64.StdEncoding.EncodeToString([]byte(content2))

	var command string
	command += fmt.Sprint("echo ", str, " | base64 -d > /root/docker-compose.yml")
	command += "\n  - docker-compose up -d -f /root/docker-compose.yml"

	b := fmt.Sprintln(baseyaml, "\n", runcmd, "\n", " -", command)
	return b
}
