package main

import (
	"fmt"
	"os"
	//	"strconv"
)

type task struct {
	id int
	description string
	status Status
	createdAt string
	updatedAt string
}

type Status string

const(
	ToDo Status = "To Do"
	InProgress Status = "In Progress"
	Done Status = "Done"
	NewDescNotProvided = "ADD: New task description not provided"
)



func main() {
	argsLength := len(os.Args)
	command := os.Args[1]

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	switch command {
	case "add":
		if argsLength < 3 {
			fmt.Println(NewDescNotProvided)
		} else if id, err := addNewTask(os.Args[2]); err == nil {
			fmt.Printf("Task added successfully (ID: 1)\n", id)
		}
		/*case "update":
		err := updateTask(strconv.Atoi(os.Args[2]), os.Args[3])
	case "delete":
		err := deleteTask(strconv.Atoi(os.Args[2]))
	case "mark-in-progress":
		err := updateProgress(strconv.Atoi(os.Args[2]), InProgress)
	case "mark-done":
		err := updateProgress(strconv.Atoi(os.Args[2]), Done)
	case "list":
		err := listTasks(nil)
	case "list done":
		err := listTasks(Done)
	case "list todo":
		err := listTasks(ToDo)
	case "list in-progress":
		err := listTasks(InProgress)*/
	default:
		fmt.Printf("Invalid Command: %s\n", command)
		
	}
	// handle errors
}

func addNewTask(description string) (id int, err error) {
	fmt.Println(description)
	return
}
