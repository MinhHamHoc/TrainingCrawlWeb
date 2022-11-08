package models

type Time struct {
	Year   int    `bson:"year,year"`
	Month  string `bson:"month,month"`
	Day    int    `bson:"day,day"`
	Domain string `bson:"domain,domain`
}
