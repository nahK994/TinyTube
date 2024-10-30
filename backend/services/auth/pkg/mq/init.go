package mq

import (
	"auth-service/pkg/app"
	"auth-service/pkg/db"
	"fmt"

	"github.com/streadway/amqp"
)

type MQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
	repo    db.Repository
}

func InitMQ(mqConfig app.MQConfig) (*MQ, error) {
	connStr := fmt.Sprintf("amqp://%s:%s@%s:%d/", mqConfig.Username, mqConfig.Password, mqConfig.Host, mqConfig.Port)
	conn, err := amqp.Dial(connStr)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
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
		return nil, err
	}

	return &MQ{
		conn:    conn,
		channel: channel,
		queue:   queue,
	}, nil
}
