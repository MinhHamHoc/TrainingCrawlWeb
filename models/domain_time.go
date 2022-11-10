package models

type Time struct {
	Date   string `bson:"date, date"`
	Domain string `bson:"domain,domain"`
}
