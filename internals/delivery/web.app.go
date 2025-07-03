package delivery

import (
	"golang-crud-api/internals/delivery/dependencies"
	"golang-crud-api/internals/interfaces/handlers"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start() {
	container := dependencies.Setup()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	err := container.Invoke(func(h *handlers.TaskHandler) {
		router.POST("/tasks", h.CreateTask)
		router.GET("/tasks", h.GetTasks)
		router.PUT("/tasks/:id", h.UpdateTask)
		router.DELETE("/tasks/:id", h.DeleteTask)

		log.Println("Server running on http://localhost:8080")
		router.Run(":8080")
	})
	if err != nil {
		log.Fatalf("error while setting up dependencies\n\tError: %v", err)
	}
}
