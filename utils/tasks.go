package utils

import (
	"errors"
	"fmt"
)

type Task struct {
	Id          int
	Name        string
	Description string
	IsComplete  bool
	IsActive    bool
}

type TaskList struct {
	Tasks []*Task
}

func (t *Task) print() {
	fmt.Printf("%+v\n", t)
}

func (t *TaskList) search(id int) (Task, error) {
	for _, v := range t.Tasks {
		if v.Id == id {
			return *v, nil
		}
	}
	return Task{}, errors.New("Task not found")
}

func (t *TaskList) ShowTasks() {
	for _, v := range t.Tasks {
		v.print()
	}
}

// Create a new task
func (t *TaskList) CreateTask(name string, description string) {
	id := len(t.Tasks)
	newTask := Task{id + 1, name, description, false, true}
	t.Tasks = append(t.Tasks, &newTask)
	fmt.Print("Task created: ")
	newTask.print()
}

func (t *TaskList) CompleteTask(id int) {
	task, err := t.search(id)
	if err != nil {
		fmt.Println(err)
	} else {
		task.IsComplete = true
	}
}

func (t *TaskList) DeleteTask(id int) {
	task, err := t.search(id)
	if err != nil {
		fmt.Println(err)
	} else {
		task.IsActive = false
	}
}
