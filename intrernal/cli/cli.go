package cli

import (
	"fmt"
	"strconv"
	"task-tracker/intrernal/usecase"
)

// CLI управляет интерфейсом командной строки
type CLI struct {
	usecase *usecase.TaskUsecase
}

// NewCLI создает новый экземпляр CLI
func NewCLI(usecase *usecase.TaskUsecase) *CLI {
	return &CLI{usecase: usecase}
}

// Run запускает CLI
func (c *CLI) Run(args []string) {
	if len(args) < 2 {
		printUsage()
		return
	}

	command := args[1]

	switch command {
	case "add":
		if len(args) < 3 {
			fmt.Println("Error: Missing task description.")
			return
		}
		id, err := c.usecase.AddTask(args[2])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Task added successfully (ID: %d)\n", id)

	case "update":
		if len(args) < 4 {
			fmt.Println("Error: Missing task ID or new description.")
			return
		}
		taskID, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Error: Task ID must be an integer.")
			return
		}
		err = c.usecase.UpdateTask(taskID, args[3])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Task updated successfully (ID: %d)\n", taskID)

	case "delete":
		if len(args) < 3 {
			fmt.Println("Error: Missing task ID.")
			return
		}
		taskID, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Error: Task ID must be an integer.")
			return
		}
		err = c.usecase.DeleteTask(taskID)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Task deleted successfully (ID: %d)\n", taskID)

	case "mark-in-progress":
		if len(args) < 3 {
			fmt.Println("Error: Missing task ID.")
			return
		}
		taskID, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Error: Task ID must be an integer.")
			return
		}
		err = c.usecase.MarkTaskStatus(taskID, "in-progress")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Task marked as in-progress (ID: %d)\n", taskID)

	case "mark-done":
		if len(args) < 3 {
			fmt.Println("Error: Missing task ID.")
			return
		}
		taskID, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Error: Task ID must be an integer.")
			return
		}
		err = c.usecase.MarkTaskStatus(taskID, "done")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Task marked as done (ID: %d)\n", taskID)

	case "list":
		status := ""
		if len(args) > 2 {
			status = args[2]
			if status != "todo" && status != "in-progress" && status != "done" {
				fmt.Println("Error: Invalid status. Use 'todo', 'in-progress', or 'done'.")
				return
			}
		}
		tasks, err := c.usecase.ListTasks(status)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}
		for _, task := range tasks {
			fmt.Printf("ID: %d, Description: %s, Status: %s\n", task.ID, task.Description, task.Status)
		}

	default:
		fmt.Printf("Error: Unknown command '%s'.\n", command)
		printUsage()
	}
}

// printUsage выводит справку по использованию CLI
func printUsage() {
	fmt.Println("Usage: task-cli <command> [args]")
	fmt.Println("Commands:")
	fmt.Println("  add <description>                     Add a new task")
	fmt.Println("  update <id> <new-description>         Update a task's description")
	fmt.Println("  delete <id>                           Delete a task")
	fmt.Println("  mark-in-progress <id>                Mark a task as in-progress")
	fmt.Println("  mark-done <id>                       Mark a task as done")
	fmt.Println("  list [status]                        List tasks (optional: filter by status)")
	fmt.Println("                                       Valid statuses: todo, in-progress, done")
}
