package models

type MessageJSON struct {
	Message string `json:"message"`
}

func NewErrorMsg(err error) MessageJSON {
	return MessageJSON{Message: err.Error()}
}
