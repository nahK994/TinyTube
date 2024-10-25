package mq

import (
	"log"
)

func (mq *MQ) ConsumeMessages() error {
	msgs, err := mq.channel.Consume(
		mq.queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			log.Printf("Received message: %s", d.Body)
			// Process the message here
		}
	}()
	return nil
}
