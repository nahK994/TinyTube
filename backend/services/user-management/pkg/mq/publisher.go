package mq

import (
	"log"

	"github.com/streadway/amqp"
)

func (mq *MQ) PublishMessage(body string) error {
	err := mq.channel.Publish(
		"",            // exchange
		mq.queue.Name, // routing key (queue name)
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		return err
	}
	log.Printf("Message published: %s", body)
	return nil
}
