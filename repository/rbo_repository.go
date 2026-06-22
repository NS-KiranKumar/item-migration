package repository

import (
	"context"
	"errors"
	"packages/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RboRepository struct {
	Collection *mongo.Collection
}

func NewRboRepository(collection *mongo.Collection) *RboRepository {
	return &RboRepository{
		Collection: collection,
	}
}

func (r *RboRepository) CreateRbo(rbo *models.RBO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := r.Collection.InsertOne(ctx, rbo)
	if err != nil {
		return err
	}
	return nil
}

func (r *RboRepository) GetRbobyID(id primitive.ObjectID) (*models.RBO, error) {
	var rbo models.RBO
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{
		"_id": id,
	}
	err := r.Collection.FindOne(ctx, filter).Decode(&rbo)
	if err != nil {
		return nil, err
	}

	return &rbo, nil
}

func (r *RboRepository) DeleteRbobyID(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{
		"_id": id,
	}
	result, err := r.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("rbo not found")
	}
	return nil
}

func (r *RboRepository) UpdateRbobyID(id primitive.ObjectID, rbo *models.RBO) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"_id": id,
	}

	update := bson.M{
		"$set": bson.M{
			"rboCode":   rbo.RBOCode,
			"rboName":   rbo.RBOName,
			"updatedAt": rbo.UpdatedAt,
		},
	}

	result, err := r.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func (r *RboRepository) GetallRbos() ([]models.RBO, error) {
	var rbos []models.RBO
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &rbos)
	if err != nil {
		return nil, err
	}
	return rbos, err
}
