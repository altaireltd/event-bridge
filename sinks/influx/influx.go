package influx

import (
	"fmt"
	"github.com/altaireltd/event-bridge/event"
	client "github.com/influxdb/influxdb/client/v2"
)

type Sink struct {
	c client.Client
}

func New(url string) *Sink {
	if c, err := client.NewHTTPClient(client.HTTPConfig{Addr: url}); err != nil {
		panic(err)
	} else {
		return &Sink{c}
	}
}

func (s *Sink) Write(e *event.Event) {
	batch, err := client.NewBatchPoints(client.BatchPointsConfig{Database: "testing", Precision: "s"})
	if err != nil {
		panic(err)
	}
	tags := map[string]string{
		"unit":     e.Unit,
		"host":     e.Hostname,
		"priority": fmt.Sprint(e.Priority),
	}
	fields := map[string]interface{}{
		"value":   1,
		"message": fmt.Sprint(e.Message),
		"command": e.Command,
	}
	point, err := client.NewPoint("event", tags, fields, e.Timestamp)
	if err != nil {
		panic(err)
	}
	batch.AddPoint(point)
	err = s.c.Write(batch)
	if err != nil {
		panic(err)
	}
}
