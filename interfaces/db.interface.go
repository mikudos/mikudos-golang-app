package interfaces

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// IDB IDB
type IDB interface {
	GetDB() *mongo.Database
	GetClient() *mongo.Client
	TestCollection() *mongo.Collection
}
