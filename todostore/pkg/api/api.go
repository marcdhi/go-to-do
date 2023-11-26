package api

import (
	"gotodo/todostore/pkg/models"
	"gotodo/todostore/pkg/repo"
	"gotodo/todostore/pkg/service"

	"github.com/gofiber/fiber/v2"
)

type API struct {

	ToDoService *service.TodoService

}

func SetupApp() *fiber.App {
	app := fiber.New()

	api := &API{
		ToDoService: &service.TodoService{
			TodoRepository: &repo.TodoRepo{},
		},
	}

	v1 := app.Group("/api/v1")
	{
		v1.Get("/todos", api.getTodos)
		v1.Post("/todos", api.addTodos)
	}

	return app
}

func (api *API) getTodos(c *fiber.Ctx) error {
	todos := api.ToDoService.GetAll()
	return c.JSON(todos)
}

func (api *API) addTodos(c *fiber.Ctx) error {
	var todo models.Todo

	if err := c.BodyParser(&todo); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	api.ToDoService.Add(todo)
	return c.SendStatus(201)
}