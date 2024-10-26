package mq

import (
	"encoding/json"
	"log"
)

type MessageProcessor interface {
	ConsumeMessages() error
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

func (mq *MQ) StartConsumeMessages() error {
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
		var info MessageAction
		for d := range msgs {
			if err := json.Unmarshal(d.Body, &info); err != nil {
				log.Fatal(err)
			}
			log.Printf("Received message: %v", info)
		}
	}()
	return nil
}
