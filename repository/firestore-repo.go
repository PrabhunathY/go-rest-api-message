package repository

import (
	"context"
	"log"
	"message/model"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

const (
	projectId      string = "go-api-XXXXXX" // change your firebase project ID
	collectionName string = "posts"
)

type repo struct{}

// New Firestore Repository create a repository
func NewFirestoreRepository() MessageRepository {
	return &repo{}
}

func (r *repo) PostMessages(message *model.Message) (*model.Message, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}
	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    message.ID,
		"Title": message.Title,
		"Text":  message.Text,
	})

	if err != nil {
		log.Fatalf("Failed to add a new post: %v", err)
		return nil, err
	}

	return message, nil
}

func (r *repo) GetMessages() ([]model.Message, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}
	defer client.Close()

	var mMessages []model.Message
	iter := client.Collection(collectionName).Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			return nil, err
		}
		message := model.Message{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		mMessages = append(mMessages, message)
	}

	return mMessages, nil
}
