package mq

type MessageProcessor interface {
	ConsumeMessages() error
}
