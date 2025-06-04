package blud

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	db "awesomeness/db"
)

func showTasks(){
	tasks, err := db.GetAllTasks()
    if err != nil {
        log.Printf("Error retrieving tasks: %v", err)
    } else {
        for _, task := range tasks {
            // log.Printf("Task ID: %d, Task: %s, Completed: %d, Created At: %s, Updated At: %s",
            //     task.ID, task.Task, task.Completed, task.CreatedAt, task.UpdatedAt)
            // fmt.Printf("taaskID: %d, Task: %s, Completed: %d, CreatedAt :%s, UpdatedAt :%s",task.ID, task.Task, task.Completed, task.CreatedAt, task.UpdatedAt)
            fmt.Println("TaskID:",task.ID,"Task:",task.Task,"Completed: ",task.Completed, "CreatedAt:",task.Completed,"UpdatedAt:",task.UpdatedAt)
            fmt.Println("-----------------------------")
    }
}
}

func showOptions() {
		fmt.Println("Select a option")
    	fmt.Println("Select option 1 to create a new Task!")
    	fmt.Println("Select option 2 to Complete a Task!")
    	fmt.Println("Select option 3 to show all Tasks")
    	fmt.Println("Type 'exit' to exit the program!")
}

func AddTask(task string){
	Start()
	// fmt.Println("add")
    // 	reader := bufio.NewReader(os.Stdin)
    // fmt.Println("Type your Task!")

    // input, err := reader.ReadString('\n')
    // if err != nil {
    // 	fmt.Println("Error reading input: ",err)
    // 	return
    // }
	testComplete := 1
	err := db.AddTask(task,testComplete)
	if err != nil {
		fmt.Printf("Error adding task %v ",err)
	}
}

func AddTaskOriginal(){
	Start()
	fmt.Println("add")
    	reader := bufio.NewReader(os.Stdin)
    fmt.Println("Type your Task!")

    input, err := reader.ReadString('\n')
    if err != nil {
    	fmt.Println("Error reading input: ",err)
    	return
    }
	testComplete := 1
	err = db.AddTask(input,testComplete)
	if err != nil {
		fmt.Printf("Error adding task %v ",err)
	}
}



func SwitchCase() {
	Start()

	 reader := bufio.NewReader(os.Stdin)
    for {
   		showOptions()

    input, err := reader.ReadString('\n')
    if err != nil {
    	fmt.Println("Error reading input: ",err)
    	return
    }

    input = strings.TrimSpace(input)

	switch input {
    case "1":
    	AddTaskOriginal()

   	case "2":
   		fmt.Println("complete tasks")
   		showTasks()
   		reader := bufio.NewReader(os.Stdin)
    fmt.Println("Type your Task to Complete!")

    input, err := reader.ReadString('\n')
    if err != nil {
    	fmt.Println("Error reading input: ",err)
    	return
    }
    input = strings.TrimSpace(input)
    taskToDelete,err := strconv.Atoi(input)
    if err != nil {
    	fmt.Println("error converting string to int",err)
    }
	err = db.DeleteTask(taskToDelete)
	if err != nil {
		fmt.Println("-------------------------")
		fmt.Printf("Error deleting task %v\n",err)
		fmt.Println("-------------------------")

	}

	showTasks()

   	case "3":
   		fmt.Println("Show tasks")
   		showTasks()
   	case "exit":
   		fmt.Println("Exiting program")
   		os.Exit(0)

   	default :
   		fmt.Println("you chose none")
    }

 }
}






func Start() {


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


	// tasks, err := db.GetAllTasks()
    // if err != nil {
    //     log.Printf("Error retrieving tasks: %v", err)
    // } else {
    //     for _, task := range tasks {
    //         log.Printf("Task ID: %d, Task: %s, Completed: %d, Created At: %s, Updated At: %s",
    //             task.ID, task.Task, task.Completed, task.CreatedAt, task.UpdatedAt)
    //     }
    // }

    fmt.Println("------WELCOME TO TASK MANAGER-------")
    
 

}

