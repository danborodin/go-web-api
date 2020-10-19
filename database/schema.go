package database

import (
	"database/sql"
	"errors"
	"log"
	"todoapi/model"
)

//GetAllTasks get all tasks from dt
func GetAllTasks(db *sql.DB) []model.Task {
	var tasks []model.Task

	row, err := db.Query("SELECT * FROM tasks ORDER BY Date")
	if err != nil {
		log.Println("Error while selecting all the tasks from db", err)
	}
	defer row.Close()

	for row.Next() {
		var id int
		var Name string
		var Details string
		var Date string
		var Done int
		var _done bool
		row.Scan(&id, &Name, &Details, &Date, &Done)
		if Done == 0 {
			_done = false
		} else {
			_done = true
		}
		tasks = append(tasks, model.Task{ID: id, Name: Name, Details: Details, Date: Date, Done: _done})
	}

	return tasks
}

//InsertNewTask add new task in database
func InsertNewTask(db *sql.DB, task model.Task) error {
	_, err := db.Exec("INSERT INTO tasks (Name, Details, Done, Date) VALUES($1, $2, 0, $3)",
		task.Name, task.Details, task.Date)

	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Task added")
	}

	return err
}

//UpdateTaskByID update a task by id in db
func UpdateTaskByID(db *sql.DB, id int, newTask model.Task) error {

	task := model.Task{}

	row, err := db.Query("SELECT * FROM tasks WHERE id = $1", id)
	if err != nil {
		log.Println("Error while selecting a task by id", err)
	}
	defer row.Close()

	for row.Next() {
		var (
			Name    string
			Details string
			Date    string
			Done    bool
		)
		row.Scan(&id, &Name, &Details, &Date, &Done)
		task.Name = Name
		task.Details = Details
		task.Done = Done
	}

	if newTask.Name == "" {
		newTask.Name = task.Name
	}
	if newTask.Details == "" {
		newTask.Details = task.Details
	}

	_, err = db.Exec("UPDATE tasks SET Name = $1, Details = $2, Done = $3 WHERE id = $4",
		newTask.Name, newTask.Details, newTask.Done, id)

	if err != nil {
		log.Println("Error while updating task")
	} else {
		log.Printf("Task with id %d updated", id)
	}
	return err
}

//DeleteTaskByID delete a task by id :)
func DeleteTaskByID(db *sql.DB, id int) error {
	var _id int
	var err error
	row, err := db.Query("SELECT id FROM tasks WHERE id = $1", id)
	if err != nil {
		log.Println("Error while geting id", err)
	}

	defer row.Close()

	for row.Next() {
		row.Scan(&_id)
	}
	if _id == id {
		log.Printf("Task with id %d deleted", id)
		_, err = db.Exec("DELETE FROM tasks WHERE id = $1", id)
		if err != nil {
			log.Println("Error deleting a task by id", err)
		}
	} else {
		return errors.New("Id not found in db")
	}

	return err
}
