package common

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConnection is a hub which deal with mongodb connections.
type MongoConnection struct {
	ConnectionString string
	Database         string
	Client           *mongo.Client
	Connected        bool
	Verbose          bool
}

// Connect keeps a singleton connection with mongodb.
func (r *MongoConnection) Connect(database string) (*mongo.Client, error) {

	if r.Connected {
		return r.Client, nil
	}

	if r.Verbose {
		log.Println("Connecting to mongo database...")
	}

	clientOptions := options.Client().ApplyURI(r.ConnectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return nil, err
	}

	r.Client = client

	log.Println("Connected to mongo [ok]")

	r.Connected = true

	return r.Client, err
}

// ReadyCollection connects to database and return a ready collection
func (r *MongoConnection) ReadyCollection(database string, collection string) (*mongo.Collection, error) {
	c, err := r.Connect(database)
	if err != nil {
		return nil, err
	}
	return c.Database(database).Collection(collection), nil
}
