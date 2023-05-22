package repository

import (
	"context"
	"os"
	"time"

	"github.com/dikyayodihamzah/simrs/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collRoom = os.Getenv("ROOM_COLLECTION")

type RoomRepository interface {
	CountRooms(c context.Context) (int64, error)
	CountNotEmptyRooms(c context.Context) ([]bson.M, error)
	CountBeds(c context.Context) ([]bson.M, error)
}

type roomRepository struct {
	Client *mongo.Client
}

func NewRoomRepository(client *mongo.Client) RoomRepository {
	return &roomRepository{
		Client: client,
	}
}

func (repository *roomRepository) CountRooms(c context.Context) (int64, error) {
	ctx, cancel := context.WithTimeout(c, time.Duration(setTimeout)*time.Second)
	defer cancel()

	db := repository.Client.Database(dbName)
	coll := db.Collection(collRoom)

	result, err := coll.CountDocuments(ctx, bson.M{})
	if err != nil {
		exception.ErrInternalServer(err.Error())
	}

	return result, nil
}

func (repository *roomRepository) CountNotEmptyRooms(c context.Context) ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(c, time.Duration(setTimeout)*time.Second)
	defer cancel()

	db := repository.Client.Database(dbName)
	coll := db.Collection(collRoom)

	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "person_stay", Value: bson.D{{Key: "$ne", Value: 0}}}}}}
	countStage := bson.D{{Key: "$count", Value: "room_total"}}

	cursor, err := coll.Aggregate(ctx, mongo.Pipeline{matchStage, countStage})
	if err != nil {
		return []bson.M{}, exception.ErrInternalServer(err.Error())
	}
	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return []bson.M{}, exception.ErrInternalServer(err.Error())
	}

	return results, nil
}

func (repository *roomRepository) CountBeds(c context.Context) ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(c, time.Duration(setTimeout)*time.Second)
	defer cancel()

	db := repository.Client.Database(dbName)
	coll := db.Collection(collRoom)

	groupStage := bson.D{
		{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$class"},
			{Key: "bed_total", Value: bson.D{{Key: "$sum", Value: "$capacity"}}},
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

// func (repository *roomRepository) CountRooms(c context.Context) (int64, error) {
// 	ctx, cancel := context.WithTimeout(c, time.Duration(setTimeout)*time.Second)
// 	defer cancel()

// 	db := repository.Client.Database(dbName)
// 	coll := db.Collection(collRoom)

// 	result, err := coll.CountDocuments(ctx, bson.M{param: value})

// 	return result, err
// }
// func (repository *roomRepository) CountRooms(c context.Context) (int64, error) {
// 	ctx, cancel := context.WithTimeout(c, time.Duration(setTimeout)*time.Second)
// 	defer cancel()

// 	db := repository.Client.Database(dbName)
// 	coll := db.Collection(collRoom)

// 	result, err := coll.CountDocuments(ctx, bson.M{param: value})

// 	return result, err
// }
