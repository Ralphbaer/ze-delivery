package zmongo

import "go.mongodb.org/mongo-driver/bson/primitive"

// MongoObjectIDGenerator represents the mongo implementation of id.IDGenerator
type MongoObjectIDGenerator struct {}

// New creates a new primitive.ObjectID
func (gen *MongoObjectIDGenerator) New() string {
	return primitive.NewObjectID().Hex()
}