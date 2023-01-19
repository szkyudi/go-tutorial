package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"szkyudi/go-tutorial/model"
)

type TodoController interface {
	FetchTodos(w http.ResponseWriter, r *http.Request)
	AddTodo(w http.ResponseWriter, r *http.Request)
	ChangeTodo(w http.ResponseWriter, r *http.Request)
	DeleteTodo(w http.ResponseWriter, r *http.Request)
}

type todoController struct {
	todoModel model.TodoModel
}

func NewTodoController(todoModel model.TodoModel) TodoController {
	return &todoController{todoModel}
}

func (controller *todoController) FetchTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := controller.todoModel.FetchTodos()

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	json, err := json.Marshal(todos)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, string(json))
}

func (controller *todoController) AddTodo(w http.ResponseWriter, r *http.Request) {
	result, err := controller.todoModel.AddTodo(r)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	json, err := json.Marshal(result)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, string(json))
}

func (controller *todoController) ChangeTodo(w http.ResponseWriter, r *http.Request) {
	result, err := controller.todoModel.ChangeTodo(r)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	json, err := json.Marshal(result)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, string(json))
}

func (controller *todoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	result, err := controller.todoModel.DeleteTodo(r)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	json, err := json.Marshal(result)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, string(json))
}
