package zmongo

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// WithObjectID adds a bson.M as "_id" = ObjectID(id) to find a document by it's ID
var WithObjectID = func(id string) MongoQueryBuilderOption {
	return func(b *MongoQueryBuilder) {
		objID, err := ObjectIDFromString(id)
		if err != nil {
			objID = primitive.NilObjectID
		}
		b.Filter["_id"] = objID
	}
}

// WithFilter adds a generic bson.M filter to filter map
var WithFilter = func(name string, filter interface{}) MongoQueryBuilderOption {
	k := reflect.TypeOf(filter).Kind()
	if (k == reflect.String || k == reflect.Bool || k == reflect.Int || k == reflect.Int64 || k == reflect.Float64) && reflect.ValueOf(filter).IsZero() {
		return func(b *MongoQueryBuilder) {}
	}

	return func(b *MongoQueryBuilder) {
		b.Filter[name] = filter
	}
}
