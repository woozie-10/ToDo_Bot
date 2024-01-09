package handlers

import (
	"ToDo_bot/database"
	"ToDo_bot/entities"
	"fmt"
	"log"
)

// Retrieve and send tasks information
func GetTasks(tg_id int64) string {
	rows, err := database.DB.Query("select * from todo_db.Tasks where TgId = ?", tg_id)
	if err != nil {
		log.Printf("ERROR: %v", err.Error())
		return "Error. Code: 0.01"
	}
	defer rows.Close()
	var tasks []entities.Task

	for rows.Next() {
		t := entities.Task{}
		err := rows.Scan(&t.Id, &t.TgId, &t.Text)
		if err != nil {
			log.Printf("ERROR: %v", err.Error())
			return "Error. Code: 0.02"
		}
		tasks = append(tasks, t)
	}
	var result string = ""

	for _, t := range tasks {
		result += fmt.Sprintf("<b>%d</b>. %s\n", t.Id, t.Text)
	}
	if len(result) == 0 {
		result = "No saved tasks"
	}
	return result
}

// Add a new task
func AddTask(tg_id int64, text string) string {
	_, err := database.DB.Exec("insert into todo_db.Tasks (TgId, Text) values (?, ?)", tg_id, text)
	if err != nil {
		log.Printf("ERROR: %v", err.Error())
		return "Error. Code: 0.03"
	}
	return "Your task has been added"
}

// Get a task by ID
func GetTaskByID(tg_id int64, id string) string {
	row := database.DB.QueryRow("select * from todo_db.Tasks where TgId = ? and Id = ?", tg_id, id)
	var task entities.Task
	err := row.Scan(&task.Id, &task.TgId, &task.Text)
	if err != nil {
		log.Printf("ERROR: %v", err.Error())
		return "Error. Code: 0.04"
	}
	return fmt.Sprintf("<b>%d</b>. %s\n", task.Id, task.Text)
}

// Delete a task by ID
func DelTaskByID(tg_id int64, id string) string {
	_, err := database.DB.Exec("delete from todo_db.Tasks where TgId = ? and Id = ?", tg_id, id)
	if err != nil {
		log.Printf("ERROR: %v", err.Error())
		return "Error. Code: 0.05"
	}
	return "Your task has been deleted"
}

// Update a task by ID with new content
func UpdTask(tg_id int64, id, newTask string) string {
	_, err := database.DB.Exec("update todo_db.Tasks set Text = ? where TgId = ? and Id = ?", newTask, tg_id, id)
	if err != nil {
		log.Printf("ERROR: %v", err.Error())
		return "Error. Code: 0.06"
	}
	return "Task has been updated"
}

// Clear all tasks
func ClearTasks(tg_id int64) string {
	_, err := database.DB.Exec("delete from todo_db.Tasks where TgId = ?", tg_id)
	if err != nil {
		log.Printf("ERROR: %v", err.Error())
		return "Error. Code: 0.07"
	}
	return "Your tasks have been deleted"
}
