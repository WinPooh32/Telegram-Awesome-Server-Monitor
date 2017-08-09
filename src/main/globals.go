package main

import "time"

const(
	DELAY_SEC = 5
	DELAY = DELAY_SEC * time.Second
	POINTS = (60 / DELAY_SEC) + 1

	EVENT_KILL         = "kill"
	EVENT_ACTION       = "act"
	EVENT_REFRESH      = "refresh"
	EVENT_TO_MAIN      = "tomain"
	EVENT_TO_REALTIME  = "torealtime"
	EVENT_TO_STEPPED   = "tostep"
	EVENT_TO_LAST      = "tolast"

	REQ_SEND = iota
	REQ_DELETE
	REQ_EDIT

	RES_SEND = iota
	RES_DELETE
	RES_EDIT

	STATE_MAIN = iota
	STATE_REALTIME
	STATE_STEPPED
)