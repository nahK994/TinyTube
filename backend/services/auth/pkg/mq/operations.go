package mq

import (
	"auth-service/pkg/db"
	"auth-service/pkg/utils"
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
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

func (mq *MQ) Start() error {
	return mq.startConsumeMessages()
}

func (mq *MQ) startConsumeMessages() error {
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

	go mq.processMessages(msgs)
	return nil
}

func (mq *MQ) processMessages(msgs <-chan amqp.Delivery) {
	for d := range msgs {
		var info MessageAction
		if err := json.Unmarshal(d.Body, &info); err != nil {
			log.Printf("Failed to unmarshal message: %v", err)
			d.Nack(false, false)
			continue
		}

		fmt.Println("Received message:", info)
		if err := mq.processMessage(info); err != nil {
			log.Printf("Failed to process message %v: %v", info, err)
			d.Nack(false, false)
		} else {
			d.Ack(false)
		}
	}
}

func (mq *MQ) processMessage(info MessageAction) error {
	switch info.ActionType {
	case UserCreate:
		return mq.processUserCreate(info)
	case UserDelete:
		return mq.processUserDelete(info)
	default:
		log.Printf("Unknown ActionType: %v", info.ActionType)
		return fmt.Errorf("unknown action type")
	}
}

func (mq *MQ) processUserCreate(info MessageAction) error {
	var msg CreateMessage
	if err := parseMessage(info, &msg); err != nil {
		return err
	}

	var hashedPassword string
	err := utils.HashPassword(msg.Password, &hashedPassword)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	return mq.repo.CreateUser(&db.UserCreate{
		ID:       msg.Id,
		Email:    msg.Email,
		Password: hashedPassword,
	})
}

func (mq *MQ) processUserDelete(info MessageAction) error {
	var id int
	if err := parseMessage(info, &id); err != nil {
		return err
	}
	return mq.repo.DeleteUser(int(id))
}

func parseMessage(info MessageAction, v interface{}) error {
	switch info.ActionType {
	case UserCreate:
		messageBytes, err := json.Marshal(info.Message)
		if err != nil {
			return fmt.Errorf("failed to marshal message: %w", err)
		}
		if err := json.Unmarshal(messageBytes, v); err != nil {
			return fmt.Errorf("failed to unmarshal message: %w", err)
		}

	case UserDelete:
		switch id := info.Message.(type) {
		case int:
			if intPtr, ok := v.(*int); ok {
				*intPtr = id
			} else {
				return fmt.Errorf("v is not of type *int for UserDelete")
			}
		default:
			return fmt.Errorf("unexpected type for UserDelete: %T", info.Message)
		}
	}
	return nil
}
