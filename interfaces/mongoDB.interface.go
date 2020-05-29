package interfaces

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// IMongoDB MongoDB interface
type IMongoDB interface {
	GetDB() *mongo.Database
	GetClient() *mongo.Client
	GetCollection(collectionName string) *mongo.Collection
}
