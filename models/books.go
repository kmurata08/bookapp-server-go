package models

import (
	"context"
	"os"

	"cloud.google.com/go/datastore"
)

type Book struct {
	ID       int64  `json:"id"`
	Name     string `json:"name" datastore:"Name,noindex"`
	Category int    `json:"category" datastore:"Category,noindex"`
}

type ModelClient struct {
	dsClient *datastore.Client
}

func NewClient() (*ModelClient, error) {
	// 参考URL(https://cloud.google.com/datastore/docs/datastore-api-tutorial?hl=ja)
	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, os.Getenv("DATASTORE_PROJECT_ID"))
	if err != nil {
		return nil, err
	}

	client := ModelClient{dsClient: dsClient}
	return &client, nil
}

func (client *ModelClient) ListBook(ctx context.Context) ([]*Book, error) {
	var books []*Book

	query := datastore.NewQuery("Book")
	keys, err := client.dsClient.GetAll(ctx, query, &books)
	if err != nil {
		return nil, err
	}

	for i, key := range keys {
		books[i].ID = key.ID
	}

	return books, nil
}

func (client *ModelClient) GetBook(ctx context.Context, bookId int64) (*Book, error) {
	k := datastore.IDKey("Book", bookId, nil)
	book := new(Book)
	if err := client.dsClient.Get(ctx, k, book); err != nil {
		return nil, err
	}

	return book, nil
}

func (client *ModelClient) CreateBook(ctx context.Context, book *Book) error {
	newKey := datastore.IncompleteKey("Book", nil)
	if _, err := client.dsClient.Put(ctx, newKey, book); err != nil {
		return err
	}

	return nil
}