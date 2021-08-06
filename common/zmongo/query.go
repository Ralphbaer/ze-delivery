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

func getTypeName(s interface{}) string {
	if reflect.ValueOf(s).Kind() == reflect.Ptr {
		return reflect.TypeOf(s).Elem().Name()
	}
	return reflect.ValueOf(s).String()
}

