package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/777777miSSU7777777/go-ass/api"
	"github.com/777777miSSU7777777/go-ass/repository"
	"github.com/777777miSSU7777777/go-ass/service"
)

func main() {
	var connectionString string
	var storageLocation string

	flag.StringVar(&connectionString, "connection_string", "", "DB connection string")
	homePath := os.Getenv("HOME")
	flag.StringVar(&storageLocation, "storage_location", homePath+"/goass/storage", "Storage location")
	flag.Parse()

	repo := repository.NewRepository(connectionString)
	svc := service.New(repo)
	storageManager := api.NewStorageManager(storageLocation)
	apiHandlers := api.NewAPI(svc, storageManager)
	streamAPIHandlers := api.NewStreamAPI(storageManager)

	app := fiber.New(fiber.Config{BodyLimit: 8 * 1024 * 1024 * 1024})
	app.Use(cors.New())

	api.SetupAPIRouter(app, apiHandlers)
	api.SetupAuthRouter(app, apiHandlers)
	api.SetupStreamRouter(app, streamAPIHandlers)

	app.Get("/health-check", func(ctx *fiber.Ctx) error {
		ctx.Status(200).Send([]byte{})
		return nil
	})

	fmt.Println("App started")
	err := app.Listen(":8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
