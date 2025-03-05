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
	var command string
	
	// handles any pannic(err) calls due to an error
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%s: %s\n",command, r)
		}
	}()
	
	if len(os.Args) < 2 {
		panic(errors.New(NotEnoghArgs))
	}
	
	// add, delete, update etc.
	command = os.Args[1]

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


/* handles add: creates a new task and returns it's unique id
 * args: description of new task at index 2
 */
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

	
	if count >= expectedCount {
		// convert to int, raises error in err if args[2] is not a number
		id, err = strconv.Atoi(args[2])
		if err != nil{
			panic(fmt.Errorf("Not a valid ID: %s", args[2]))
		}
	} else {
		panic(errors.New(NotEnoghArgs))
	}

	tasks := readTasksFromFile(TasksFile)

	// present is false if id not in tasks
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
		// updated at set to current time
		task.UpdatedAt = time.Now().Format(TimeFormat)
		tasks[id] = task
	}

	writeTasksToFile(tasks, TasksFile)
}

/* handles list: list saved tasks
 * args: can contain status filters at index 2 (done, inprogress, done)
 */
func listTasks(args []string) {
	var valid bool = true
	var status Status
	
	if len(args) < 3 {
		status = None
	} else {
		// valid is false if args[2] not in allowed status
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

// Generates a (map[id int]tasks Task) from fileName(json file)
func readTasksFromFile(fileName string) map[int]Task {

	// check if json file exists, creates it if it doesnt
	if _, err := os.Stat(TasksFile); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Creating New Task File")

			// creates a newFile with permision 0600: only user can read and write
			if err = os.WriteFile(TasksFile, []byte("{}"),0600); err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	// reads all data in [FileName]
	data, err := os.ReadFile(TasksFile)
	if err != nil {
		panic(err)
	}

	// converts read data to a map of id(int): task
	var tasks map[int] Task
	if err := json.Unmarshal(data, &tasks); err != nil{
		panic(err)
	}

	return tasks
}

// writes tasks to a json file(FileName)
func writeTasksToFile(tasks map[int]Task, tasksFile string) {
	// convert tasks to json
	data, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}

	if err = os.WriteFile(tasksFile, data, 0466); err != nil {
		panic(err)
	}
}

// generates new unique ids for tasks
func getNextId(tasks map[int]Task) (highestId int) {
	for id := range tasks {
		if id  > highestId {
			highestId = id
		}
	}
	highestId++
	return
}

// Prints each of task's field on a newline
func printTask(id int, task Task) {
	indent := "    "
	fmt.Printf("ID: %d, Description: %s\n", id, task.Description)
	fmt.Printf("%sStatus: %s\n", indent, task.Status)
	fmt.Printf("%sCreated: %s\n", indent, task.CreatedAt)

	// only print updatedAt if task as been updated before
	if task.UpdatedAt != "" {
		fmt.Printf("%sLast Update: %s\n", indent, task.UpdatedAt)
	}
}

