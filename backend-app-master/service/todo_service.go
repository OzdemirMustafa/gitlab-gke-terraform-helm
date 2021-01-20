package service

import (
	"backend-app/model"
)

type TodoRepository interface {
	AddTodoListItemByTitle(string) (model.TodoItem, error)
	GetTodoListItems() ([]model.TodoList, error)
	UpdateTodoListItem(int, bool) (model.TodoItem, error)
	DeleteTodoItemByID(int) (model.TodoItem, error)
}

type TodoService struct {
	TodoRepository TodoRepository
}

func (t TodoService) AddTodoListItem(title string) (model.TodoItem, error) {
	addTodoItem, err := t.TodoRepository.AddTodoListItemByTitle(title)
	if err != nil {
		return model.TodoItem{}, err
	}

	return addTodoItem, nil
}

func (t TodoService) GetTodoList() ([]model.TodoList, error) {
	getTodoItemList, err := t.TodoRepository.GetTodoListItems()
	if err != nil {
		return []model.TodoList{}, err
	}

	return getTodoItemList, nil
}

func (t TodoService) UpdateItemCompleted(id int, completed bool) (model.TodoItem, error) {
	updateTodoItem, err := t.TodoRepository.UpdateTodoListItem(id, completed)
	if err != nil {
		return model.TodoItem{}, err
	}

	return updateTodoItem, nil
}

func (t TodoService) DeleteTodoItem(id int) (model.TodoItem, error) {
	deleteTodoItem, err := t.TodoRepository.DeleteTodoItemByID(id)
	if err != nil {
		return model.TodoItem{}, err
	}

	return deleteTodoItem, nil
}
