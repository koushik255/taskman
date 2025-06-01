package main

import (
	"fmt"
	"log"

	db "awesomeness/db"
)

func main() {
	dbPath:= "./tasks.db"
	err := db.InitDatabase(dbPath)
	if err != nil {
		log.Fatalf("failed to init database: %v",err)
	}
	fmt.Printf("Database initlizaed at : %s\n", dbPath)
	
	// testTask := "Walk dog"
	// testComplete := 1
	// err = db.AddTask(testTask,testComplete)
	// if err != nil {
	// 	fmt.Printf("Error adding task %v ",err)
	// }

	// err = db.DeleteTask(2) 
	// if err != nil {
	// 	fmt.Printf("Error deleteing Task! %v",err)
	// }

	tasks, err := db.GetAllTasks()
    if err != nil {
        log.Printf("Error retrieving tasks: %v", err)
    } else {
        for _, task := range tasks {
            log.Printf("Task ID: %d, Task: %s, Completed: %d, Created At: %s, Updated At: %s",
                task.ID, task.Task, task.Completed, task.CreatedAt, task.UpdatedAt)
        }
    }
}

