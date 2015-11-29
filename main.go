package main

import (
	"github.com/altaireltd/event-bridge/event"
	"github.com/altaireltd/event-bridge/sinks/gelf"
	"github.com/altaireltd/event-bridge/sinks/influx"
	"github.com/altaireltd/event-bridge/sources/journald"
	"os"
)

type Source interface {
	Read(*event.Event)
}

type Sink interface {
	Write(*event.Event)
}

func main() {
	source := journald.New()
	sinks := []Sink{
		influx.New(os.Getenv("INFLUX_URL")),
		gelf.New("udp", os.Getenv("GRAYLOG_ADDRESS")),
	}
	for {
		var event event.Event
		source.Read(&event)
		for _, sink := range sinks {
			sink.Write(&event)
		}
	}
}
