package repository

import (
	"context"
	"os"
	"time"

	"github.com/dikyayodihamzah/simrs/model/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collEMR = os.Getenv("EMR_COLLECTION")

type EMRRepository interface {
	FindOneEMR(c context.Context, key, value string) (domain.EMR, error)
}

type emrRepository struct {
	Client *mongo.Client
}

func NewEMRRepository(client *mongo.Client) EMRRepository {
	return &emrRepository{
		Client: client,
	}
}

func (repository *emrRepository) FindOneEMR(c context.Context, key, value string) (domain.EMR, error) {
	ctx, cancel := context.WithTimeout(c, time.Duration(setTimeout)*time.Second)
	defer cancel()

	db := repository.Client.Database(dbName)
	coll := db.Collection(collEMR)

	result := coll.FindOne(ctx, bson.M{key: value})

	var emr domain.EMR
	result.Decode(&emr)
	return emr, result.Err()
}
