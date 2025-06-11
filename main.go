package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Todo struct {
	ID        bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Completed bool          `json:"completed"`
	Body      string        `json:"body"`
}

var collection *mongo.Collection

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	MONGODB_URI := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	if err = client.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB Atlas")

	collection = client.Database("golang_db").Collection("todos")

	app := fiber.New()

	app.Get("/api/todos", getTodos)
	app.Post("/api/todos", createTodo)
	app.Patch("/api/todos/:id", updateTodo)
	app.Delete("/api/todos/:id", deleteTodo)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}

	log.Fatal(app.Listen(":" + PORT))
}

func getTodos(c *fiber.Ctx) error {
	var todos []Todo

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo Todo
		if err := cursor.Decode(&todo); err != nil {
			return err
		}
		todos = append(todos, todo)
	}

	return c.JSON(todos)
}

func createTodo(c *fiber.Ctx) error {
	todo := new(Todo)

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

func updateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid todo ID"})
	}

	var todo Todo

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

func deleteTodo(c *fiber.Ctx) error {
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
