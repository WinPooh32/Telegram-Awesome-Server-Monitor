package main

import "time"

var(
	DELAY_SEC = 5
	DELAY = time.Duration(DELAY_SEC) * time.Second
	POINTS = (60 / DELAY_SEC) + 1
)

const(
	EVENT_KILL         = "kill"
	EVENT_ACTION       = "act"
	EVENT_REFRESH      = "refresh"
	EVENT_TO_MAIN      = "tomain"
	EVENT_TO_REALTIME  = "torealtime"
	EVENT_TO_STEPPED   = "tostep"
	EVENT_TO_LAST      = "tolast"
	EVENT_TO_SETTINGS  = "tosettings"
	EVENT_SET_3        = "3sec"
	EVENT_SET_5        = "5sec"
	EVENT_SET_10       = "10sec"

	REQ_SEND = iota
	REQ_DELETE
	REQ_EDIT

	RES_SEND = iota
	RES_DELETE
	RES_EDIT

	STATE_MAIN = iota
	STATE_REALTIME
	STATE_STEPPED
	STATE_WAIT_REFRESH
)

func SetDelay(delay int){
	DELAY_SEC = delay
	DELAY = time.Duration(delay) * time.Second
}