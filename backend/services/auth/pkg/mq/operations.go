package mq

import (
	"auth-service/pkg/db"
	"auth-service/pkg/handlers"
	"encoding/json"
	"fmt"
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

func (mq *MQ) Start(handler *handlers.Handler) error {
	return mq.startConsumeMessages(handler)
}

func (mq *MQ) startConsumeMessages(handler *handlers.Handler) error {
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
			var info MessageAction
			if err := json.Unmarshal(d.Body, &info); err != nil {
				log.Printf("Failed to unmarshal message: %v", err)
				continue // Skip to the next message
			}

			fmt.Println("received message", info)

			switch info.ActionType {
			case UserCreate:
				var msg CreateMessage
				messageBytes, _ := json.Marshal(info.Message) // Convert to JSON bytes
				if err := json.Unmarshal(messageBytes, &msg); err != nil {
					log.Printf("Failed to unmarshal message to CreateMessage: %v", err)
					continue
				}

				processErr := handler.CreateUser(db.UserCreate{
					ID:       msg.Id,
					Email:    msg.Email,
					Password: msg.Password,
				})
				if processErr != nil {
					log.Printf("Failed to process message %v: %v", info, processErr)
					d.Nack(false, false)
				}

			case UserDelete:
				if id, ok := info.Message.(float64); ok { // JSON numbers are float64 by default
					processErr := handler.DeleteUser(int(id))
					if processErr != nil {
						log.Printf("Failed to process message %v: %v", info, processErr)
						d.Nack(false, false)
					}
				} else {
					log.Printf("Invalid message type for UserDelete: %v", info.Message)
					d.Nack(false, false)
				}
			}
		}
	}()
	return nil
}
