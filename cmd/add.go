package cmd

import (
	"fmt"
	"strings"

	"github.com/Basics/src/github.com/TinStay/Task/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		// Get tasks name from input
		task :=	strings.Join(args, " ")

		// Create task
		_, err := db.CreateTask(task)

		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
			return
		}
		
		fmt.Printf("Added \"%s\" to your task list.\n", task)
	},
}

func init(){
	RootCmd.AddCommand(addCmd)
}