# Task Tracker CLI

Task Tracker CLI is a simple command-line application written in Go for task management. It allows you to add, update, delete tasks, and track their status.

### Features

- Add a task
- Update task description
- Delete a task
- Mark a task as in progress
- Mark a task as done
- List all tasks
- List tasks by status (done, todo, in progress)
- Store tasks in a JSON file

### Installation

1. Make sure you have Go (>=1.18) installed.

2. Clone the repository:
   ```sh
   git clone https://github.com/che1nov/task-tracker-cli.git
   cd task-tracker-cli
   ```

3. Build and install:
   ```sh
   go build -o task-cli
   ```

### Usage

#### Add a new task

```sh
task-cli add "Buy groceries"
# Output: Task successfully added (ID: 1)
```

#### Update a task

```sh
task-cli update 1 "Buy groceries and cook dinner"
```

#### Delete a task

```sh
task-cli delete 1
```

#### Change task status

```sh
task-cli mark-in-progress 1
```

```sh
task-cli mark-done 1
```

#### List tasks

```sh
task-cli list
```

#### List tasks by status

```sh
task-cli list done
```

```sh
task-cli list todo
```

```sh
task-cli list in-progress
```

### Task Structure

Each task is represented in a JSON file with the following fields:

```json
{
  "id": 1,
  "description": "Buy groceries",
  "status": "todo",
  "createdAt": "2025-03-16T12:00:00Z",
  "updatedAt": "2025-03-16T12:00:00Z"
}
```

### Data Storage

The `tasks.json` file is used to store the list of tasks. If the file does not exist, it will be created automatically.

### Errors and Exception Handling

- Check for the existence of a task before updating or deleting
- Validate user input
- Handle errors when working with the JSON file

### License

This project is licensed under the MIT License.

https://roadmap.sh/projects/task-tracker