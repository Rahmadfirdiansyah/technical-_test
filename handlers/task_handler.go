package handlers

import (
	"strconv"

	"stability-test-task-api/models"
	"stability-test-task-api/store"

	"github.com/gofiber/fiber/v2"
)

func GetTasks(c *fiber.Ctx) error {
	tasks := store.GetAllTasks()
	return c.JSON(tasks)
}

func GetTask(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid task ID",
		})
	}

	task := store.GetTaskByID(id)

	if task == nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "task not found",
		})
	}

	return c.JSON(task)
}

func CreateTask(c *fiber.Ctx) error {
	var task models.Task

	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if task.Title == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "title is required",
		})
	}

	store.AddTask(task)

	return c.JSON(task)
}

func DeleteTask(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, _ := strconv.Atoi(idParam)

	deleted := store.DeleteTask(id)

	if !deleted {
		return c.Status(404).JSON(fiber.Map{
			"error": "task not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "deleted",
	})
}
