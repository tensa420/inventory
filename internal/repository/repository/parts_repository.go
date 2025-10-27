package repository

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"inventory/internal/entity"
	"inventory/internal/repository/model"
)

type InventoryRepository struct {
	collection *mongo.Collection
}

func NewInventoryRepository(db *mongo.Database) *InventoryRepository {
	collection := db.Collection("Inventory")

	indexModels := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "title", Value: 1}},
			Options: options.Index().SetUnique(false),
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.Indexes().CreateMany(ctx, indexModels)
	if err != nil {
		log.Fatalf("Failed to create indexes %v", err)
	}
	return &InventoryRepository{
		collection: collection,
	}
}

func (r *InventoryRepository) GetPart(partUUID string) (*entity.Part, error) {
	var part model.Part

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.collection.FindOne(ctx, bson.M{"UUID": partUUID}).Decode(&part)
	if err != nil {
		return &entity.Part{}, errors.Wrap(err, "part not found")
	}
	return model.ConvertRepoModelToEntity(part), nil
}

func (r *InventoryRepository) ListParts(filter entity.PartsFilter) ([]*entity.Part, error) {
	parts := make([]*entity.Part, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for _, uuid := range filter.UUIDS {
		var part model.Part
		err := r.collection.FindOne(ctx, bson.M{"UUID": uuid}).Decode(&part)
		if err != nil {
			return nil, errors.Wrap(err, "part not found")
		}
		parts = append(parts, model.ConvertRepoModelToEntity(part))
	}

	for _, name := range filter.Names {
		var part model.Part
		err := r.collection.FindOne(ctx, bson.M{"Name": name}).Decode(&part)
		if err != nil {
			return nil, errors.Wrap(err, "part not found")
		}
		parts = append(parts, model.ConvertRepoModelToEntity(part))
	}
	for _, tag := range filter.Tags {
		var part model.Part
		err := r.collection.FindOne(ctx, bson.M{"Tags": tag}).Decode(&part)
		if err != nil {
			return nil, errors.Wrap(err, "part not found")
		}
		parts = append(parts, model.ConvertRepoModelToEntity(part))
	}
	for _, cate := range filter.Categories {
		var part model.Part
		err := r.collection.FindOne(ctx, bson.M{"Category": cate}).Decode(&part)
		if err != nil {
			return nil, errors.Wrap(err, "part not found")
		}
		parts = append(parts, model.ConvertRepoModelToEntity(part))
	}
	for _, manu := range filter.ManufacturerContries {
		var part model.Part
		err := r.collection.FindOne(ctx, bson.M{"Manufacturer": manu}).Decode(&part)
		if err != nil {
			return nil, errors.Wrap(err, "part not found")
		}
		parts = append(parts, model.ConvertRepoModelToEntity(part))
	}
	return parts, nil
}
