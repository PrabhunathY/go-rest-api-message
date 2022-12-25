package controller

import (
	"encoding/json"
	"message/errors"
	"message/model"
	"message/service"
	"net/http"
)

type controller struct{}

var (
	messageService service.MessageService = service.NewMessageService()
)

type MessageController interface {
	GetMessages(resp http.ResponseWriter, req *http.Request)
	AddMessages(resp http.ResponseWriter, req *http.Request)
}

func NewMessageController() MessageController {
	return &controller{}
}

func (*controller) GetMessages(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	messages, err := messageService.GetAllMessage()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error getting the posts"})
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(messages)
}

func (*controller) AddMessages(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	var message model.Message
	err := json.NewDecoder(req.Body).Decode(&message)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}

	err1 := messageService.ValidateMessage(&message)
	if err1 != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}

	result, err2 := messageService.CreateMessage(&message)
	if err2 != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error saving the post"})
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(result)
}
