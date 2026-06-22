package repository

import (
	"context"
	"packages/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type ItemRepository struct {
	Collection *mongo.Collection
}

func NewItemRepository(collection *mongo.Collection) *ItemRepository {
	return &ItemRepository{
		Collection: collection,
	}
}

func (r *ItemRepository) CreateItem(item *models.Item) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := r.Collection.InsertOne(ctx, item)
	if err != nil {
		return err
	}

	return nil
}
