package domain

import "time"

type Patient struct {
	NIK         string    `json:"nik" bson:"nik"`
	Name        string    `json:"name" bson:"name"`
	DoB         string    `json:"dob" bson:"dob"`
	Age         int       `json:"age" bson:"age"`
	Address     string    `json:"address" bson:"address"`
	Phone       string    `json:"phone" bson:"phone"`
	BloodType   string    `json:"blood_type" bson:"blood_type"`
	IsImpatient int       `json:"is_impatient" bson:"is_impatient"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	ImpatientAt time.Time `json:"impatient_at" bson:"impatient_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}
