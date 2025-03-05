package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
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
	Delete Status = "Delete"
	None Status = "None"
	NewDescNotProvided = "New description not provided."
	NotEnoghArgs = "Not Enogh Args"
	NotFound = "Not Found"
	TasksFile = "TASKS.json"
	TimeFormat = "Jan 02 2006, 03:04"
	
)

var allowedStatus = map[string]Status{
	"done": Done, "todo": ToDo, "inprogress": InProgress,
}

func main() {
	// add, delete, update etc.
	command := os.Args[1]

	// handles any pannic(err) calls due to an error
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
		updateTask(os.Args, None)
	case "delete":
		updateTask(os.Args, Delete)
	case "mark-in-progress":
		updateTask(os.Args, InProgress)
	case "mark-done":
		updateTask(os.Args, Done)
	case "list":
		listTasks(os.Args)
	default:
		// an unknown command was entered
		fmt.Printf("Invalid Command: %s\n", command)
	}
}

/* function to handle commands */


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

// handles add: creates a new task and returns it's unique id
func addNewTask(args []string) int {
	// if user did not provide all required arguments
	if len(args) < 3 {
		panic(errors.New(NewDescNotProvided))
	}

	tasks := readTasksFromFile(TasksFile)

	// generates a new unique id
	newId := getNextId(tasks)
	
	tasks[newId] = Task {
		Description: args[2],
		Status: ToDo,
		// curent time
		CreatedAt: time.Now().Format(TimeFormat),	
	}

	writeTasksToFile(tasks, TasksFile)
	return newId
}

/*handles update, delete, mark-in-progress and mark-done based on the value passed as status
 *args: should contain task id at index 2 and description at 3 (if applicable)
 *Status
 * ****None: updates a tasks description
 * ****Delete: removes a task
 * ****InProgress: changes a tasks status to In progress
 * ****Done: changes a tasks status to Done
 */
func updateTask(args []string, status Status) {
	count := len(args)

	var id int
	var err error

	// status = None, expects the description argument after index
	var expectedCount int = 4
	if status != None{
		expectedCount = 3
	}

	
	if count > (expectedCount) {
		// convert to int, raises error in err if args[2] is not a number
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

	if status == Delete {
		delete(tasks, id)
	} else {
		if status == None {
			task.Description = args[3]
		} else {
			task.Status = status
		}
		task.UpdatedAt = time.Now().Format(TimeFormat)
		tasks[id] = task
	}

	writeTasksToFile(tasks, TasksFile)
	
}

func listTasks(args []string) {
	var valid bool = true
	var status Status
	if len(args) < 3 {
		status = None
	} else {
		status, valid = allowedStatus[args[2]]
		if !valid {
			panic(fmt.Errorf("Unkown status: %s",args[2]))
		}
	}
	tasks := readTasksFromFile(TasksFile)
	for id, task := range tasks {
		if task.Status == status || status == None {
			printTask(id, task)
		}
	}
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

func printTask(id int, task Task) {
	indent := "    "
	fmt.Printf("ID: %d, Description: %s\n", id, task.Description)
	fmt.Printf("%sStatus: %s\n", indent, task.Status)
	fmt.Printf("%sCreated: %s\n", indent, task.CreatedAt)
	if task.UpdatedAt != "" {
		fmt.Printf("%sLast Update: %s\n", indent, task.UpdatedAt)
	}
}
