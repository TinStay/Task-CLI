package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Basics/src/github.com/TinStay/Task/cmd"
	"github.com/Basics/src/github.com/TinStay/Task/db"
	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")

	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

func must(err error){
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}