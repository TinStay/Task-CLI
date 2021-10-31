package cmd

import (
	"fmt"
	"os"

	"github.com/Basics/src/github.com/TinStay/Task/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		// Get all tasks
		tasks, err := db.GetTasks()

		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
			os.Exit(1)
		}

		// No tasks are stored in the database
		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete")
			return
		}

		// // Print all tasks
		// fmt.Println("You have the following tasks: ")
		// for i, task := range tasks {
		// 	fmt.Printf("%d. %s \n", i+1 ,task.Value)
		// }
		
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
