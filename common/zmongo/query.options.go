package zmongo

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// WithSort sorts the results by given sort argument
// You can use any of these options for ascending: [asc, ASC, 1]
// You can use any of these options for descending: [desc, DESC, -1]
var WithSort = func(field string, sort string) MongoQueryBuilderOption {
	return func(b *MongoQueryBuilder) {
		m := map[string]int{
			"asc":  1,
			"ASC":  1,
			"1":    1,
			"desc": -1,
			"DESC": -1,
			"-1":   -1,
		}
		sortOrder := -1
		if s, found := m[sort]; found {
			sortOrder = s
		}
		b.Sorts[field] = sortOrder
		b.Opts.SetSort(b.Sorts)
	}
}

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
