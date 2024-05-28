package socket

type Message struct {
	Room    string `json:"room"`
	Message string `json:"message"`
}
