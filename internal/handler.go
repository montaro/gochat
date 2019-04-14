package internal

import "encoding/json"

func HandleMessage(message []byte) *Message {
	var msg *Message
	json.Unmarshal(message, &msg)
	return msg
}
