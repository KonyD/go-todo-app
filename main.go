package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/KonyD/go-todo-app/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

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

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin,Content-Type,Accept",
	}))

	app.Get("/api/todos", handlers.GetTodos(collection))
	app.Post("/api/todos", handlers.CreateTodo(collection))
	app.Patch("/api/todos/:id", handlers.UpdateTodo(collection))
	app.Delete("/api/todos/:id", handlers.DeleteTodo(collection))

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}

	log.Fatal(app.Listen(":" + PORT))
}
