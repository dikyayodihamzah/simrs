package web

import (
	"time"

	"github.com/dikyayodihamzah/simrs.git/model/domain"
)

type PatientResponse struct {
	NIK       string    `json:"nik" bson:"nik"`
	Name      string    `json:"name" bson:"name"`
	PDoB      string    `json:"pdob" bson:"pdob"`
	Address   string    `json:"address" bson:"address"`
	Phone     string    `json:"phone" bson:"phone"`
	BloodType string    `json:"blood_type" bson:"blood_type"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

func NewPatientResponse(patient domain.Patient) PatientResponse {
	return PatientResponse{
		NIK:       patient.NIK,
		Name:      patient.Name,
		PDoB:      patient.PDoB,
		Address:   patient.Address,
		Phone:     patient.Phone,
		BloodType: patient.BloodType,
		CreatedAt: patient.CreatedAt,
		UpdatedAt: patient.UpdatedAt,
	}
}

func NewAllPatientResponse(patients []domain.Patient) []PatientResponse {
	var patientResponses []PatientResponse

	for _, patient := range patients {
		patientResponses = append(patientResponses, NewPatientResponse(patient))
	}

	return patientResponses
}
