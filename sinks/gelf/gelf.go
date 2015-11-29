package gelf

import (
	"encoding/json"
	"github.com/altaireltd/event-bridge/event"
	"net"
)

type gelfEvent struct {
	Version   string  `json:"version"`
	Timestamp float64 `json:"timestamp"`
	Host      string  `json:"host"`
	Priority  int64   `json:"level"`
	Unit      string  `json:"_unit"`
	Message   string  `json:"short_message"`
	Command   string  `json:"_command"`
}

type Sink struct {
	e *json.Encoder
}

func New(network, address string) *Sink {
	conn, err := net.Dial(network, address)
	if err != nil {
		panic(err)
	}
	return &Sink{json.NewEncoder(conn)}
}

func (s *Sink) Write(e *event.Event) {
	var out gelfEvent
	out.Version = "1.1"
	out.Host = e.Hostname
	out.Timestamp = float64(e.Timestamp.UnixNano()) / 1000000000.0
	out.Priority = e.Priority
	out.Unit = e.Unit
	out.Message = e.Message
	out.Command = e.Command
	s.e.Encode(out)
}
