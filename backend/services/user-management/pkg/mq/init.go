package mq

import (
	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var channel *amqp.Channel

func InitRabbitMQ() error {
	var err error
	conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	defer conn.Close()

	channel, err = conn.Channel()
	if err != nil {
		return err
	}
	channel.Close()

	// Declare queue, exchange, etc.
	_, err = channel.QueueDeclare(
		"queue_name", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	return err
}
