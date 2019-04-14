package internal

import (
	"time"
)

type Message struct {
	Cmd  string    `json:"cmd,omitempty"`
	Time time.Time `json:"timestamp,omitempty"`
	Data map[string]string `json:data,omitempty`
}
