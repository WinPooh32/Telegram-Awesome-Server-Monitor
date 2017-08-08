package main

import (
	"os/exec"
	"bufio"
	"strings"
	"net"
	"time"
)

//key is ip, values is username
var Logins = make(map[string] string)

func UpdateLastLogins() []string{
	cmd := exec.Command("last", "-w", "-a", "-i")
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	in := bufio.NewScanner(stdout)

	newips := make([]string, 0)

	for in.Scan() {
		list := strings.Split(in.Text(), " ")

		user := list[0]
		ip := list[len(list) - 1]

		parsed := net.ParseIP(ip)
		if parsed.To4() == nil {
			continue
		}

		_, ok := Logins[ip]

		if !ok{
			Logins[ip] = user
			newips = append(newips, ip)
		}
	}

	return newips[:]
}

func StartMonitoringLast(lastChan chan []string){
	for true {
		last := UpdateLastLogins()

		if len(last) > 0{
			lastChan <-last
		}

		time.Sleep(time.Second * 5)
	}
}