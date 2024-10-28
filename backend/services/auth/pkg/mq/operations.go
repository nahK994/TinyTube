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
		var info MessageAction
		for d := range msgs {
			if err := json.Unmarshal(d.Body, &info); err != nil {
				log.Printf("Failed to unmarshal message: %v", err)
				// d.Nack(false, false) // Do not acknowledge the message, allow re-queuing
				continue // Skip to the next message
			}
			fmt.Println("received message", info)
			var processErr error
			switch info.ActionType {
			case UserCreate:
				processErr = handler.CreateUser(db.UserCreate{
					ID:       info.Message.Id,
					Email:    info.Message.Email,
					Password: info.Message.Password,
				})
			case ChangePassword:
				processErr = handler.UpdatePassword(db.PasswordUpdate{
					Email:    info.Message.Email,
					Password: info.Message.Password,
				})
			case UserDelete:
				processErr = handler.DeleteUser(info.Message.Id)
			}

			if processErr != nil {
				log.Printf("Failed to process message %v: %v", info, processErr)
				d.Nack(false, false) // Do not acknowledge the message
			}
			// else {
			// 	d.Ack(false) // Acknowledge the message
			// }
		}
	}()
	return nil
}
