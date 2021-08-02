package zmongo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

// Aggregate is a utility function which finds all records in given collection using mongo query builder
// The given slice must be a pre created slice of any type. Eg. make([]*struct,0)
func Aggregate(ctx context.Context, coll *mongo.Collection, b *MongoAggregateBuilder, slice interface{}) error {
	if coll == nil {
		return errors.New("the given collection (coll) is nil")
	}

	cursor, err := coll.Aggregate(context.TODO(), mongo.Pipeline{b.Stage}, b.Opts)
	if err != nil {
		return err
	}

	defer cursor.Close(context.TODO())

	if err := cursor.All(context.TODO(), slice); err != nil {
		return err
	}

	return nil
}
