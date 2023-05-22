package domain

type Room struct {
	RoomNumber    string             `json:"room_number" bson:"room_number"`
	Floor         string             `json:"floor" bson:"floor"`
	Class         string             `json:"class" bson:"class"`
	Capacity      string             `json:"capacity" bson:"capacity"`
	Cluster       string             `json:"cluster" bson:"cluster"`
	MedEquipments []MedicalEquipment `json:"med_equipments" bson:"med_equipments"`
}
