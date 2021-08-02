package repository

import (
	"context"

	"github.com/Ralphbaer/ze-delivery/common"
	"github.com/Ralphbaer/ze-delivery/common/zmongo"
	e "github.com/Ralphbaer/ze-delivery/partner/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// PartnerMongoRepositoryOption is the options to PartnerMongoRepository
type PartnerMongoRepositoryOption struct {
	DatabaseName   string
	CollectionName string
}

// PartnerMongoRepository represents a MongoDB implementation of PartnerRepository interface
type PartnerMongoRepository struct {
	connection *common.MongoConnection
	opts       *PartnerMongoRepositoryOption
}

// NewPartnerMongoRepository creates an instance of repository.PartnerMongoRepository
func NewPartnerMongoRepository(c *common.MongoConnection) *PartnerMongoRepository {
	return &PartnerMongoRepository{
		connection: c,
		opts: &PartnerMongoRepositoryOption{
			DatabaseName:   "partner",
			CollectionName: "partners",
		},
	}
}

// Find returns a specific Partner given an id
func (r PartnerMongoRepository) Find(ctx context.Context, eventID string) (*e.Partner, error) {
	coll, err := r.connection.ReadyCollection(r.opts.DatabaseName, r.opts.CollectionName)
	if err != nil {
		return nil, err
	}

	doc := &PartnerMongoModel{}
	b := zmongo.NewMongoQueryBuilder(
		zmongo.WithObjectID(eventID),
	)

	if _, err := zmongo.FindOne(ctx, coll, b.Filter, doc); err != nil {
		return nil, err
	}

	event := doc.ToEntity()

	return event, nil
}

// Save stores the given entity.Event into Mongo
func (r *PartnerMongoRepository) Save(ctx context.Context, p *e.Partner) (*string, error) {

	coll, err := r.connection.ReadyCollection(r.opts.DatabaseName, r.opts.CollectionName)
	if err != nil {
		return nil, err
	}

	doc := &PartnerMongoModel{}
	doc.FromEntity(p)

	res, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		if werr, ok := err.(mongo.WriteException); ok {
			if common.ContainsError(werr, 11000) {
				return nil, common.ErrMongoDuplicatedDocument
			}
		}
		return nil, err
	}

	id := res.InsertedID.(primitive.ObjectID).Hex()
	
	return &id, nil
}


/*
// FindNearest returns a specific Partner given an id
func (r PartnerMongoRepository) FindNearest(ctx context.Context, long, lat float64) (*e.Partner, error) {
	coll, err := r.connection.ReadyCollection(r.opts.DatabaseName, r.opts.CollectionName)
	if err != nil {
		return nil, err
	}

	doc := &PartnerMongoModel{}

	b := zmongo.NewMongoQueryBuilder(
		zmongo.WithFilter("location", ),
	)
	
	if _, err := zmongo.FindOne(ctx, bson.M{
		"location": bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{long, lat},
				},
				"$maxDistance": 50,
			},
		},
	}, nil); err != nil {
		return nil, err
	}

	event := doc.ToEntity()

	return event, nil
}*/