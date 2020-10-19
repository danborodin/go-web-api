package handler

import (
	"log"
	"strconv"
	"todoapi/database"
	"todoapi/model"

	"github.com/gofiber/fiber/v2"
)

//GetAllTasks convert all tasks from database in json and return them
func GetAllTasks(c *fiber.Ctx) error {
	var tasks []model.Task
	tasks = database.GetAllTasks(database.DB)
	err := c.JSON(tasks)
	if err != nil {
		log.Println("Error while convert tasks to json", err)
		c.Status(500).SendString("Internal server error")
	}
	return err
}

//AddTask add new task
func AddTask(c *fiber.Ctx) error {
	task := new(model.Task)

	err := c.BodyParser(task)
	if err != nil {
		log.Println("Error while parsing json", err)
		return c.Status(400).SendString("Client bad request")
	}

	err = database.InsertNewTask(database.DB, *task)
	if err != nil {
		log.Println("Error while adding new task", err)
		return c.Status(500).SendString("Internal server error")
	}
	c.SendString("Task added")

	return err
}

//UpdateTask update a task
func UpdateTask(c *fiber.Ctx) error {
	task := new(model.Task)

	err := c.BodyParser(task)
	if err != nil {
		log.Println(err)
		return c.Status(400).SendString("Client bad request")
	}

	id, err := strconv.Atoi(c.Params("id"))

	err = database.UpdateTaskByID(database.DB, id, *task)
	if err != nil {
		log.Println("Error while trying to update an task with unexistent id")
		return c.Status(500).SendString("Id not found")
	}
	c.SendString("Task updated")

	return err
}

//DeleteTask delete a task
func DeleteTask(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println("Error converting string to integer")
		c.Status(400).SendString("Client bad request")
	}

	err = database.DeleteTaskByID(database.DB, id)
	if err != nil {
		c.SendStatus(400)
		return c.SendString("id not found")
	}
	c.SendString("Task deleted")

	return err
}
