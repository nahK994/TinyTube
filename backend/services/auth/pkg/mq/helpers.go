package mq

const (
	UserCreate     = "Create"
	UserDelete     = "Delete"
	ChangePassword = "ChangePassword"
)

type Message struct {
	Email    string `json:"email"`
	Id       int    `json:"id"`
	Password string `json:"password"`
}

type MessageAction struct {
	ActionType string  `json:"actionType"`
	Message    Message `json:"message"`
}
