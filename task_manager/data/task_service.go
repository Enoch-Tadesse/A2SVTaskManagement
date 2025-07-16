package data

import (
	"fmt"
	"slices"
	"sync"
	"task_manager/models"
)

type TaskService struct {
	tasks []*models.Task
	mutex sync.RWMutex
	next  int
}

// NewTaskService returns a new
// initialized TaskService struct
func NewTaskService() *TaskService {
	return &TaskService{
		tasks: make([]*models.Task, 0),
		next:  1,
	}
}

// GetAllTasks is a function that return
// all tasks currently available
func (ts *TaskService) GetAllTasks() []models.Task {
	ts.mutex.RLock()
	defer ts.mutex.RUnlock()
	result := make([]models.Task, len(ts.tasks))
	for i, task := range ts.tasks {
		result[i] = *task
	}
	return result
}

// AddTask is a method used to add
// a new task into the tasks slice
func (ts *TaskService) AddTask(task models.Task) (models.Task, error) {
	// lock the count to increment it
	ts.mutex.Lock()
	task.ID = ts.next
	ts.next += 1
	ts.tasks = append(ts.tasks, &task)
	ts.mutex.Unlock()

	return task, nil

}

// GetTaskByID iterates over all tasks and returns the task
// if not found, returns an error
func (ts *TaskService) GetTaskByID(id int) (models.Task, error) {
	ts.mutex.RLock()
	defer ts.mutex.RUnlock()

	for _, t := range ts.tasks {
		if t.ID == id {
			return *t, nil
		}
	}
	return models.Task{}, fmt.Errorf("task does not exist")
}

// UpdateTask finds a task by id and updates it with the new modesl
// if task is not found it returns an error
func (ts *TaskService) UpdateTask(id int, modTask models.Task) (models.Task, error) {
	ts.mutex.Lock()
	defer ts.mutex.Unlock()

	// find the task and update it
	for _, task := range ts.tasks {
		if task.ID == id {
			task.Title = modTask.Title
			task.Description = modTask.Description
			task.DueDate = modTask.DueDate
			task.Status = modTask.Status
			// return the new updated task
			return *task, nil
		}
	}

	return models.Task{}, fmt.Errorf("task not found")

}

// DeleteTaskByID removes a task from tasks slice
// returns an error if task is not found
func (ts *TaskService) DeleteTaskByID(id int) error {
	ts.mutex.Lock()
	defer ts.mutex.Unlock()

	for i, task := range ts.tasks {
		if task.ID == id {
			ts.tasks = slices.Delete(ts.tasks, i, i+1)
			return nil
		}
	}
	return fmt.Errorf("task with id %d is not found", id)
}
