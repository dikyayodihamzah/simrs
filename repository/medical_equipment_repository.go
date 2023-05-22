package repository

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/dikyayodihamzah/simrs.git/model/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	dbName        = os.Getenv("DB_NAME")
	collMeqEq     = os.Getenv("MEDICAL_EQUIPMENT_COLLECTION")
	setTimeout, _ = strconv.Atoi(os.Getenv("DB_TIMEOUT"))
)

type MedEqRepository interface {
	FindOne(c context.Context, params, value string) (domain.MedicalEquipment, error)
}

type medEqRepository struct {
	Client *mongo.Client
}

func (repository *medEqRepository) FindOne(c context.Context, params, value string) (domain.MedicalEquipment, error) {
	ctx, cancel := context.WithTimeout(c, time.Duration(setTimeout)*time.Second)
	defer cancel()

	db := repository.Client.Database(dbName)
	coll := db.Collection(collMeqEq)

	result := coll.FindOne(ctx, bson.M{params: value})

	var account domain.MedicalEquipment
	result.Decode(&account)
	return account, result.Err()
}
