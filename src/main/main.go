package main

import (
	"bytes"
	"log"
	"bufio"
	"os"
	"runtime"
)

func ReadToken(filename string)(string) {
	file,_ := os.Open(filename)
	reader := bufio.NewReader(file)
	lineBytes,_,_ := reader.ReadLine()
	file.Close()
	return string(lineBytes)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//channel for sharing plot image
	var monChan = make(chan *bytes.Buffer)
	//channel for last logins list
	var lastChan = make(chan []string)

	//init
	UpdateLastLogins()

	InitKeyboards()

	go StartMonitoringResources(monChan)
	go StartMonitoringLast(lastChan)

	log.Printf("Begin to serve bot")
	ServeBot(ReadToken("token.line"), monChan, lastChan)
}