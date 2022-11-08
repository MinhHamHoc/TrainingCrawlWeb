package controllers

import (
	"Crawl_Web/configs"
	"Crawl_Web/models"
	"Crawl_Web/responses"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var domainCollection *mongo.Collection = configs.GetCollection(configs.DatabaseConfig, "domain3")

func GetAllDomain() []primitive.M {
	cur, err := domainCollection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatalln(err)
	}
	var results []primitive.M

	for cur.Next(context.Background()) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatalln(err)
		}
		results = append(results, result)
	}
	defer cur.Close(context.Background())
	return results
}

func ReadDomain(c *fiber.Ctx) error {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var allDomain = GetAllDomain()
	defer cancel()

	return c.Status(http.StatusOK).JSON(responses.DomainResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": allDomain}})
}

func GetADomain(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	domainId := c.Params("domainId")
	var domain models.Domains
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(domainId)

	err := domainCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&domain)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.DomainResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.DomainResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": domain}})
}

func CreateDomain(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var domain models.Domains
	defer cancel()

	newDomain := models.Domains{
		Id:     primitive.NewObjectID(),
		Year:   domain.Year,
		Month:  domain.Month,
		Day:    domain.Day,
		Domain: domain.Domain,
	}

	result, err := domainCollection.InsertOne(ctx, newDomain)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.DomainResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.DomainResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}