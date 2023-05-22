package repository

import (
	"context"
	"os"
	"time"

	"github.com/dikyayodihamzah/simrs/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collImpatient = os.Getenv("IMPATIENT_COLLECTION")

type ImpatientRepository interface {
	CountPatients(c context.Context) (int64, error)
	CountPatientbyClass(c context.Context) ([]bson.M, error)
}

type impatientRepository struct {
	Client *mongo.Client
}

func NewImpatientRepository(client *mongo.Client) ImpatientRepository {
	return &impatientRepository{
		Client: client,
	}
}

func (repository *impatientRepository) CountPatients(c context.Context) (int64, error) {
	ctx, cancel := context.WithTimeout(c, time.Duration(setTimeout)*time.Second)
	defer cancel()

	db := repository.Client.Database(dbName)
	coll := db.Collection(collImpatient)

	result, err := coll.CountDocuments(ctx, bson.M{})
	if err != nil {
		exception.ErrInternalServer(err.Error())
	}

	return result, nil
}

func (repository *impatientRepository) CountPatientbyClass(c context.Context) ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(c, time.Duration(setTimeout)*time.Second)
	defer cancel()

	db := repository.Client.Database(dbName)
	coll := db.Collection(collImpatient)

	groupStage := bson.D{
		{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$class"},
			{Key: "patient_total", Value: bson.D{{Key: "$sum", Value: 1}}},
		}}}

	cursor, err := coll.Aggregate(ctx, mongo.Pipeline{groupStage})
	if err != nil {
		return []bson.M{}, exception.ErrInternalServer(err.Error())
	}

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return []bson.M{}, exception.ErrInternalServer(err.Error())
	}

	return results, nil
}
