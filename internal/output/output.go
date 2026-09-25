package output

import (
	"encoding/json"
	"fmt"
)

type Event string

const (
	EventMessage  Event = "message"
	EventInspect  Event = "inspect"
	EventProgress Event = "progress"
	EventError    Event = "error"
)

type Message struct {
	Event           Event   `json:"event"`
	Message         any     `json:"message,omitempty"`
	PercentProgress float64 `json:"percent_progress,omitempty"`
	BytesProgress   int64   `json:"bytes_progress,omitempty"`
	Total           int64   `json:"total,omitempty"`
	FilePath        string  `json:"filepath,omitempty"`
}

func (m Message) MessageText() string {
	switch v := m.Message.(type) {
	case string:
		return v
	case fmt.Stringer:
		return v.String()
	default:
		return ""
	}
}

type Output struct {
	formatJSON bool
}

func (o Output) Println(msg Message) {
	if o.formatJSON {
		b, _ := json.Marshal(msg)
		fmt.Println(string(b))
		return
	}
	switch msg.Event {
	case EventMessage, EventInspect:
		fmt.Printf("%s\n", msg.MessageText())
	default:
		fmt.Printf("%s: %s\n", msg.Event, msg.MessageText())
	}
}

func (o Output) Progress(msg Message) {
	if o.formatJSON {
		b, _ := json.Marshal(msg)
		fmt.Println(string(b))
		return
	}
	message := fmt.Sprintf("%d bytes downloaded", msg.BytesProgress)
	if msg.Total > 0 {
		message = fmt.Sprintf("%.1f%% downloaded", float64(msg.BytesProgress)/float64(msg.Total)*100)
	}
	fmt.Printf("\r%s: %s", msg.Event, message)
}

func NewOutput(formatJSON bool) Output {
	return Output{
		formatJSON: formatJSON,
	}
}
