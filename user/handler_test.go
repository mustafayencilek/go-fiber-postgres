package user

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/mustafayencilek/go-fiber-postgres/db"
	"github.com/stretchr/testify/assert"
)

func TestGetHandler(t *testing.T) {
	database, err := db.Connect()
	assert.Nil(t, err)
	repo := NewRepository(database)
	service := NewService(repo)
	handler := NewHandler(service)

	app := fiber.New()
	app.Get("/users/:id", handler.Get)
	id, err := repo.Create(Model{Name: "ali", Email: "ali@gmail.com"})
	assert.Nil(t, err)
	req := httptest.NewRequest("GET", fmt.Sprintf("/users/%d", id), nil)
	resp, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}
