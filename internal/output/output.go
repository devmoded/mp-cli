package output

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Event           string  `json:"event"` // TODO: Придумать что-то с типом событий
	Message         string  `json:"message,omitempty"`
	PercentProgress float64 `json:"percent_progress,omitempty"`
	BytesProgress   int64   `json:"bytes_progress,omitempty"`
	Total           int64   `json:"total,omitempty"`
	FilePath        string  `json:"filepath,omitempty"`
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
	case "message":
		fmt.Printf("%s\n", msg.Message)
	default:
		fmt.Printf("%s: %s\n", msg.Event, msg.Message)
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
