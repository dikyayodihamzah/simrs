package domain

type EMR struct {
	ID         string      `json:"id" bson:"_id"`
	Patient    Patient     `json:"patient" bson:"patient"`
	Screenings []Screening `json:"screenings" bson:"screenings"`
}
