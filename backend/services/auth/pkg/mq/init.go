package mq

import (
	"auth-service/pkg/app"
	"fmt"

	"github.com/streadway/amqp"
)

type MQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func InitMQ(mqConfig app.MQConfig) (*MQ, error) {
	connStr := fmt.Sprintf("amqp://%s:%s@%s:%d/", mqConfig.Username, mqConfig.Password, mqConfig.Host, mqConfig.Port)
	conn, err := amqp.Dial(connStr)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
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
	if err != nil {
		return nil, err
	}

	return &MQ{
		conn:    conn,
		channel: channel,
	}, nil
}
