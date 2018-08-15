## Telegram Awesome Server Monitor
Telegram bot for monitoring remote server in realtime.

Rendered image example:

[![screenshot](https://github.com/WinPooh32/AwesomeServerMonitor/raw/readmedata/photo_2017-08-07_15-21-32.jpg)]()

## Features
* CPU and RAM usage realtime graph
* Sends report when user connected from new unknown ip
* Shows list of last connected users

## Install and Build
Install dependencies:
```
$ go get -u github.com/wcharczuk/go-chart \
github.com/shirou/gopsutil \
github.com/go-telegram-bot-api/telegram-bot-api
```
Build:
```
$ cd AwesomeServerMonitor

$ touch token.line
$ echo 'YOUR_BOT_TOKEN' > token.line

$ go build *.go
```
