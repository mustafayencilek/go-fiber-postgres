package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/mustafayencilek/go-fiber-postgres/db"
	"github.com/mustafayencilek/go-fiber-postgres/user"
)

func main() {
	database, err := db.Connect()

	if err != nil {
		log.Fatalf("cannot connected to db  %v", err)
	}
	fmt.Println(database.Config)

	repo := user.NewRepository(database)
	err = repo.Migration()
	if err != nil {
		log.Fatal(err)
	}

	service := user.NewService(repo)
	handler := user.NewHandler(service)
	app := fiber.New()
	app.Get("/users/:id", handler.Get)

	app.Post("/users", handler.Create)

	app.Listen(":8000")

}
