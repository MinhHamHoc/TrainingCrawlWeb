package configs

import (
	"Crawl_Web/models"
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var DomainCollection *mongo.Collection = GetCollection(DatabaseConfig, "domain3")

func Saving(pathFile, date string) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	file, err := os.Open(pathFile)
	if err != nil {
		fmt.Print(err)
	}

	dataScan := bufio.NewScanner(file)

	for dataScan.Scan() {
		line := dataScan.Text()
		dataTime := models.Time{Date: date, Domain: line}
		dataMongo, err := DomainCollection.InsertOne(ctx, dataTime)
		if err != nil {
			log.Fatal(err, dataMongo)
		}
	}

}
