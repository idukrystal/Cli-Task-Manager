package main

import (
	"encoding/json"
	"fmt"
	"os"
)

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
