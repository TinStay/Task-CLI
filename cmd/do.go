package cmd

import (
	"fmt"
	"strconv"

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
		if(len(ids) > 1){
			fmt.Printf("Tasks \"%v\" are marked as completed", ids)
		}else{
			fmt.Printf("Tasks \"%v\" is marked as completed", ids)
		}

	},
}

func init() {
	RootCmd.AddCommand(doCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
