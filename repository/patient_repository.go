package repository

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/dikyayodihamzah/simrs/exception"
	"github.com/dikyayodihamzah/simrs/model/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var dbName = os.Getenv("DB_NAME")
var collPatient = os.Getenv("PATIENT_COLLECTION")
var setTimeout, _ = strconv.Atoi(os.Getenv("DB_TIMEOUT"))

type PatientRepository interface {
	CreatePatient(c context.Context, patient domain.Patient) (*mongo.InsertOneResult, error)
	FindOnePatient(c context.Context, params, value string) (domain.Patient, error)
	UpdatePatient(c context.Context, nik string, patient domain.Patient) error
	DeleteOne(c context.Context, nik string) error
}

type patientRepository struct {
	Client *mongo.Client
}

func NewPatientRepository(client *mongo.Client) PatientRepository {
	return &patientRepository{
		Client: client,
	}
}

func (repository *patientRepository) CreatePatient(c context.Context, patient domain.Patient) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(c, time.Duration(setTimeout)*time.Second)
	defer cancel()

	db := repository.Client.Database(dbName)
	coll := db.Collection(collPatient)

	result, err := coll.InsertOne(ctx, patient)
	if err != nil {
		return nil, exception.ErrInternalServer(err.Error())
	}

	return result, err
}

func (repository *patientRepository) FindOnePatient(c context.Context, params, value string) (domain.Patient, error) {
	ctx, cancel := context.WithTimeout(c, time.Duration(setTimeout)*time.Second)
	defer cancel()

	db := repository.Client.Database(dbName)
	coll := db.Collection(collPatient)

	result := coll.FindOne(ctx, bson.M{params: value})

	var patient domain.Patient
	result.Decode(&patient)
	return patient, result.Err()
}

func (repository *patientRepository) FindAllPatients(c context.Context) ([]domain.Patient, error) {
	ctx, cancel := context.WithTimeout(c, time.Duration(setTimeout)*time.Second)
	defer cancel()

	db := repository.Client.Database(dbName)
	coll := db.Collection(collPatient)

	result, err := coll.Find(ctx, bson.M{})
	if err != nil {
		return []domain.Patient{}, result.Err()
	}

	var patients []domain.Patient
	result.Decode(&patients)
	return patients, result.Err()
}

func (repository *patientRepository) UpdatePatient(c context.Context, nik string, patient domain.Patient) error {
	ctx, cancel := context.WithTimeout(c, time.Duration(setTimeout)*time.Second)
	defer cancel()

	db := repository.Client.Database(dbName)
	coll := db.Collection(collPatient)

	update := bson.M{
		"nik":        patient.NIK,
		"name":       patient.Name,
		"pdob":       patient.PDoB,
		"address":    patient.Address,
		"phone":      patient.Phone,
		"blood_type": patient.BloodType,
		"updated_at": patient.UpdatedAt,
	}

	if _, err := coll.UpdateByID(ctx, nik, bson.M{"$set": update}); err != nil {
		return exception.ErrInternalServer(err.Error())
	}

	return nil
}

func (repository *patientRepository) DeleteOne(c context.Context, nik string) error {
	ctx, cancel := context.WithTimeout(c, time.Duration(setTimeout)*time.Second)
	defer cancel()

	db := repository.Client.Database(dbName)
	coll := db.Collection(collPatient)

	if _, err := coll.DeleteOne(ctx, bson.M{"nik": nik}); err != nil {
		return exception.ErrInternalServer(err.Error())
	}

	return nil
}
