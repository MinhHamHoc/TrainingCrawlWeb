package SaveIntoMongo

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Time struct {
	Year   int    `bson:"year,year"`
	Month  string `bson:"month,month"`
	Day    int    `bson:"day,day"`
	Domain string `bson:"domain,domain`
}

func Saving(pathFile string, year int, month string, day int) {

	uriMongo := viper.GetString("database.URI")

	//Set client options
	clientOptions := options.Client().ApplyURI(uriMongo)

	//Connect to MongoDb
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	nameDatabase := viper.GetString("database.name")
	nameCollection := viper.GetString("database.collection")
	DomainCollection := client.Database(nameDatabase).Collection(nameCollection)

	file, err := os.Open(pathFile)
	if err != nil {
		fmt.Print(err)
	}

	dataScan := bufio.NewScanner(file)

	for dataScan.Scan() {
		line := dataScan.Text()
		dataTime := Time{Year: year, Month: month, Day: day, Domain: line}
		dataMongo, err := DomainCollection.InsertOne(ctx, dataTime)
		if err != nil {
			log.Fatal(err, dataMongo)
		}
	}

}
