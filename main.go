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
		} else {
			id = addNewTask(os.Args[2]);
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

func readTasksFromFile(fileName string) []Task {
	if _, err := os.Stat(TasksFile); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Creating New Task File")
			if err = os.WriteFile(TasksFile, []byte("[]"),0600); err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}
	
	data, err := os.ReadFile(TasksFile)
	if err != nil {
		panic(err)
	}
	
	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil{
		panic(err)
	}

	return tasks
}

func addNewTask(description string) int {
	tasks := readTasksFromFile(TasksFile)
	
	task := Task {
		ID: getNextId(tasks) + 1,
		Description: description,
		Status: ToDo,
	}

	tasks = append(tasks, task)

	writeTasksToFile(tasks, TasksFile)

	return task.ID
}

func writeTasksToFile(tasks []Task, tasksFile string) {
	data, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}

	if err = os.WriteFile(tasksFile, data, 0466); err != nil {
		panic(err)
	}
}

func getNextId(tasks []Task) (id int) {
	for _, task := range tasks {
		if task.ID > id {
			id = task.ID
		}
	}
	return
}
