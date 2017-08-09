package main

import "time"

const(
	DELAY_SEC = 5
	DELAY = DELAY_SEC * time.Second
	POINTS = (60 / DELAY_SEC) + 1
)