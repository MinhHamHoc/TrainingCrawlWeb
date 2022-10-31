package SaveIntoMongo

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Saving(pathFile string, year int, month string, day int) {

	//Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

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

	DomainCollection := client.Database("mydb").Collection("domain")

	file, err := os.Open(pathFile)
	if err != nil {
		fmt.Print(err)
	}

	dataScan := bufio.NewScanner(file)

	for dataScan.Scan() {
		line := dataScan.Text()
		newBson := bson.D{{Key: "Year", Value: year}, {Key: "Month", Value: month}, {Key: "Day", Value: day}, {Key: "Domain", Value: line}}
		dataMongo, err := DomainCollection.InsertOne(ctx, newBson)
		if err != nil {
			log.Fatal(err, dataMongo)
		}
	}

}
