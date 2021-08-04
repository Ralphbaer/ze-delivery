package repository

import (
	"context"

	"github.com/Ralphbaer/ze-delivery/common"
	"github.com/Ralphbaer/ze-delivery/common/zmongo"
	e "github.com/Ralphbaer/ze-delivery/partner/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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



// FindNearest returns the nearest Partner given a long & lat
func (r PartnerMongoRepository) FindNearest(ctx context.Context, long, lat float64) (*e.Partner, error) {
	coll, err := r.connection.ReadyCollection(r.opts.DatabaseName, r.opts.CollectionName)
	if err != nil {
		return nil, err
	}


	var docs []*PartnerMongoModel
	var docResult *e.Partner

	pipeline := []bson.M{
		{
			"$geoNear": bson.M{
				"includeLocs":   "address",
				"distanceField": "distance",
				"spherical":     true,
				"near": bson.M{
					"type":        "Point",
					"coordinates": []float64{long, lat},
				},
				"query": bson.M{ 
					"coverageArea": bson.M{ 
						"$geoIntersects": bson.M{ 
							"$geometry": bson.M{ 
								"type": "Point", 
								"coordinates": []float64{long, lat},
								}, 
								},
								},
							},

			},
		},
		{
			"$limit": 1,
		},
	}

	cursor, err := coll.Aggregate(ctx, pipeline, options.Aggregate())
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &docs); err != nil {
		return nil, err
	}

	if len(docs) != 0 {
		docResult = docs[0].ToEntity()
	}

	return docResult, nil
}