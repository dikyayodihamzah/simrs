package domain

type Disease struct {
	ID               string             `json:"id" bson:"_id"`
	Name             string             `json:"name" bson:"name"`
	Symptoms         []Symptom          `json:"symptoms" bson:"symptoms"`
	ReqMedEquipments []MedicalEquipment `json:"req_med_equipments" bson:"req_med_equipments"`
}
