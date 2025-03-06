package task

import (
	"tasks/src/status"
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	Description string
	Status status.Status
	CreatedAt string
	UpdatedAt string
}

func GetAllowedStatus(s string) (status.Status, bool) {
	allowedStauses := map[string]status.Status{
		"done": status.Done,
		"todo": status.ToDo,
		"inprogress": status.InProgress,
	}

	allowedStatus, allowed := allowedStauses[s]
	return allowedStatus, allowed
}

// Generates a (map[id int]tasks Task) from fileName(json file)
func ReadTasksFromFile(fileName string) map[int]Task {

	// check if json file exists, creates it if it doesnt
	if _, err := os.Stat(fileName); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Creating New Task File")

			// creates a newFile with permision 0600: only user can read and write
			if err = os.WriteFile(fileName, []byte("{}"),0600); err != nil {
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
	var tasks map[int] Task
	if err := json.Unmarshal(data, &tasks); err != nil{
		panic(err)
	}

	return tasks
}

// writes tasks to a json file(FileName)
func WriteTasksToFile(tasks map[int]Task, tasksFile string) {
	// convert tasks to json
	data, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}

	if err = os.WriteFile(tasksFile, data, 0466); err != nil {
		panic(err)
	}
}
