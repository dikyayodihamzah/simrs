package web

import (
	"time"

	"github.com/dikyayodihamzah/simrs/model/domain"
)

type PatientResponse struct {
	EMRID       string             `json:"emr_id" bson:"emr_id"`
	NIK         string             `json:"nik" bson:"nik"`
	Name        string             `json:"name" bson:"name"`
	DoB         string             `json:"dob" bson:"dob"`
	Age         int                `json:"age" bson:"age"`
	Address     string             `json:"address" bson:"address"`
	Phone       string             `json:"phone" bson:"phone"`
	BloodType   string             `json:"blood_type" bson:"blood_type"`
	IsImpatient int                `json:"is_impatient" bson:"is_impatient"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	ImpatientAt time.Time          `json:"impatient_at" bson:"impatient_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
	Screenings  []domain.Screening `json:"screenings" bson:"screenings"`
}

func NewPatientResponse(emr domain.EMR) PatientResponse {
	return PatientResponse{
		EMRID:       emr.ID,
		NIK:         emr.Patient.NIK,
		Name:        emr.Patient.Name,
		DoB:         emr.Patient.DoB,
		Age:         emr.Patient.Age,
		Address:     emr.Patient.Address,
		Phone:       emr.Patient.Phone,
		BloodType:   emr.Patient.BloodType,
		CreatedAt:   emr.Patient.CreatedAt,
		ImpatientAt: emr.Patient.ImpatientAt,
		UpdatedAt:   emr.Patient.UpdatedAt,
		Screenings:  emr.Screenings,
	}
}

func NewAllPatientResponse(emrs []domain.EMR) []PatientResponse {
	var patientResponses []PatientResponse

	for _, emr := range emrs {
		patientResponses = append(patientResponses, NewPatientResponse(emr))
	}

	return patientResponses
}
