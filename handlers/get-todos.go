package handlers

import (
	"context"

	"github.com/KonyD/go-todo-app/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func GetTodos(collection *mongo.Collection) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var todos []models.Todo
		cursor, err := collection.Find(context.Background(), bson.M{})

		if err != nil {
			return nil
		}

		defer cursor.Close(context.Background())

		for cursor.Next(context.Background()) {
			var todo models.Todo

			if err := cursor.Decode(&todo); err != nil {
				return nil
			}

			todos = append(todos, todo)
		}

		return c.JSON(todos)
	}
}
