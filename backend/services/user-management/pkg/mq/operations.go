package mq

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

type MessageProcessor interface {
	PublishMessage(info MessageAction) error
	Close()
}

func (mq *MQ) Close() {
	if mq.channel != nil {
		mq.channel.Close()
	}
	if mq.conn != nil {
		mq.conn.Close()
	}
}

func (mq *MQ) PublishMessage(info MessageAction) error {
	body, err := json.Marshal(info)
	if err != nil {
		return err
	}
	err = mq.channel.Publish(
		"",            // exchange
		mq.queue.Name, // routing key (queue name)
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		return err
	}
	log.Printf("Message published: %s", body)
	return nil
}
