package mq

import (
	"log"

	"github.com/streadway/amqp"
)

func (mq *MQ) PublishMessage(body string) error {
	err := mq.channel.Publish(
		"",           // exchange
		"queue_name", // routing key (queue name)
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Printf("Failed to publish message: %s", err)
		return err
	}
	log.Printf("Message published: %s", body)
	return nil
}
