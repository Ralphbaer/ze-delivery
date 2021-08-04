package zmongo

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

// WithMatch adds a generic bson.D match to match map
var WithMatch = func(name string, v interface{}) MongoAggregateBuilderOption {
	return func(b *MongoAggregateBuilder) {
		b.Stage = bson.D{{Key: "$match", Value: bson.D{{Key: name, Value: v}}}}
	}
}

// WithSample adds the $sample stage that can randomly selects the specified number of documents from its input.
var WithSample = func(size int) MongoAggregateBuilderOption {
	return func(b *MongoAggregateBuilder) {
		b.Stage = bson.D{{Key: "$sample", Value: bson.D{{Key: "size", Value: 1}}}}
	}
}

// WithGroup adds the $group stage that groups input documents by the specified _id expression
// and for each distinct grouping, outputs a document. The _id field of each output document
// contains the unique group by value. The output documents can also contain computed fields
// that hold the values of some accumulator expression.
var WithGroup = func(name string, v interface{}) MongoAggregateBuilderOption {
	return func(b *MongoAggregateBuilder) {
		b.Stage = bson.D{{Key: "$group", Value: bson.D{{Key: "_id", Value: bson.D{{Key: name, Value: fmt.Sprintf("$%s", v)}}}}}}
	}
}

// WithCount adds the $count stage that can count of the number of documents input to the stage.
var WithCount = func(name string) MongoAggregateBuilderOption {
	return func(b *MongoAggregateBuilder) {
		b.Stage = bson.D{{Key: "$count", Value: name}}
	}
}

// WithGeoNear adds the $geoNear stage that can count of the number of documents input to the stage.
var WithGeoNear = func(name string) MongoAggregateBuilderOption {
	return func(b *MongoAggregateBuilder) {
	
		b.Stage = bson.D{{Key: "$geoNear", Value: bson.D{{Key: "near", Value: bson.D{{Key: "type", Value: "Point"}}}}}}
	}
}
