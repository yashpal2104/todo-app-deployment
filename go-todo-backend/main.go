// go-todo-backend/main.go

package main

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Task represents the structure of our task/note object.
type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// --- In-Memory "Database" ---
// For this simple example, we'll store tasks in memory.
// In a real application, we would use a database like PostgreSQL or SQLite.
var (
	tasks      = make(map[int]Task)
	nextTaskID = 1
	lock       = sync.Mutex{} // A mutex to handle concurrent requests safely
)

func main() {
	// Initialize Gin router
	router := gin.Default()

	// --- CORS Middleware ---
	// This is crucial for allowing the frontend (running on a different port)
	// to communicate with this backend.
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"} // The default Svelte dev server port
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	router.Use(cors.New(config))

	// --- API Routes ---
	api := router.Group("/api")
	{
		api.GET("/tasks", getTasks)
		api.POST("/tasks", addTask)
		api.PUT("/tasks/:id", updateTask)
		api.DELETE("/tasks/:id", deleteTask)
	}
	
	// Pre-populate with some data for demonstration
	tasks[nextTaskID] = Task{ID: nextTaskID, Title: "Learn Go & Gin", Completed: true}
	nextTaskID++
	tasks[nextTaskID] = Task{ID: nextTaskID, Title: "Build a Svelte Frontend", Completed: false}
	nextTaskID++

	// Start the server
	router.Run("0.0.0.0:8080") // Listen and serve on http://localhost:8080
}

// --- Route Handlers ---

// getTasks responds with the list of all tasks as JSON.
func getTasks(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()

	// Convert map to slice for consistent JSON output
	taskList := make([]Task, 0, len(tasks))
	for _, task := range tasks {
		taskList = append(taskList, task)
	}

	c.JSON(http.StatusOK, taskList)
}

// addTask adds a new task from the JSON received in the request body.
func addTask(c *gin.Context) {
	var newTask Task

	// Bind the received JSON to the newTask struct
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lock.Lock()
	defer lock.Unlock()

	newTask.ID = nextTaskID
	newTask.Completed = false // New tasks are always incomplete
	tasks[newTask.ID] = newTask
	nextTaskID++

	c.JSON(http.StatusCreated, newTask)
}

// updateTask updates an existing task's status (e.g., toggles 'completed').
func updateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedTask Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lock.Lock()
	defer lock.Unlock()

	if _, ok := tasks[id]; ok {
		tasks[id] = updatedTask // Replace the whole task
		c.JSON(http.StatusOK, updatedTask)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
	}
}

// deleteTask removes a task by its ID.
func deleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	lock.Lock()
	defer lock.Unlock()

	if _, ok := tasks[id]; ok {
		delete(tasks, id)
		c.Status(http.StatusNoContent) // Success, no content to return
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
	}
}
