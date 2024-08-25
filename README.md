![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/radek-nowak/todo/.github%2Fworkflows%2Fgo.yml)

# Todo CLI Application
A simple command-line application for managing tasks, built using Go and the Cobra library.

### Features
- Add, update, delete, and complete tasks
- View tasks as a table
- Persist tasks between sessions using JSON file storage

## Installation

### Prerequisites
Go 1.x: Make sure Go is installed on your system. You can download it from the official [Go website](https://go.dev/dl/).

### Via `go install`
```
go install github.com/yourusername/go_todo_app@latest
```

## Usage
When you run the todo CLI application for the first time, it automatically sets up a storage file where all your tasks will be saved. This ensures that your tasks are persistently stored across sessions and can be easily managed by the CLI.

By default, the tasks are stored in a file located at `~/.todo_app/data/todo_data.json`.
If this directory or file does not exist, the application automatically creates them during the first run.

### Commands
- `todo add [task]` - Add a new task.
- `todo show` - Show tasks (by dafault shows top 30 tasks).
- `todo show --top [number of tasks]` - Show specified number of tasks to display.
- `todo update [id] [new task]` - Update an existing task.
- `todo complete [id]` - Mark a task as completed.
- `todo delete [id]` - Delete a task by ID. You can either specify the ID directly or be prompted to enter it.
- `todo delete --from [start_index] --to [end_index]` - Delete a range of tasks by specifying the start (`--from`) and end (`--to`) indexes (inclusive). These flags are optional:
    - If neither `--from` nor `--to` is provided: You can delete a single task by ID.
    - If only `--from` is provided: Tasks from the specified index to the last task are deleted.
    - If only `--to` is provided: Tasks from the first task to the specified index are deleted.
    - If both `--from` and `--to` are provided: The specified range of tasks is deleted.
