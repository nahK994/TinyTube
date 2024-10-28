package mq

const (
	UserCreate = "Create"
	UserDelete = "Delete"
)

type CreateMessage struct {
	Email    string `json:"email"`
	Id       int    `json:"id"`
	Password string `json:"password"`
}

type MessageAction struct {
	ActionType string      `json:"actionType"`
	Message    interface{} `json:"message"`
}
