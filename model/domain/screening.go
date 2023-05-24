package domain

import "time"

type Screening struct {
	ID       string    `json:"id" bson:"_id"`
	Patient  Patient   `json:"patient" bson:"patient"`
	Symptoms []Symptom `json:"symptoms" bson:"symptoms"`
	Diseases []Disease `json:"diseases" bson:"diseases"`
	AddInf   string    `json:"add_inf" bson:"add_inf"`
	Date     time.Time `json:"date" bson:"date"`
}
