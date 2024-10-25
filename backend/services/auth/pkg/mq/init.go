package mq

import (
	"auth-service/pkg/app"
	"fmt"

	"github.com/streadway/amqp"
)

type MQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

func InitMQ(mqConfig app.MQConfig) (*MQ, error) {
	connStr := fmt.Sprintf("amqp://%s:%s@%s:%d/", mqConfig.Username, mqConfig.Password, mqConfig.Host, mqConfig.Port)
	conn, err := amqp.Dial(connStr)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	queue, err := channel.QueueDeclare(
		"queue_name",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		conn.Close()
		channel.Close()
		return nil, err
	}

	return &MQ{
		conn:    conn,
		channel: channel,
		queue:   queue,
	}, nil
}
