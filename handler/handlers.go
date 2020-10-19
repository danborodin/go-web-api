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
		c.Response
	}
	return err
}

//AddTask add new task
func AddTask(c *fiber.Ctx) error {
	task := new(model.Task)

	err := c.BodyParser(task)
	if err != nil {
		log.Fatal(err)
	}

	database.InsertNewTask(database.DB, *task)

	return err
}

//UpdateTask update a task
func UpdateTask(c *fiber.Ctx) error {
	task := new(model.Task)

	err := c.BodyParser(task)
	if err != nil {
		log.Fatal(err)
	}

	id, err := strconv.Atoi(c.Params("id"))

	database.UpdateTaskByID(database.DB, id, *task)

	return err
}

//DeleteTask delete a task
func DeleteTask(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Fatal("Error converting string to integer")
	}

	database.DeleteTaskByID(database.DB, id)

	return err
}
