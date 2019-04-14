package internal

func (message *Message) validateMessage() bool {
	if message.Cmd == "" || message.Time.IsZero() || len(message.Data) == 0 {
		return false
	}
	switch message.Cmd {
	case "sname":
		return message.validateSnameCmd()
	case "scname":
		return message.validateScnameCmd()
	case "smsg":
		return message.validateSmsgCmd()
	case "rmsg":
		return message.validateRmsgCmd()
	}
	return false
}

func (message *Message) validateSnameCmd() bool {
	if message.Data["name"] != "" {
		return true
	}
	return false
}

func (message *Message) validateScnameCmd() bool {
	if message.Data["name"] != "" {
		return true
	}
	return false
}

func (message *Message) validateSmsgCmd() bool {
	if message.Data["msg"] != "" {
		return true
	}
	return false
}

func (message *Message) validateRmsgCmd() bool {
	if message.Data["msg"] != "" {
		return true
	}
	return false
}
