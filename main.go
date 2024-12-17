package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"timesheet-manager-backend/api/routes"
	"timesheet-manager-backend/pkg/book"
	"timesheet-manager-backend/pkg/task"
	"timesheet-manager-backend/pkg/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	db, cancel, err := databaseConnection()
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	fmt.Println("Database connection success!")

	// API - Books
	bookCollection := db.Collection("books")
	bookRepo := book.NewRepo(bookCollection)
	bookService := book.NewService(bookRepo)

	// API - Users
	userCollection := db.Collection("users")
	userRepo := user.NewRepo(userCollection)
	userService := user.NewService(userRepo)

	// API - Tasks
	taskCollection := db.Collection("tasks")
	taskRepo := task.NewRepo(taskCollection)
	taskService := task.NewService(taskRepo)

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to the timesheet-manager-backend mongo book shop!"))
	})
	api := app.Group("/api")
	routes.BookRouter(api, bookService)
	routes.UserRouter(api, userService)
	routes.TaskRouter(api, taskService)
	defer cancel()
	log.Fatal(app.Listen(getPort()))
}

func databaseConnection() (*mongo.Database, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DB_URI")).SetServerSelectionTimeout(5*time.
		Second))
	if err != nil {
		cancel()
		return nil, nil, err
	}
	db := client.Database(os.Getenv("DB_NAME"))
	return db, cancel, nil
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3001"
	} else {
		port = ":" + port
	}

	return port
}
