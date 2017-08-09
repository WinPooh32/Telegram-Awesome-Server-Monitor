package main

import (
	"os/exec"
	"bufio"
	"strings"
	"net"
	"time"
	"fmt"
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

	//its very important to use this line, otherwise it will leak tty channels :C
	cmd.Wait()

	return newips[:]
}

func FixDate(date string) string {
	list := strings.Split(date, " ")
	listTime := strings.Split(list[3], ":")
	return fmt.Sprintf("%s %s %s %s:%s", list[0], list[1], list[2], listTime[0], listTime[1])
}

func MakeLastReport() string {
	cmd := exec.Command("utmpdump", "/var/log/wtmp")

	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	in := bufio.NewScanner(stdout)

	var report string = "```"

	for in.Scan() {
		list := strings.Split(in.Text(), "[")

		if list[1] != "7] " {continue}

		user := strings.TrimSpace(strings.TrimRight(list[4], "]  "))
		ip := strings.TrimSpace(strings.TrimRight(list[7], "]  "))
		date := strings.TrimSpace(strings.TrimRight(list[8], "]  "))

		report += fmt.Sprintf("\n%-6s  %-10s  %s", user, ip, FixDate(date))
	}

	cmd.Wait()

	return report + "\n```"
}

func StartMonitoringLast(lastChan chan []string){
	for true {
		last := UpdateLastLogins()

		if len(last) > 0{
			lastChan <-last
		}

		time.Sleep(DELAY)
	}
}