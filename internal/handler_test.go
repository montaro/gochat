package internal

import (
	"fmt"
	"testing"
	"time"
)

func TestHandleMessage(t *testing.T) {
	var now = time.Now()
	var formattedNow = now.Format(time.RFC3339)
	//var formattedNow = now.UTC().String()
	//fmt.Println("Now", now)
	chatMsg := `{"msg": "Hello World!","timestamp":"` + formattedNow + `"}`
	fmt.Printf("msg: %v", chatMsg)

	unmarshalChatMessage := Message{
		Msg:  "Hello World!",
		Time: now,
	}

	var msg *Message
	msg = HandleMessage([]byte(chatMsg))
	println()
	fmt.Printf("msg: %v", *msg)
	println()
	fmt.Printf("msg: %v", unmarshalChatMessage)
	println()

	if msg.Msg != unmarshalChatMessage.Msg {
		t.Errorf("Expected message: %v doesn't equal: %v", msg.Msg, unmarshalChatMessage.Msg)
	}
	if msg.Time.Equal(unmarshalChatMessage.Time) {
		t.Errorf("Expected message: %v doesn't equal: %v", msg.Time, unmarshalChatMessage.Time)
	}
}
