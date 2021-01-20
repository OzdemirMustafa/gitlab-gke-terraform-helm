package handler

import (
	"backend-app/model"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"strings"
)

type TodoService interface {
	AddTodoListItem(string) (model.TodoItem, error)
	GetTodoList() ([]model.TodoList, error)
	UpdateItemCompleted(int, bool) (model.TodoItem, error)
	DeleteTodoItem(int) (model.TodoItem, error)
}

type TodoHandler struct {
	TodoService TodoService
}

func CheckTodoItemTitleEmpty(title string) bool {
	return strings.TrimSpace(title) == ""
}

func (t TodoHandler) AddTodoItem(context echo.Context) error {
	if CheckTodoItemTitleEmpty(context.Param("title")) {
		return context.NoContent(http.StatusBadRequest)
	}

	addTodoList, _ := t.TodoService.AddTodoListItem(context.Param("title"))
	if CheckTodoItemTitleEmpty(addTodoList.Title) {
		return context.NoContent(http.StatusNotFound)
	}

	return context.JSON(http.StatusOK, addTodoList)
}

func (t TodoHandler) GetTodoList(context echo.Context) error {
	getTodoList, err := t.TodoService.GetTodoList()
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, getTodoList)
}

func (t TodoHandler) UpdateTodoItemCompleted(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	completed, err := strconv.ParseBool(context.Param("completed"))
	if err != nil {
		_ = context.NoContent(http.StatusBadRequest)
		return err
	}

	updateTodoItem, err := t.TodoService.UpdateItemCompleted(id, completed)
	if err != nil {
		_ = context.NoContent(http.StatusNotFound)
		return err
	}

	return context.JSON(http.StatusOK, updateTodoItem)
}

func (t TodoHandler) DeleteTodoItem(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		_ = context.JSON(http.StatusBadRequest, err)

		return err
	}

	deleteTodoItem, err := t.TodoService.DeleteTodoItem(id)
	if err != nil {
		_ = context.JSON(http.StatusNotFound, err)

		return err
	}

	return context.JSON(http.StatusOK, deleteTodoItem)
}
