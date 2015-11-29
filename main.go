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
	sinks := []Sink{}
	if v, ok := os.LookupEnv("INFLUX_URL"); ok {
		sinks = append(sinks, influx.New(v))
	}
	if v, ok := os.LookupEnv("GRAYLOG_ADDRESS"); ok {
		sinks = append(sinks, gelf.New("udp", v))
	}
	if len(sinks) == 0 {
		panic("no sinks configured: no point running")
	}
	for {
		var event event.Event
		source.Read(&event)
		for _, sink := range sinks {
			sink.Write(&event)
		}
	}
}
