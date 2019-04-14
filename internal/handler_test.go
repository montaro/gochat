package internal

import (
	"testing"
	"time"
)

func TestHandleMessage(t *testing.T) {
	var now = time.Now()
	var formattedNow = now.Format(time.RFC3339)
	chatMsg := `{"cmd": "smsg","timestamp":"` + formattedNow + `","data":{
	"msg": "Hello World"}}`
	unmarshalChatMessage := Message{
		Cmd:  "smsg",
		Time: now,
		Data: map[string]string{"msg": "Hello World"},
	}
	var msg *Message
	msg = HandleMessage([]byte(chatMsg))
	if msg.Cmd != unmarshalChatMessage.Cmd {
		t.Errorf("Expected cmd: %v doesn't equal: %v", msg.Cmd, unmarshalChatMessage.Cmd)
	}
	if msg.Time.Equal(unmarshalChatMessage.Time) {
		t.Errorf("Expected time: %v doesn't equal: %v", msg.Time, unmarshalChatMessage.Time)
	}
	if msg.Data["msg"] != unmarshalChatMessage.Data["msg"] {
		t.Errorf("Expected message: %v doesn't equal: %v", msg.Data["msg"], unmarshalChatMessage.Data["msg"])
	}
}
