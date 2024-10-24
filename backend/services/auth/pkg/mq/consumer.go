package mq

import (
	"log"
)

func (mq *MQ) ConsumeMessages() {
	msgs, err := mq.channel.Consume(
		"queue_name", // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
		return
	}

	go func() {
		for d := range msgs {
			log.Printf("Received message: %s", d.Body)
			// Process the message here
		}
	}()
}
