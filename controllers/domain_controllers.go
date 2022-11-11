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
	cur, err := domainCollection.Find(context.Background(), bson.M{})
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

func ReadAllDomain(c *fiber.Ctx) error {
	var allDomain = GetAllDomain()

	return c.Status(http.StatusOK).JSON(responses.DomainResponse{
		Status:  http.StatusCreated,
		Message: "success",
		Data:    &fiber.Map{"data": allDomain}},
	)
}

func GetADomain(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	domainId := c.Params("domainId")
	var domain models.Domains
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(domainId)

	err := domainCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&domain)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.DomainResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()}},
		)
	}

	return c.Status(http.StatusOK).JSON(responses.DomainResponse{
		Status:  http.StatusCreated,
		Message: "success",
		Data:    &fiber.Map{"data": domain}},
	)
}

func CreateADomain(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var domain models.Domains
	defer cancel()

	if err := c.BodyParser(&domain); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.DomainResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"dataParse": err.Error()}},
		)
	}

	newDomain := models.Domains{
		Date:   domain.Date,
		Domain: domain.Domain,
	}

	result, err := domainCollection.InsertOne(ctx, newDomain)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.DomainResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"dataMongo": err.Error()}},
		)
	}

	return c.Status(http.StatusCreated).JSON(responses.DomainResponse{
		Status:  http.StatusCreated,
		Message: "success",
		Data:    &fiber.Map{"data": result}},
	)
}

func DeleteADomain(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	domainId := c.Params("domainId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(domainId)

	result, err := domainCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.DomainResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()}},
		)
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(responses.DomainResponse{
			Status:  http.StatusNotFound,
			Message: "error",
			Data: &fiber.Map{
				"data": "Domain ID not found !",
			},
		})
	}

	return c.Status(http.StatusOK).JSON(responses.DomainResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    &fiber.Map{"data": "Domain is deleted"},
	})
}

func UpdateADomain(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	domainId := c.Params("domainId")
	var domain models.Domains
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(domainId)

	if err := c.BodyParser(&domain); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.DomainResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()}},
		)
	}

	update := bson.M{
		"date":   domain.Date,
		"domain": domain.Domain,
	}

	result, err := domainCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.DomainResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()}},
		)
	}

	var updateDomain models.Domains
	if result.MatchedCount == 1 {
		err := domainCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updateDomain)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.DomainResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()}},
			)
		}
	}
	return c.Status(http.StatusOK).JSON(responses.DomainResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    &fiber.Map{"data": updateDomain},
	})
}
