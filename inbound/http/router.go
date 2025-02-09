package http

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"net/http"
	"sourav.kabiraj/goboilerplate/logic/todos"
	"sourav.kabiraj/goboilerplate/logic/todos/models"
	"strconv"
)

func initialize(app *dig.Container) {
	app.Invoke(func(router *Router) {
		r := gin.Default()

		r.GET("/todos", router.GetTodos)
		r.POST("/todos", router.AddTodo)
		r.PUT("/todos", router.UpdateTodo)
		r.DELETE("/todos/:id", router.DeleteTodo)

		r.Run(":8080")
	})
}

type Router struct {
	manager *todos.Manager
}

func NewRouter(manager *todos.Manager) *Router {
	return &Router{manager: manager}
}

func (h *Router) GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, h.manager.GetTodos())
}

func (h *Router) AddTodo(c *gin.Context) {
	var newTodo models.Todo
	if err := c.Bind(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo := h.manager.AddTodo(newTodo)
	c.JSON(http.StatusCreated, todo)
}

func (h *Router) UpdateTodo(c *gin.Context) {
	var updatedTodo models.Todo
	if err := c.Bind(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := h.manager.UpdateTodo(updatedTodo)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func (h *Router) DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.manager.DeleteTodoByID(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
