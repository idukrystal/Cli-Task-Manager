// Functions to handle command calls add, delete etc

package command

import (
	"github.com/idukrystal/Cli-Task-Manager/src/status"
	"github.com/idukrystal/Cli-Task-Manager/src/task"
	"errors"
	"fmt"
	"strconv"
    "time"
)

const(
	NewDescNotProvided = "New description not provided."
	NotEnoghArgs = "Not Enogh Args"
	NotFound = "Not Found"
	TasksFile = "TASKS.json"
	// all date/time printed in this format
	TimeFormat = "Jan 02 2006, 03:04"
)

/* function to handle commands */

/* handles add: creates a new task and returns it's unique id
 * args: description of new task at index 2
 */
func AddNewTask(args []string) int {
	// if user did not provide all required arguments
	if len(args) < 3 {
		panic(errors.New(NewDescNotProvided))
	}

	tasks := task.ReadTasksFromFile(TasksFile)

	// generates a new unique id
	newId := getNextId(tasks)
	
	tasks[newId] = task.New(newId, args[2])
	
	task.WriteTasksToFile(tasks, TasksFile)
	return newId
}

/*handles update, delete, mark-in-progress and mark-done based on the value passed as status
 *args: should contain task id at index 2 and description at 3 (if applicable)
 *updateStatus
 * ****None: updates a github.com/idukrystal/Cli-Task-Manager description
 * ****Delete: removes a task
 * ****InProgress: changes a github.com/idukrystal/Cli-Task-Manager status to In progress
 * ****Done: changes a github.com/idukrystal/Cli-Task-Manager status to Done
 */
func UpdateTask(args []string, updateStatus status.Status) {
	count := len(args)

	var id int
	var err error

	// status = None, expects the description argument after index
	var expectedCount int = 4
	if updateStatus != status.None{
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

	tasks := task.ReadTasksFromFile(TasksFile)

	// present is false if id not in github.com/idukrystal/Cli-Task-Manager
	currentTask, present := tasks[id]
	if !present {
		panic(errors.New(NotFound))
	}

	if updateStatus == status.Delete {
		delete(tasks, id)
	} else {
		if updateStatus == status.None {
			currentTask.Description = args[3]
		} else {
			currentTask.Status = updateStatus
		}
		// updated at set to current time
		currentTask.UpdatedAt = time.Now().Format(TimeFormat)
		tasks[id] = currentTask
	}

	task.WriteTasksToFile(tasks, TasksFile)
}

/* handles list: list saved github.com/idukrystal/Cli-Task-Manager
 * args: can contain status filters at index 2 (done, inprogress, done)
 */
func ListTasks(args []string) {
	var valid bool = true
	var updateStatus status.Status
	
	if len(args) < 3 {
		updateStatus = status.None
	} else {
		// valid is false if args[2] not in allowed status
		updateStatus, valid = task.GetAllowedStatus(args[2])
		if !valid {
			panic(fmt.Errorf("Unkown status: %s",args[2]))
		}
	}
	
	tasks := task.ReadTasksFromFile(TasksFile)
	for id, task := range tasks {
		if task.Status == updateStatus || updateStatus == status.None {
			printTask(id, task)
		}
	}
}


// Prints each of task's field on a newline
func printTask(id int, task task.Task) {
	indent := "    "
	fmt.Printf("ID: %d, Description: %s\n", id, task.Description)
	fmt.Printf("%sStatus: %s\n", indent, task.Status)
	fmt.Printf("%sCreated: %s\n", indent, task.CreatedAt)

	// only print updatedAt if task as been updated before
	if task.UpdatedAt != "" {
		fmt.Printf("%sLast Update: %s\n", indent, task.UpdatedAt)
	}
}

// generates new unique ids for github.com/idukrystal/Cli-Task-Manager
func getNextId(tasks map[int]task.Task) (highestId int) {
	for id := range tasks {
		if id  > highestId {
			highestId = id
		}
	}
	highestId++
	return
}
