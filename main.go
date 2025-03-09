// Task Manager Main Package

package main

import (
	"github.com/idukrystal/Cli-Task-Manager/src/command"
	"github.com/idukrystal/Cli-Task-Manager/src/status"
	"errors"
	"fmt"
	"os"
)

func main() {
	var cmd string
	
	// handles any pannic(err) calls due to an error
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%s: %s\n",cmd, r)
		}
	}()
	
	if len(os.Args) < 2 {
		panic(errors.New(command.NotEnoghArgs))
	}
	
	// add, delete, update etc.
	cmd = os.Args[1]

	switch cmd {
	case "add":
		id := command.AddNewTask(os.Args)
		fmt.Printf("Task added successfully (ID: %d)\n", id)
	case "update":
		command.UpdateTask(os.Args, status.None)
	case "delete":
		command.UpdateTask(os.Args, status.Delete)
	case "mark-in-progress":
		command.UpdateTask(os.Args, status.InProgress)
	case "mark-done":
		command.UpdateTask(os.Args, status.Done)
	case "list":
		command.ListTasks(os.Args)
	default:
		// an unknown command was entered
		fmt.Printf("Invalid Command: %s\n", cmd)
	}
}


