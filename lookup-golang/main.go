package main

import (
	"fmt"
	"os/exec"
	"strings"
	"strconv"
	"net"
)

type PingResult struct {
	Ip   net.IP  `json:"ip"`
	Min  float64 `json:"min"`
	Avg  float64 `json:"avg"`
	Max  float64 `json:"max"`
	Mdev float64 `json:"mdev"`
}

func main()  {
	fmt.Println(validateConnection(net.ParseIP("1.1.1.1")))
	fmt.Println(validateConnection(net.ParseIP("8.8.8.8")))
	fmt.Println(validateConnection(net.ParseIP("127.0.0.1")))
}

func validateConnection(addr net.IP) PingResult {
	// Ping
	command := fmt.Sprintf("ping %s -c 5 -W 3 | grep rtt", addr.String())
	cmd := exec.Command("/bin/sh", "-c", command)
	
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
	}

	// Parse ping result
	s := strings.Split(string(stdout), " ")
	x := strings.Split(s[3], "/")
	x0, _ := strconv.ParseFloat(x[0], 64)
	x1, _ := strconv.ParseFloat(x[1], 64)
	x2, _ := strconv.ParseFloat(x[2], 64)
	x3, _ := strconv.ParseFloat(x[3], 64)

	// Show answer
	res := PingResult{
		Ip: addr,
		Min: x0, 
		Avg: x1, 
		Max: x2, 
		Mdev: x3,
	}

	return res
}