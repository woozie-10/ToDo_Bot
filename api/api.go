package api

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

// Get environment variables for MySQL connection
var MYSQL_DATABASE = os.Getenv("MYSQL_DATABASE")
var MYSQL_USER = os.Getenv("MYSQL_USER")
var MYSQL_PASSWORD = os.Getenv("MYSQL_PASSWORD")

// Define a struct to represent a task
type task struct {
	Id   int
	Task string
}

// InitDB initializes the MySQL database
func InitDB() {
	// Open a MySQL database connection
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(db:3306)/%s", MYSQL_USER, MYSQL_PASSWORD, MYSQL_DATABASE))

	// Create the 'Tasks' table if it doesn't exist
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS Tasks (
            Id INT AUTO_INCREMENT PRIMARY KEY,
            Task MEDIUMTEXT
        )
    `)
	if err != nil {
		panic(err)
	}
}

// GetInfo retrieves task information from the database
func GetInfo() string {
	// Open a MySQL database connection
	var db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(db:3306)/%s", MYSQL_USER, MYSQL_PASSWORD, MYSQL_DATABASE))

	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Execute a SELECT query to retrieve tasks
	rows, err := db.Query("select * from todo_db.Tasks")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	tasks := []task{}

	// Iterate through the result rows and populate the 'tasks' slice
	for rows.Next() {
		t := task{}
		err := rows.Scan(&t.Id, &t.Task)
		if err != nil {
			fmt.Println(err)
			continue
		}
		tasks = append(tasks, t)
	}
	var result string = ""

	// Format the task information as a string
	for _, t := range tasks {
		result += fmt.Sprintf("<b>%d</b>. %s\n", t.Id, t.Task)
	}
	if len(result) == 0 {
		result = "No saved tasks"
	}
	return result
}

// AddInfo adds a new task to the database
func AddInfo(userTask string) string {
	// Open a MySQL database connection
	var db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(db:3306)/%s", MYSQL_USER, MYSQL_PASSWORD, MYSQL_DATABASE))

	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Execute an INSERT query to add a new task
	_, err = db.Exec("insert into todo_db.Tasks (Task) values (?)", userTask)

	if err != nil {
		panic(err)
	}
	return "Your task has been added"
}

// UpdInfo updates an existing task in the database
func UpdInfo(id int, newTask string) string {
	// Open a MySQL database connection
	var db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(db:3306)/%s", MYSQL_USER, MYSQL_PASSWORD, MYSQL_DATABASE))

	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Execute an UPDATE query to update an existing task
	_, err = db.Exec("update todo_db.Tasks set Task = ? where Id = ?", newTask, id)
	if err != nil {
		panic(err)
	}
	return "Task has been updated"
}

// DelInfo deletes a task from the database
func DelInfo(id int) string {
	// Open a MySQL database connection
	var db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(db:3306)/%s", MYSQL_USER, MYSQL_PASSWORD, MYSQL_DATABASE))

	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Execute a DELETE query to delete a task
	_, err = db.Exec("delete from todo_db.Tasks where id = ?", id)
	if err != nil {
		panic(err)
	}
	return "Your task has been deleted"
}

// ClearInfo deletes all tasks from the database
func ClearInfo() string {
	// Open a MySQL database connection
	var db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(db:3306)/%s", MYSQL_USER, MYSQL_PASSWORD, MYSQL_DATABASE))

	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Execute a DELETE query to delete all tasks
	_, err = db.Exec("delete from todo_db.Tasks")
	if err != nil {
		panic(err)
	}
	return "Your tasks have been deleted"
}
