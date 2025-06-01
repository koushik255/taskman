package db

// neeed to open database add tasks, delete taks, edit taks and get tasks

import (
	"database/sql"
	"time"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Task struct {
	ID     int      
	Task   string
	Completed int
	CreatedAt time.Time
	UpdatedAt time.Time
}

var db *sql.DB

func InitDatabase(dbPath string) error {
	var err error 
	db, err = sql.Open("sqlite3",dbPath)
	if err != nil {
		return fmt.Errorf("failed to open database: %v",err)
	}

	if err = db.Ping(); err !=nil {
		return fmt.Errorf("Failed to ping database: %v",err)
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task TEXT UNIQUE NOT NULL,
		completed INTEGER NOT NULL DEFAULT 0,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);`

	_,err = db.Exec(createTableSQL)
	if err != nil {
		return fmt.Errorf("failed to create table: %v",err)
	}

	log.Println("database initizalzed succesfully")
	return nil
}

func SaveTask(task string, completed int) error {
	if db == nil {
		return fmt.Errorf("datanase not inited")
	}

	updateSQL := `UPDATE tasks SET completed = ?, updated_at = CURRENT_TIMESTAMP WHERE task = ?`
	result, err := db.Exec(updateSQL,task,completed)
	if err != nil {
		return fmt.Errorf("Failed to update tasK :%v",err)
	}

	rowsAffected , err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v",err)

	}

	if rowsAffected == 0 {
		insertSQL := `INSERT INTO tasks (task,completed) VALUES (?,?)`
		_,err = db.Exec(insertSQL,task,completed)
		if err != nil {
			return fmt.Errorf("Failed to insert task: %v",err)
		}
		log.Printf("Inserted new task %s with completed status of %d",task,completed)
	} else {
		log.Printf("Updated task %s with completed status of %d",task,completed )
	}
	return nil
}

func DeleteTask(taskID int) error {
	if db == nil {
		return fmt.Errorf("database not init")
	}

	deleteSQL := `DELETE FROM tasks WHERE id = ?`
	result, err := db.Exec(deleteSQL, taskID)
	if err != nil {
		return fmt.Errorf("failed to delete task: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no task found with id %d", taskID)
	}

	log.Printf("Deleted task with id %d", taskID)
	return nil
}



func AddTask(task string,completed int) error {
	err := SaveTask(task,completed)
	if err != nil {
		return fmt.Errorf("Error adding task to db %v",err)
	}
	return nil
}

func GetAllTasks() ([]Task, error) {
    if db == nil {
        return nil, fmt.Errorf("database not initialized")
    }

    selectSQL := `SELECT id, task, completed, created_at, updated_at FROM tasks`
    rows, err := db.Query(selectSQL)
    if err != nil {
        return nil, fmt.Errorf("failed to query tasks: %v", err)
    }
    defer rows.Close() // Ensure rows are closed after processing

    var tasks []Task
    for rows.Next() {
        var task Task
        err := rows.Scan(&task.ID, &task.Task, &task.Completed, &task.CreatedAt, &task.UpdatedAt)
        if err != nil {
            return nil, fmt.Errorf("failed to scan task: %v", err)
        }
        tasks = append(tasks, task)
    }

    // Check for errors during iteration
    if err = rows.Err(); err != nil {
        return nil, fmt.Errorf("error occurred during rows iteration: %v", err)
    }

    return tasks, nil
}


