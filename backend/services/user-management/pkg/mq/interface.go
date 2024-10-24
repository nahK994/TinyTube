package mq

type MessageProcessor interface {
	PublishMessage(body string) error
}
