package cmd

import (
	"github.com/andres15alvarez/learning_go/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(taskApp)
	taskApp.AddCommand(createTask)
	taskApp.AddCommand(showTasks)
}

var firstTask = utils.Task{
	Id:          1,
	Name:        "Primera",
	Description: "Primera task",
	IsComplete:  false,
	IsActive:    true,
}
var tasks = []*utils.Task{&firstTask}
var taskList = utils.TaskList{Tasks: tasks}

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
		taskList.CreateTask(args[0], args[1])
	},
}

var showTasks = &cobra.Command{
	Use:   "show",
	Short: "Show all tasks.",
	Long:  `Show all tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		taskList.ShowTasks()
	},
}
