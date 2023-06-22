package main

import (
	"context"
	"log"
	"time"

	"github.com/cleverextechnology/site-generator/src"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg MongoInstance

const dbName = "fiber-UserPortfolioData"
const mongoURI = "mongodb://localhost:27017/" + dbName

func Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		return err
	}

	mg = MongoInstance{
		Client: client,
		Db:     db,
	}
	return nil
}

func main() {

	if err := Connect(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	//GET all user

	app.Get("/user", func(c *fiber.Ctx) error {

		query := bson.D{{}}

		cursor, err := mg.Db.Collection("PortfolioData").Find(c.Context(), query)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		var users []src.Data = make([]src.Data, 0)

		if err := cursor.All(c.Context(), &users); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		//fmt.Print(users)
		return c.JSON(users)
	})

	//GET single user

	app.Get("/user/:id", func(c *fiber.Ctx) error {

		idParam := c.Params("id")

		userID, err := primitive.ObjectIDFromHex(idParam)

		if err != nil {
			log.Fatal(err)
		}

		query := bson.D{{Key: "_id", Value: userID}}

		var user = new(src.Data)

		err = mg.Db.Collection("PortfolioData").FindOne(c.Context(), query).Decode(user)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		// Parser and Hydrate
		tpl, err := src.ParseTemplate()
		if err != nil {
			return c.Status(500).JSON(err.Error())
		}
		if err := src.HydrateTemplate(tpl, *user); err != nil {
			return c.Status(500).JSON(err.Error())
		}

		return c.Status(200).JSON(user)

	})

	// INSERT user

	app.Post("/user", func(c *fiber.Ctx) error {
		collection := mg.Db.Collection("PortfolioData")

		user := new(src.Data)

		if err := c.BodyParser(user); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		insertionResult, err := collection.InsertOne(c.Context(), user)

		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
		createdRecord := collection.FindOne(c.Context(), filter)

		createdUser := &src.Data{}
		createdRecord.Decode(createdUser)

		return c.Status(201).JSON(createdUser)

	})

	// UPDATE user based on ID

	app.Put("/user/:id", func(c *fiber.Ctx) error {
		idParam := c.Params("id")

		userID, err := primitive.ObjectIDFromHex(idParam)

		if err != nil {
			return c.SendStatus(400)
		}

		user := new(src.Data)

		if err := c.BodyParser(user); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		query := bson.D{{Key: "_id", Value: userID}}
		update := bson.D{
			{Key: "$set",
				Value: bson.D{
					{Key: "pageTitle", Value: user.PageTitle},
					{Key: "headerTitle", Value: user.HeaderTitle},
				},
			},
		}

		err = mg.Db.Collection("PortfolioData").FindOneAndUpdate(c.Context(), query, update).Err()

		if err != nil {
			if err == mongo.ErrNoDocuments {
				return c.SendStatus(400)
			}
			return c.SendStatus(500)
		}

		user.ID = userID

		return c.Status(200).JSON(user)

	})

	//DELETE user based on ID

	app.Delete("/user/:id", func(c *fiber.Ctx) error {

		userID, err := primitive.ObjectIDFromHex(c.Params("id"))

		if err != nil {
			return c.SendStatus(400)
		}

		query := bson.D{{Key: "_id", Value: userID}}
		result, err := mg.Db.Collection("PortfolioData").DeleteOne(c.Context(), &query)

		if err != nil {
			return c.SendStatus(500)
		}

		if result.DeletedCount < 1 {
			return c.SendStatus(404)
		}

		return c.Status(200).JSON("record deleted")
		//return c.JSON(result)

	})
	///////////////////////////////////////////////////////////////////////////////////////////////////////////

	log.Fatal(app.Listen(":3000"))

}
