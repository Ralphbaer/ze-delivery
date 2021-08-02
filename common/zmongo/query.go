package zmongo

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/Ralphbaer/ze-delivery/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// FindAll is a utility function which finds all records in given collection using mongo query builder
// The given slice must be a pre created slice of any type. Eg. make([]*struct,0)
func FindAll(ctx context.Context, coll *mongo.Collection, b *MongoQueryBuilder, slice interface{}) error {

	if coll == nil {
		return errors.New("the given collection (coll) is nil")
	}

	cursor, err := coll.Find(context.TODO(), b.Filter, b.Opts)

	if err != nil {
		return err
	}

	defer cursor.Close(context.TODO())

	if err := cursor.All(context.TODO(), slice); err != nil {
		return err
	}

	return nil
}

// Count is a utility function which returns the number of documents in the collection.
//  For a fast count of the documents in the collection, see the EstimatedDocumentCount method.
func Count(ctx context.Context, coll *mongo.Collection, b *MongoQueryBuilder) (int64, error) {
	if coll == nil {
		return 0, errors.New("the given collection (coll) is nil")
	}

	result, err := coll.CountDocuments(context.TODO(), b.Filter, nil)
	if err != nil {
		return 0, err
	}

	return result, nil
}


// FindOne finds a document in MongoDB using a filter
func FindOne(ctx context.Context, coll *mongo.Collection, filter interface{}, s interface{}) (*mongo.SingleResult, error) {
	sr := coll.FindOne(context.TODO(), filter)
	if err := sr.Err(); err != nil {
		switch e := err.Error(); e {
		case "mongo: no documents in result":
			return nil, common.EntityNotFoundError{
				EntityType: fmt.Sprintf("(%s)", getTypeName(s)),
				Err:        err,
			}
		default:
			return nil, err
		}
	}
	if err := sr.Decode(s); err != nil {
		return sr, err
	}
	return sr, nil
}

// FindByID finds a document in MongoDB by given ID in ObjectID format
func FindByID(ctx context.Context, coll *mongo.Collection, ID string, s interface{}) (*mongo.SingleResult, error) {

	objID, err := ObjectIDFromString(ID)

	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}

	return FindOne(ctx, coll, filter, s)
}

// ObjectIDFromString converts a HexID in string format to primitive.ObjectID
func ObjectIDFromString(HexID string) (primitive.ObjectID, error) {
	objID, err := primitive.ObjectIDFromHex(HexID)
	if err != nil {
		if err.Error() == "encoding/hex: odd length hex string" {
			return primitive.NilObjectID, errors.New("invalid ID format")
		}
		return primitive.NilObjectID, err
	}
	return objID, nil
}

// MustObjectIDFromString converts a HexID in string format to primitive.ObjectID anyway
func MustObjectIDFromString(hexID string) primitive.ObjectID {
	obj, err := ObjectIDFromString(hexID)
	if err != nil {
		return primitive.NilObjectID
	}
	return obj
}

// ObjectIDFromList converts a list of string in a list of primitive.ObjectID
func ObjectIDFromList(HexIDs []string) ([]primitive.ObjectID, error) {
	objIDs := make([]primitive.ObjectID, 0)
	for _, id := range HexIDs {
		i, err := ObjectIDFromString(id)
		if err != nil {
			return nil, err
		}
		objIDs = append(objIDs, i)
	}
	return objIDs, nil
}

func getTypeName(s interface{}) string {
	if reflect.ValueOf(s).Kind() == reflect.Ptr {
		return reflect.TypeOf(s).Elem().Name()
	}
	return reflect.ValueOf(s).String()
}

