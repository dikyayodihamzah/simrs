package repository

import (
	"context"
	"os"
	"time"

	"github.com/dikyayodihamzah/simrs.git/exception"
	"github.com/dikyayodihamzah/simrs.git/model/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collPatient = os.Getenv("PATIENT_COLLECTION")

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
		"name":        patient.Name,
		"patientname": patient.patientname,
		"email":       patient.Email,
		"password":    patient.Password,
		"phone":       patient.Phone,
		"is_merchant": patient.IsMerchant,
		"updated_at":  time.Now(),
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

	if _, err := coll.DeleteOne(ctx, bson.M{"_id": nik}); err != nil {
		return exception.ErrInternalServer(err.Error())
	}

	return nil
}
