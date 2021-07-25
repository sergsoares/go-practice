package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/digitalocean/godo"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(os.Getenv("HOME") + "/.config/doctl/token.env")
	var dropletName string
	flag.StringVar(&dropletName, "name", "", "Give a name for your droplet")
	flag.Parse()

	client := godo.NewFromToken(os.Getenv("ACCESS_TOKEN"))

	createRequest := &godo.DropletCreateRequest{
		Name:     dropletName,
		Region:   "nyc3",
		Size:     "s-1vcpu-1gb",
		UserData: encodeFile(),
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

	ipv4, _ := newDroplet.PublicIPv4()
	fmt.Println("Access instance")
	fmt.Println("ssh root@" + ipv4)
	fmt.Println(newDroplet)
}

func encodeFile() string {
	f, _ := os.Open("cloud-init.yml")

	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	return string(content)
	// return base64.StdEncoding.EncodeToString(content)
}
