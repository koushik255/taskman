package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "A simple CLI application",
	Long:  `Made with Cobra and Go`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from Taskman!")
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

	// Add flags to greet command
	greetCmd.Flags().BoolP("uppercase", "u", false, "Display the greeting in uppercase")
	
	// Add greet command to root command
	rootCmd.AddCommand(greetCmd)
}