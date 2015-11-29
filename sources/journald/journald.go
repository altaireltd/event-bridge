package journald

import (
	"encoding/json"
	"fmt"
	"github.com/altaireltd/event-bridge/event"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"time"
)

type journalEvent struct {
	Timestamp string      `json:"__REALTIME_TIMESTAMP"`
	Hostname  string      `json:"_HOSTNAME"`
	Priority  string      `json:"PRIORITY"`
	Facility  string      `json:"SYSLOG_IDENTIFIER"`
	Message   interface{} `json:"MESSAGE"`
	Command   string      `json:"_CMDLINE"`
	Unit      string      `json:"_SYSTEMD_UNIT"`
}

type Source struct {
	d *json.Decoder
}

var re = regexp.MustCompile(`(-[0-9]+)?(@[^@]+)?\.[^.]+`)

func New() *Source {
	cmd := exec.Command("journalctl", "-fn0", "-ojson")
	pipe, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	return &Source{json.NewDecoder(pipe)}
}

func (s *Source) Read(e *event.Event) {
	var in journalEvent
	if err := s.d.Decode(&in); err != nil {
		panic(err)
	}
	timestamp, err := strconv.ParseInt(in.Timestamp, 10, 64)
	if err != nil {
		panic(err)
	}
	priority, err := strconv.ParseInt(in.Priority, 10, 64)
	if err != nil {
		panic(err)
	}
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	unit := in.Unit
	unit = re.ReplaceAllString(unit, "")
	if unit == "" {
		unit = in.Facility
	}
	e.Timestamp = time.Unix(0, timestamp*1000)
	e.Hostname = hostname
	e.Priority = priority
	e.Unit = unit
	e.Message = fmt.Sprint(in.Message)
	e.Command = in.Command
}
