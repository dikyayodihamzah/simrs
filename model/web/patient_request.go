package web

import "time"

type NewPatientRequest struct {
	NIK       string    `json:"nik" bson:"nik"`
	Name      string    `json:"name" bson:"name"`
	PDoB      string    `json:"pdob" bson:"pdob"`
	Address   string    `json:"address" bson:"address"`
	Phone     string    `json:"phone" bson:"phone"`
	BloodType string    `json:"blood_type" bson:"blood_type"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
