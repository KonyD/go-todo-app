package handlers

import (
	"context"

	"github.com/KonyD/go-todo-app/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func CreateTodo(collection *mongo.Collection) fiber.Handler {
	return func(c *fiber.Ctx) error {
		todo := new(models.Todo)

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Todo body cannot be empty"})
		}

		insertResult, err := collection.InsertOne(context.Background(), todo)
		if err != nil {
			return err
		}

		todo.ID = insertResult.InsertedID.(bson.ObjectID)

		return c.Status(201).JSON(todo)
	}
}
