package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
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
	NewDescNotProvided = "New description not provided."
	NotEnoghArgs = "Not Enogh Args"
	NotFound = "Not Found"
	TasksFile = "TASKS.json"
	
)



func main() {
	command := os.Args[1]

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%s: %s\n",command, r)
		}
	}()

	switch command {
	case "add":
		id := addNewTask(os.Args)
		fmt.Printf("Task added successfully (ID: %d)\n", id)
		
	case "update":
		updateTask(os.Args)
	/*case "delete":
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
}

func readTasksFromFile(fileName string) map[int]Task {
	if _, err := os.Stat(TasksFile); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Creating New Task File")
			if err = os.WriteFile(TasksFile, []byte("{}"),0600); err != nil {
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
	
	var tasks map[int] Task
	if err := json.Unmarshal(data, &tasks); err != nil{
		panic(err)
	}

	return tasks
}

func addNewTask(args []string) int {

	if len(args) < 3 {
		panic(errors.New(NewDescNotProvided))
	}

	tasks := readTasksFromFile(TasksFile)
	newId := getNextId(tasks)
	
	tasks[newId] = Task {
		Description: args[2],
		Status: ToDo,
	}

	writeTasksToFile(tasks, TasksFile)

	return newId
}

func updateTask(args []string) {
	count := len(args)

	var id int
	var err error
	
	if count > 3 {
	id, err = strconv.Atoi(args[2])
		if err != nil{
			panic(fmt.Errorf("Not a valid ID: %s", args[2]))
		}
	} else {
		panic(errors.New(NotEnoghArgs))
	}

	tasks := readTasksFromFile(TasksFile)

	task, present := tasks[id]
	if !present {
		panic(errors.New(NotFound))
	}

	task.Description = args[3]

	tasks[id] = task

	fmt.Println(tasks)

	writeTasksToFile(tasks, TasksFile)
	
}

func writeTasksToFile(tasks map[int]Task, tasksFile string) {
	data, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}

	if err = os.WriteFile(tasksFile, data, 0466); err != nil {
		panic(err)
	}
}

func getNextId(tasks map[int]Task) (highestId int) {
	for id := range tasks {
		if id  > highestId {
			highestId = id
		}
	}
	highestId++
	return
}
