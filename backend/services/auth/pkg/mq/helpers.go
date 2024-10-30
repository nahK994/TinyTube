package mq

const (
	UserCreate = "Create"
	UserDelete = "Delete"
)

type MessageAction struct {
	ActionType string      `json:"actionType"`
	Message    interface{} `json:"message"`
}
