package tasks

import (
	"bufio"
	"fmt"
	"os"

	"strings"

	"github.com/spf13/cobra"
)

var tasks []string
var todoFile = "tasks.txt"

var RootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A simple to-do list CLI app",
}

func loadTasks() {
	file, err := os.OpenFile(todoFile, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening or creating tasks file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tasks = append(tasks, scanner.Text())
	}
}

func saveTasks() {
	file, err := os.OpenFile(todoFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening or creating tasks file:", err)
		return
	}
	defer file.Close()

	for _, task := range tasks {
		fmt.Fprintln(file, task)
	}
}

func addTask(cmd *cobra.Command, args []string) {
	task := args[0]
	tasks = append(tasks, task)
	saveTasks()
	fmt.Printf("Added task: %s\n", task)
}

func listTasks(cmd *cobra.Command, args []string) {
	if len(tasks) == 0 {
		fmt.Println("No tasks in the to-do list.")
	} else {
		fmt.Println("To-Do List:")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task)
		}
	}
}

func completeTask(cmd *cobra.Command, args []string) {
	task := args[0]
	for i, t := range tasks {
		if strings.EqualFold(t, task) {
			tasks = append(tasks[:i], tasks[i+1:]...)
			saveTasks()
			fmt.Printf("Completed task: %s\n", task)
			return
		}
	}
	fmt.Printf("Task not found: %s\n", task)
}

func updateTask(cmd *cobra.Command, args []string) {
	oldTask := args[0]
	newTask := args[1]
	for i, t := range tasks {
		if strings.EqualFold(t, oldTask) {
			tasks[i] = newTask
			saveTasks()
			fmt.Printf("Updated task: %s -> %s\n", oldTask, newTask)
			return
		}
	}
	fmt.Printf("Task not found: %s\n", oldTask)
}

func deleteTask(cmd *cobra.Command, args []string) {
	task := args[0]
	for i, t := range tasks {
		if strings.EqualFold(t, task) {
			tasks = append(tasks[:i], tasks[i+1:]...)
			saveTasks()
			fmt.Printf("Deleted task: %s\n", task)
			return
		}
	}
	fmt.Printf("Task not found: %s\n", task)
}

func AddCommands() {
	var addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add a new task to the to-do list",
		Args:  cobra.ExactArgs(1),
		Run:   addTask,
	}

	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all tasks in the to-do list",
		Run:   listTasks,
	}

	var completeCmd = &cobra.Command{
		Use:   "complete",
		Short: "Mark a task as complete",
		Args:  cobra.ExactArgs(1),
		Run:   completeTask,
	}

	var updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update a task in the to-do list",
		Args:  cobra.ExactArgs(2),
		Run:   updateTask,
	}

	var deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete a task from the to-do list",
		Args:  cobra.ExactArgs(1),
		Run:   deleteTask,
	}

	RootCmd.AddCommand(addCmd, listCmd, completeCmd, updateCmd, deleteCmd)
}
