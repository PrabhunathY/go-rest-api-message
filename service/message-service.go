package service

import (
	"errors"
	"math/rand"
	"message/model"
	"message/repository"
)

type MessageService interface {
	ValidateMessage(message *model.Message) error
	CreateMessage(message *model.Message) (*model.Message, error)
	GetAllMessage() ([]model.Message, error)
}

type service struct{}

var (
	repo repository.MessageRepository = repository.NewFirestoreRepository()
)

func NewMessageService() MessageService {
	return &service{}
}

func (*service) ValidateMessage(message *model.Message) error {
	if message == nil {
		err := errors.New("post is empy")
		return err
	}

	if message.Title == "" {
		err := errors.New("post title is empy")
		return err
	}
	return nil
}

func (*service) CreateMessage(message *model.Message) (*model.Message, error) {
	message.ID = rand.Int63()
	return repo.PostMessages(message)
}

func (*service) GetAllMessage() ([]model.Message, error) {
	return repo.GetMessages()
}
