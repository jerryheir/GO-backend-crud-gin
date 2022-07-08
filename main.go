package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

type response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var todos = []todo{
	{ID: "1", Item: "Have Call with Sola", Completed: false},
	{ID: "2", Item: "Pay for Light", Completed: false},
	{ID: "3", Item: "Go Crash Course", Completed: false},
}

func getMain(context *gin.Context) {
	var value = response{Status: "success", Message: "🎸 Hello World!"}

	context.IndentedJSON(http.StatusOK, value)
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, todos)
}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	}
	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}

func filterTodoById(id string) (*[]todo) {
	var newTodos = []todo{}
	for i, t := range todos {
		if t.ID != id {
			newTodos = append(newTodos, todos[i])
		}
	}
	return &newTodos
}

func deleteTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	}
	newTodos := filterTodoById(todo.ID)
	todos = *newTodos
	context.IndentedJSON(http.StatusOK, todos)
}

func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func main() {
	router := gin.Default()
	router.GET("/", getMain)
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodo)
	router.DELETE("/todos/:id", deleteTodo)
	router.Run("localhost:9090")
}
