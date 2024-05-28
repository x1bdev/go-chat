package socket

import "time"

type Message struct {
	Room    string    `json:"room"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}
