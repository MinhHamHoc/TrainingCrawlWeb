package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Domains struct {
	Id     primitive.ObjectID `json:"id,omitempty"`
	Year   int                `json:"Year,omitempty`
	Month  string             `json:"Month,omitempty"`
	Day    int                `json:"Day,omitempty"`
	Domain string             `json:"Domain,omitempty"`
}
