package handlers

import (
	"context"

	"github.com/KonyD/go-todo-app/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func UpdateTodo(collection *mongo.Collection) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		objectID, err := bson.ObjectIDFromHex(id)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid todo ID"})
		}

		var todo models.Todo

		filter := bson.M{"_id": objectID}

		if err := collection.FindOne(context.Background(), filter).Decode(&todo); err != nil {
			return nil
		}

		update := bson.M{"$set": bson.M{"completed": !todo.Completed}}

		if _, err := collection.UpdateOne(context.Background(), filter, update); err != nil {
			return nil
		}

		return c.Status(200).JSON(fiber.Map{"success": true})
	}
}
