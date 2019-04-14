package internal

import (
	"testing"
	"time"
)

func TestValidateSnameCmd(t *testing.T) {
	var now = time.Now()
	unmarshalChatMessage := Message{
		Cmd:  "smsg",
		Time: now,
		Data: map[string]string{"msg": "Hello World"},
	}
	if !unmarshalChatMessage.validateMessage() {
		t.Fail()
	}
}