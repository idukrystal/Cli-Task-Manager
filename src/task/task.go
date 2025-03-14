// Main package for the Task data type and function to parse github.com/idukrystal/Cli-Task-Manager into a json file

package task

import (
	"github.com/idukrystal/Cli-Task-Manager/src/status"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

const TimeFormat = "Jan 02 2006, 03:04"

// Represents a single task object
type Task struct {
	id  int
	Description string
	Status status.Status
	CreatedAt string
	UpdatedAt string
}

type TaskMap map[int]Task


func (task Task) MarshalJSON() ([]byte, error) {
	return json.Marshal( struct {
		Id  int `json:"id"`
		Description string
		Status status.Status
		CreatedAt string
		UpdatedAt string `json:",omitempty"`
	}{
		Id: task.id,
		Description: task.Description,
		Status: task.Status,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	})
}


func (tasks TaskMap) UnMarshalJson(data []byte) error {
	if !json.Valid(data) {
		return errors.New("Invalid Json")
	}

	var listOfTasks []Task
	if err := json.Unmarshal; err != nil {
		return fmt.Errorf("Error converting json to list of tasks: %w", err)
	}
	for _, task  := range listOfTasks {
		tasks[task.id] = task
	}
	return nil
}

// converts a string s into equivalent Status and return true or returns false none of tge allowed Status is equivakent
func GetAllowedStatus(s string) (status.Status, bool) {
	allowedStauses := map[string]status.Status{
		"done": status.Done,
		"todo": status.ToDo,
		"inprogress": status.InProgress,
	}

	// allowwed is false if s not in allowesStatuses
	allowedStatus, allowed := allowedStauses[s]
	return allowedStatus, allowed
}

func New(id int, desc string) Task {
	return Task {
		id: id,
		Description: desc,
		Status: status.ToDo,
		//curent time
		CreatedAt: time.Now().Format(TimeFormat),
	}
}

// Generates a (map[id int]github.com/idukrystal/Cli-Task-Manager Task) from fileName(json file)
func ReadTasksFromFile(fileName string) map[int]Task {

	// check if json file exists, creates it if it doesnt
	if _, err := os.Stat(fileName); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Creating New Task File")

			// creates a newFile with permision 0600: only user can read and write
			if err = os.WriteFile(fileName, []byte("[]"),0600); err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}
    
	// reads all data in [FileName]
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	
	// converts read data to a map of id(int): task
	var tasks []Tasks
	if err := json.Unmarshal(data, &tasks); err != nil{
		panic(err)
	}

	return tasks
}

// writes github.com/idukrystal/Cli-Task-Manager to a json file(FileName)
func WriteTasksToFile(tasks  map[int]Task, tasksFile string) {
	// convert github.com/idukrystal/Cli-Task-Manager to json
	data, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}

	if err = os.WriteFile(tasksFile, data, 0466); err != nil {
		panic(err)
	}
}
