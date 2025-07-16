package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

var ts *data.TaskService = data.NewTaskService()

// GetAllTasks handles the HTTP GET request to retrieve all tasks.
func GetAllTasks(c *gin.Context) {
	tasks := ts.GetAllTasks()
	c.JSON(http.StatusOK, tasks)
}

// GetTaskByID handles the HTTP GET request to retrive a task by ID.
func GetTaskByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	// id is NaN or a negative number
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	task, err := ts.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("failed to retrieve task: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

// DeleteTask handles the HTTP DELETE request to remove a task by ID.
func DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	// id is NaN or a negative number
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	// delete the task
	if err := ts.DeleteTaskByID(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("failed to delete task: %s", err.Error()),
		})
		return
	}

	// send the
	c.JSON(http.StatusOK, gin.H{
		"message": "task deleted successfully",
	})

}

// UpdateTask handles the HTTP PUT request to fully replace a task with a new one.
func UpdateTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	// id is NaN or a negative number
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	var body models.Task

	// read the request body
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}
	// convert the status into lowercase
	body.Status = strings.ToLower(body.Status)

	// check for valid struct
	if err := isValidTask(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// update the task
	newTask, err := ts.UpdateTask(id, body)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "task updated successfully",
		"task":    newTask,
	})
}

// AddTask handles the HTTP POST request to insert a task into task collections.
func AddTask(c *gin.Context) {
	var body models.Task

	// read the request body
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	// convert the status into lowercase
	body.Status = strings.ToLower(body.Status)

	// check for valid struct
	if err := isValidTask(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newTask, err := ts.AddTask(body)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "task created successfully",
		"task":    newTask,
	})

}
