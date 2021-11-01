package cmd

import (
	"fmt"
	"strconv"

	"github.com/Basics/src/github.com/TinStay/Task/db"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as completed",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args{
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("Invalid argument: %s", arg)
				return
			}else {
				ids = append(ids, id)
			}
		}
		// Fetch tasks from database
		tasks, err := db.GetTasks()

		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}

		for _, id := range ids {
			// Check if an invalid task is typed
			if id <= 0 || id > len(tasks){
				fmt.Println("Invalid task number: ", id)
				continue
			}

			// Get task from input
			task := tasks[id-1]

			// Delete task
			err := db.DeleteTask(task.Key)

			// Print message to user
			if err != nil {
				fmt.Printf("Failed to mark  \"%d\" as completed. Error: %s\n", id, err)
			}else {
				fmt.Printf("Marked \"%d\" as completed\n", id)
			}

		}

	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
