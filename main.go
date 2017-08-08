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
	return string(lineBytes)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//channel for sharing plot image
	var monChan = make(chan *bytes.Buffer)

	go StartMonitoring(monChan)

	log.Printf("Begin to serve bot")
	ServeBot(ReadToken("token.line"), monChan)
}