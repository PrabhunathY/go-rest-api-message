package repository

import "message/model"

type MessageRepository interface {
	PostMessages(message *model.Message) (*model.Message, error)
	GetMessages() ([]model.Message, error)
}
