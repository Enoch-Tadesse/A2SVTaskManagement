package router

import (
	"task_manager/controllers"

	"github.com/gin-gonic/gin"
)

func Run() error {
	router := gin.Default()

	tasks := router.Group("/tasks")
	{
		tasks.GET("", controllers.GetAllTasks)
		tasks.GET("/:id", controllers.GetTaskByID)
		tasks.PUT("/:id", controllers.UpdateTask)
		tasks.DELETE("/:id", controllers.DeleteTask)
		tasks.POST("", controllers.AddTask)
	}

	err := router.Run(":8080")
	if err != nil {
		return err
	}
	return nil
}
