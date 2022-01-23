package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(taskApp)
	taskApp.AddCommand(createTask)
	taskApp.AddCommand(showTasks)
}

var firstTask = task{1, "Primera", "Primera task", false, true}
var tasks = TaskList{[]*task{&firstTask}}

var taskApp = &cobra.Command{
	Use:     "task",
	Aliases: []string{"tarea"},
	Short:   "App to store task in memory.",
	Long:    `With this app you can create, update and delete many tasks as you want.`,
}

var createTask = &cobra.Command{
	Use:   "create",
	Short: "Create a new task.",
	Long:  `Create a new task.`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks.CreateTask(args[0], args[1])
	},
}

var showTasks = &cobra.Command{
	Use:   "show",
	Short: "Show all tasks.",
	Long:  `Show all tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks.ShowTasks()
	},
}

type task struct {
	id          int
	name        string
	description string
	is_complete bool
	is_active   bool
}

type TaskList struct {
	tasks []*task
}

func (t *task) print() {
	fmt.Printf("%+v\n", t)
}

func (t *TaskList) search(id int) (task, error) {
	for _, v := range t.tasks {
		if v.id == id {
			return *v, nil
		}
	}
	return task{}, errors.New("Task not found")
}

func (t *TaskList) ShowTasks() {
	for _, v := range t.tasks {
		v.print()
	}
}

// Create a new task
func (t *TaskList) CreateTask(name string, description string) {
	id := len(t.tasks)
	newTask := task{id + 1, name, description, false, true}
	t.tasks = append(t.tasks, &newTask)
	fmt.Print("Task created: ")
	newTask.print()
}

func (t *TaskList) CompleteTask(id int) {
	task, err := t.search(id)
	if err != nil {
		fmt.Println(err)
	} else {
		task.is_complete = true
	}
}

func (t *TaskList) DeleteTask(id int) {
	task, err := t.search(id)
	if err != nil {
		fmt.Println(err)
	} else {
		task.is_active = false
	}
}
