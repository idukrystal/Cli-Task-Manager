# CLI Task Manager
**project url**: https://roadmap.sh/projects/task-tracker

CLI Task Manager is a simple command-line application built in Go to help you manage your tasks. It allows you to track what you need to do, what you're currently working on, and what you've completed. The tasks are stored in a JSON file, making it easy to persist and manage your tasks across sessions.

## Features

- **Add Tasks**: Add new tasks to your task list.
- **Update Tasks**: Modify existing tasks.
- **Delete Tasks**: Remove tasks from your list.
- **Mark Tasks**: Mark tasks as "in progress" or "done".
- **List Tasks**: View all tasks or filter them by status (done, in progress, or not done).

## Requirements

- Go (Golang) installed on your machine.
- Basic familiarity with the command line.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/idukrystal/Cli-Task-Manager.git
   ```
2. Navigate to the project directory:
   ```bash
   cd Cli-Task-Manager
   ```
3. Build the project:
   ```bash
   go build -o tasks
   ```
4. Run the application:
   ```bash
   ./tasks [command] [args...]
   ```

## Usage

The CLI Task Manager supports the following commands:

### Add a Task
```bash
./tasks add "Task description"
```

### Update a Task
```bash
./tasks update [task-id] "New task description"
```

### Delete a Task
```bash
./tasks delete [task-id]
```

### Mark a Task as In Progress
```bash
./tasks mark-in-progress [task-id]
```

### Mark a Task as Done
```bash
./tasks mark-done [task-id]
```

### List All Tasks
```bash
./tasks list
```

### List Done Tasks
```bash
./tasks list done
```

### List In Progress Tasks
```bash
./tasks list in-progress
```

### List Not Done Tasks
```bash
./tasks list todo
```

## Contributing

Contributions are welcome! If you'd like to contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Commit your changes.
4. Push your branch and open a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built with Go (Golang).
- Inspired by simple CLI tools for productivity.

---

Feel free to reach out if you have any questions or suggestions!
```
