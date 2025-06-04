package cmd

import (
	"fmt"
	"os"
	"strings"

	blud "awesomeness/funcs"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "A simple CLI application",
	Long:  `Made with Cobra and Go`,
	Run: func(cmd *cobra.Command, args []string) {

		blud.SwitchCase()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Add commands and flags in init()
	
	// Create greet command
	var greetCmd = &cobra.Command{
		Use:   "greet [name]",
		Short: "Greets a person",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := "World"
			if len(args) > 0 {
				name = args[0]
			}
			
			uppercase, _ := cmd.Flags().GetBool("uppercase")
			greeting := fmt.Sprintf("Hello, %s!", name)
			if uppercase {
				greeting = fmt.Sprintf("HELLO, %s!", name)
			}
			
			fmt.Println(greeting)
		},
	}

	var addTaskCmd = &cobra.Command{
		Use: "add [task]",
		Short: "Adds a task!",
		Args: cobra.MaximumNArgs(10),
		Run: func(cmd *cobra.Command, args []string) {
			task := strings.Join(args, " ")
			blud.AddTask(task)
		},
	}

	var showTaskCmd = &cobra.Command{
		Use: "show",
		Short: "Shows the tasks",
		Args: cobra.MaximumNArgs(1),
		Run : func(cmd *cobra.Command, args []string){
			// input := strings.Join(input, "")
			blud.Start()
			blud.ShowTasks()
		},
	}

	var deleteTaskCmd = &cobra.Command{
		Use: "del",
		Short: "Deletes a Task",
		Args: cobra.MaximumNArgs(1),
		Run : func(cmd *cobra.Command, args []string){
			task := strings.Join(args, "")
			blud.Start()
			blud.CompleteTask(task)

		},
	}



	// Add flags to greet command
	greetCmd.Flags().BoolP("uppercase", "u", false, "Display the greeting in uppercase")
	
	// Add greet command to root command
	rootCmd.AddCommand(greetCmd)
	rootCmd.AddCommand(addTaskCmd)
	rootCmd.AddCommand(showTaskCmd)
	rootCmd.AddCommand(deleteTaskCmd)
}