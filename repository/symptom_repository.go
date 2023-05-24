package repository

import (
	"context"
	"os"
	"time"

	"github.com/dikyayodihamzah/simrs/model/domain"
	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collSymptom = os.Getenv("SYMPTOM_COLLECTION")

type SymptomRepository interface {
	FindAllSymtoms(c context.Context) ([]domain.Symptom, error)
}

type symptomRepository struct {
	Client *mongo.Client
}

func NewSymptomRepository(clinet *mongo.Client) SymptomRepository {
	return &symptomRepository{
		Client: clinet,
	}
}

func (repository *symptomRepository) FindAllSymtoms(c context.Context) ([]domain.Symptom, error) {
	ctx, cancel := context.WithTimeout(c, time.Duration(setTimeout)*time.Second)
	defer cancel()

	db := repository.Client.Database(dbName)
	coll := db.Collection(collSymptom)

	// matchStage := bson.D{{Key: "$match", Value: bson.D{{}}}}

	result, err := coll.Find(ctx, bson.M{})
	if err != nil {
		return []domain.Symptom{}, err
	}

	var symptoms []domain.Symptom
	if err := result.All(context.TODO(), &symptoms); err != nil {
		panic(err)
	}

	return symptoms, nil
}
