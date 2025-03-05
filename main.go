package main

import (
	"encoding/json"
	"fmt"
	"os"
	//	"strconv"
)

type Task struct {
	ID int
	Description string
	Status Status
	CreatedAt string
	UpdatedAt string
}

type Status string

const(
	ToDo Status = "To Do"
	InProgress Status = "In Progress"
	Done Status = "Done"
	NewDescNotProvided = "ADD: New task description not provided"
	TasksFile = "TASKS.json"
)



func main() {
	argsLength := len(os.Args)
	command := os.Args[1]

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%s: %s\n",command, r)
		}
	}()
	var err error
	var id int
	switch command {
	case "add":
		if argsLength < 3 {
			fmt.Println(NewDescNotProvided)
		} else if id, err = addNewTask(os.Args[2]); err == nil {
			fmt.Printf("Task added successfully (ID: %d)\n", id)
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
	if err != nil {
		fmt.Println(err)
	}
}

func addNewTask(description string) (id int, err error) {

	if _, err = os.Stat(TasksFile); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Creating New Task File")
			if err = os.WriteFile(TasksFile, []byte("[]"),0600); err != nil {
				panic(err)
			}
		}
	}
	
	data, err := os.ReadFile(TasksFile)
	if err != nil {
		panic(err)
	}

	var tasks []Task
	json.Unmarshal(data, &tasks)

	
	task := Task{
		ID: getHighestId(tasks) + 1,
		Description: description,
		Status: ToDo,
	}

	tasks = append(tasks, task)

	data, err = json.Marshal(tasks)
	if err != nil {
		panic(err)
	}

	if err = os.WriteFile(TasksFile, data, 0466); err != nil {
		panic(err)
	}
	return
}

func getHighestId(tasks []Task) (id int) {
	for _, task := range tasks {
		if task.ID > id {
			id = task.ID
		}
	}
	return
}
