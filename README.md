# Task CLI

## Exercise details

This is an exercise for building a CLI tool that can be used to manage TODOs in the terminal. The basic usage of the tool is going to look roughly like this:

```
$ task
task is a CLI for managing your TODOs.

Usage:
  task [command]

Available Commands:
  add         Add a new task to your TODO list
  do          Mark a task on your TODO list as complete
  list        List all of your incomplete tasks

Use "task [command] --help" for more information about a command.

$ task add review talk proposal
Added "review talk proposal" to your task list.

$ task add clean dishes
Added "clean dishes" to your task list.

$ task list
You have the following tasks:
1. review talk proposal
2. some task description

$ task do 1
You have completed the "review talk proposal" task.

$ task list
You have the following tasks:
1. some task description
```

*Note: Lines prefixed with `$` are lines where we type into the terminal, and other lines are output from our program.*

The main functionalities this application will have are: 

- `add` - adds a new task to our list
- `list` - lists all of our incomplete tasks
- `do` - marks a task as complete
