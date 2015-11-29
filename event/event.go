package event

import (
	"time"
)

type Event struct {
	Timestamp time.Time
	Hostname  string
	Priority  int64
	Unit      string
	Message   string
	Command   string
}
