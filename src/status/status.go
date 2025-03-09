// Varriable and data types to represent a  github.com/idukrystal/Cli-Task-Manager possible state
package status

// Represents the status of a Task
type Status string

// Posible Atatus values
const (
	ToDo Status = "To Do"
	InProgress Status = "In Progress"
	Done Status = "Done"
	Delete Status = "Delete"
	None Status = "None"
)


