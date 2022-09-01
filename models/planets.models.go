package models

type Planets struct {
	ID      interface{} `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string      `json:"name" bson:"name"`
	Climate string      `json:"climate" bson:"climate"`
	Terrain string      `json:"terrain" bson:"terrain"`
}
