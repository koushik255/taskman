package blud

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	db "awesomeness/db"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"


)



func completeCheck(x int) string {
	if x == 1 {
		return "no"
	} else {
		return ""
	}
}

func ShowTasks(){
	tasks, err := db.GetAllTasks()
    if err != nil {
        log.Printf("Error retrieving tasks: %v", err)
    } else {

    	tableData := pterm.TableData{
            		{"TaskID", "Task","Completed","CreatedAt","UpdatedAt"},
            }
        for _, task := range tasks {
            // log.Printf("Task ID: %d, Task: %s, Completed: %d, Created At: %s, Updated At: %s",
            //     task.ID, task.Task, task.Completed, task.CreatedAt, task.UpdatedAt)
            // fmt.Printf("taaskID: %d, Task: %s, Completed: %d, CreatedAt :%s, UpdatedAt :%s",task.ID, task.Task, task.Completed, task.CreatedAt, task.UpdatedAt)
            // fmt.Println("TaskID:",task.ID,"Task:",task.Task,"Completed: ",task.Completed, "CreatedAt:",task.Completed,"UpdatedAt:",task.UpdatedAt)
            // fmt.Println("-----------------------------")
            // its making a whole new table every file because the data is being looped over,
            // i need a way to collect the data so that il be able to append it to 
            // the table?


            timeFormat := "2006-01-02 15:04:05"

            
            tableData = append(tableData, []string{
            	strconv.Itoa(task.ID),task.Task,completeCheck(task.Completed),task.CreatedAt.Format(timeFormat),task.UpdatedAt.Format(timeFormat),
            	 })


            // err := pterm.DefaultTable.WithHasHeader().WithData(tableData).Render()
            // if err != nil {
            // 	fmt.Println(err)
            // 	return
            // }
            
            }

            // theStyle := pterm.NewStyle(pterm.BgDarkGray)
            
            // unicode := '\u035e'
            // stringString := string(unicode)

            lilMin := "⬌"
            // ⬌⬌⬌⬌⬌⬌⬌⬌⬌⬌⬌⬌⬌⬌⬌⬌⬌⬌⬌⬌⬌⬌⬌⬌⬌⬌⬌⬌⬌⬌⬌⬌⬌⬌⬌⬌⬌⬌⬌
           



            err := pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(tableData).WithRowSeparator(lilMin).Render()
            if err != nil {
            	fmt.Println(err)
            	return

    }

}

}

func print_Confirm(x string) {
	style := pterm.NewStyle(pterm.BgGray, pterm.FgLightWhite, pterm.Bold)
	style.Printfln("%s",x)
} 


func Take_input() (string,error) {
	// reader := bufio.NewReader(os.Stdin)
	result, _ := pterm.DefaultInteractiveTextInput.Show()


	pterm.Info.Printfln("Your input: %s",result)

    // input, err := reader.ReadString('\n')
    // if err != nil {
    // 	fmt.Println("Error reading input: ",err)
    // 	return "",err
    // }
    input := strings.TrimSpace(result)
    return input, nil
}


func showOptions() {
		// fmt.Println("Select a option")
    	// fmt.Println("Select option `add` to create a new Task!")
    	// fmt.Println("Select option `del` to Complete a Task!")
    	// fmt.Println("Select option `show` to show all Tasks")
    	// fmt.Println("Type 'exit' to exit the program")

    	pterm.Info.Println("Type 'add' to create a new Task!\nType 'del' to delete a Task\nType 'show' to show all Tasks!\nType 'exit' to quit!")


}

func AddTask(task string){
	Start()
	// fmt.Println("add")
    // 	reader := bufio.NewReader(os.Stdin)
    fmt.Println("Type your Task!")

    // input, err := reader.ReadString('\n')
    // if err != nil {
    // 	fmt.Println("Error reading input: ",err)
    // 	return
    // }
    fmt.Println("Type your Task!")
	testComplete := 1
	err := db.AddTask(task,testComplete)
	if err != nil {
		fmt.Printf("Error adding task %v ",err)
	}
}

func AddTaskOriginal(){
	Start()
	fmt.Println("add")
	fmt.Println("Type your Task!")
    	input,err := Take_input()
    	if err != nil {
    		fmt.Println("Error taking input!")
    		}
    	
	testComplete := 1
	err = db.AddTask(input,testComplete)
	if err != nil {
		fmt.Printf("Error adding task %v ",err)
	}
}

func CompleteTask(task string) {
	fmt.Println("Type your Task to Complete!")
   	ShowTasks()
   	// 	reader := bufio.NewReader(os.Stdin)
    // fmt.Println("Type your Task to Complete!")

    // input, err := reader.ReadString('\n')
    // if err != nil {
    // 	fmt.Println("Error reading input: ",err)
    // 	return
    // }

    task = strings.TrimSpace(task)
    taskToDelete,err := strconv.Atoi(task)
    if err != nil {
    	fmt.Println("error converting string to int",err)
    }
	err = db.DeleteTask(taskToDelete)
	if err != nil {
		fmt.Println("-------------------------")
		fmt.Printf("Error deleting task %v\n",err)
		fmt.Println("-------------------------")
	}
	
	ShowTasks()
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
    case "add":
    	AddTaskOriginal()

   	case "del":
   		print_Confirm("Delete Tasks!")
   		print_Confirm("Type your Task to Complete!")


   		ShowTasks()
   		///
   		input,err := Take_input()
   		if err != nil {
   			fmt.Println("Erorring taking input")
   		}
    //
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


	ShowTasks()

   	case "show":
   		
   		print_Confirm("Showing Tasks!")
   		// fmt.Println("-------------------------")
   		// style.Println("Showing Tasks!")
   		// fmt.Println("Showing Tasks")
   		// fmt.Println("-------------------------")
   		ShowTasks()
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
	// fmt.Printf("Database initlizaed at : %s\n", dbPath)
	pterm.Info.Printf("Database initlizaed at: %s\n", dbPath)

	text := "Task Man"

	letters := putils.LettersFromString(text)

	pterm.DefaultBigText.WithLetters(letters).Render()
		
	

    // fmt.Println("------WELCOME TO TASK MANAGER-------")
    // pterm.DefaultHeader.Println("WELCOME TO TASK MAN!")
    pterm.DefaultHeader.WithMargin(20).WithBackgroundStyle(pterm.NewStyle(pterm.BgRed)).WithTextStyle(pterm.NewStyle(pterm.FgBlack)).Println("WELCOME TO TASK MANAGER")
	
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

    
 

}

