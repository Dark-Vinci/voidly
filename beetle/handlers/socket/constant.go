package socket

import "time"

const (
	readingLimit = 512
	writeWait    = 10 * time.Second
	pongWait     = 60 * time.Second
	pingPeriod   = 10 * time.Second
)
