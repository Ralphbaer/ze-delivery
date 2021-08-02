package zmongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoQueryBuilderOption is a kind of interface for build options
type MongoQueryBuilderOption func(b *MongoQueryBuilder)

// MongoAggregateBuilderOption is a kind of interface for build options
type MongoAggregateBuilderOption func(b *MongoAggregateBuilder)

// MongoQueryBuilder builds a query for mongodb
type MongoQueryBuilder struct {
	Opts   *options.FindOptions
	Filter bson.M
	Sorts  map[string]int
}

// MongoAggregateBuilder builds an aggregate for mongodb
type MongoAggregateBuilder struct {
	Opts  *options.AggregateOptions
	Stage bson.D
}

// NewMongoQueryBuilder creates an instance of MongoQueryBuilder
func NewMongoQueryBuilder(opts ...MongoQueryBuilderOption) *MongoQueryBuilder {
	builder := &MongoQueryBuilder{
		Opts:   options.Find(),
		Filter: bson.M{},
		Sorts:  make(map[string]int),
	}
	for _, opt := range opts {
		opt(builder)
	}

	return builder
}

// NewMongoAggregateBuilder creates an instance of MongoQueryBuilder
func NewMongoAggregateBuilder(opts ...MongoAggregateBuilderOption) *MongoAggregateBuilder {
	a := &MongoAggregateBuilder{
		Opts:  options.Aggregate(),
		Stage: bson.D{},
	}
	for _, opt := range opts {
		opt(a)
	}

	return a
}

// With adds a new option to query builder
func (q *MongoQueryBuilder) With(opt MongoQueryBuilderOption) {
	opt(q)
}

// With adds a new option to aggregate builder
func (a *MongoAggregateBuilder) With(opt MongoAggregateBuilderOption) {
	opt(a)
}
