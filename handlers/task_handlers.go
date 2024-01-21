package handlers

import (
	"ToDo_bot/database"
	"ToDo_bot/entities"
	"fmt"
	"log"
)

// Retrieve and send tasks information
func GetTasks(tg_id int64) string {
	var resp string
	var tasks []entities.Task
	if err := database.DB.Where("TgId = ?", tg_id).Find(&tasks).Error; err != nil {
		log.Printf("ERROR: %v", err.Error())
		return "Error. Code: 0.02"
	}
	for _, t := range tasks {
		resp += fmt.Sprintf("<b>%d</b>. %s\n", t.Id, t.Text)
	}
	if len(resp) == 0 {
		resp = "No saved tasks"
	}
	return resp
}

// Add a new task
func AddTask(tg_id int64, text string) string {
	task := entities.Task{TgId: tg_id, Text: text}
	result := database.DB.Create(&task)
	if result.Error != nil {
		log.Printf("ERROR: %v", result.Error.Error())
		return "Error. Code: 0.03"
	}
	if result.RowsAffected == 0 {
		log.Printf("Warning: No rows affected during task creation")
		return "Error. Code: 0.03"
	}
	return "Your task has been added"
}

// Get a task by ID
func GetTaskByID(tg_id int64, id string) string {
	var task entities.Task
	if err := database.DB.Where("TgId = ? AND Id = ?", tg_id, id).Find(&task).Error; err != nil {
		log.Printf("ERROR: %v", err.Error())
		return "Error. Code: 0.04"
	}
	return fmt.Sprintf("<b>%d</b>. %s\n", task.Id, task.Text)
}

// Delete a task by ID
func DelTaskByID(tg_id int64, id string) string {
	if err := database.DB.Where("TgId = ? AND Id = ?", tg_id, id).Delete(&entities.Task{}).Error; err != nil {
		log.Printf("ERROR: %v", err.Error())
		return "Error. Code: 0.05"
	}

	return "Your task has been deleted"
}

// Update a task by ID with new content
func UpdTask(tg_id int64, id int, newTaskText string) string {
	newTask := entities.Task{Text: newTaskText}
	if err := database.DB.Model(&entities.Task{}).Where("TgId = ? AND Id = ?", tg_id, id).Updates(newTask).Error; err != nil {
		log.Printf("ERROR: %v", err.Error())
		return "Error. Code: 0.06"
	}
	return "Task has been updated"
}

// Clear all tasks
func ClearTasks(tg_id int64) string {
	if err := database.DB.Where("TgId = ?", tg_id).Delete(&entities.Task{}).Error; err != nil {
		log.Printf("ERROR: %v", err.Error())
		return "Error. Code: 0.07"
	}

	return "Your tasks have been deleted"
}
