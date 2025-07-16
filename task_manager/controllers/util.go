package controllers

import (
	"fmt"
	"task_manager/models"
	"time"
)

// isValidTask is a helping function used to check
// whether a task is valid (contains appropriate
// inputs or not)
// returns an error message if not
func isValidTask(task models.Task) error {
	// checks for invalid task attributes
	if task.Title == "" {
		return fmt.Errorf("title can not be empty")
	}
	if task.DueDate.Before(time.Now()) {
		return fmt.Errorf("due date can not be in the past")
	}
	if task.Status != string(models.COMPLETED) && task.Status != string(models.PENDING) && task.Status != string(models.MISSED) {
		return fmt.Errorf("invalid status")
	}
	return nil
}
