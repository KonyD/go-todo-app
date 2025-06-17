package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func DeleteTodo(collection *mongo.Collection) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		objectID, err := bson.ObjectIDFromHex(id)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid todo ID"})
		}

		filter := bson.M{"_id": objectID}
		if _, err = collection.DeleteOne(context.Background(), filter); err != nil {
			return err
		}

		return c.Status(200).JSON(fiber.Map{"success": true})
	}
}
