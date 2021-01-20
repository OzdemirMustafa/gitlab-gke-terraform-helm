package service

import (
	"backend-app/.mocks"
	"backend-app/model"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var requestTitle = "First Todo Item"

func TestGivenRequestTitleWhenICallAddTodoItemRepositoryThenICanSeeTodoItemResponse(t *testing.T) {
	expectedTodoItemResponse := model.TodoItem{ID: 1, Title: requestTitle, Completed: false}

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockTodoRepository := _mocks.NewMockTodoRepository(mockController)
	mockTodoRepository.EXPECT().
		AddTodoListItemByTitle(requestTitle).
		Return(expectedTodoItemResponse, nil).Times(1)

	actualResponse, _ := TodoService{mockTodoRepository}.AddTodoListItem(requestTitle)

	assert.Equal(t, expectedTodoItemResponse, actualResponse)
}

func TestGivenRequestTitleWhenICallAddTodoItemRepositoryAndRepositoryHasErrorThenICanSeeToGettingErrors(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockTodoRepository := _mocks.NewMockTodoRepository(mockController)
	mockTodoRepository.EXPECT().
		AddTodoListItemByTitle(requestTitle).
		Return(model.TodoItem{}, errors.New("todo repository has a error")).Times(1)

	_, err := TodoService{mockTodoRepository}.AddTodoListItem(requestTitle)

	assert.Error(t, err)
}

func TestGivenRequestWhenICallGetTodoListRepositoryThenICanSeeExpectedResponse(t *testing.T) {
	expectedTodoListResponse := []model.TodoList{
		{ID: 1, Title: "First Todo Item", Completed: false},
		{ID: 2, Title: "Second Todo Item", Completed: false},
	}

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockTodoRepository := _mocks.NewMockTodoRepository(mockController)
	mockTodoRepository.EXPECT().
		GetTodoListItems().
		Return(expectedTodoListResponse, nil).Times(1)

	actualResponse, _ := TodoService{mockTodoRepository}.GetTodoList()

	assert.Equal(t, expectedTodoListResponse, actualResponse)
}

func TestGivenRequestTitleWhenICallGetTodoListMethodAndMethodyHasErrorThenICanSeeToGettingErrors(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockTodoRepository := _mocks.NewMockTodoRepository(mockController)
	mockTodoRepository.EXPECT().
		GetTodoListItems().
		Return([]model.TodoList{}, errors.New("todolist has a error")).Times(1)

	_, err := TodoService{mockTodoRepository}.GetTodoList()

	assert.Error(t, err)
}

func TestGivenRequestWhenICallUpdateTodoItemRepositoryThenICanSeeTodoItemResponse(t *testing.T) {
	expectedTodoItemResponse := model.TodoItem{ID: 1, Title: requestTitle, Completed: true}

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockTodoRepository := _mocks.NewMockTodoRepository(mockController)
	mockTodoRepository.EXPECT().
		UpdateTodoListItem(1, true).
		Return(expectedTodoItemResponse, nil).Times(1)

	actualResponse, _ := TodoService{mockTodoRepository}.UpdateItemCompleted(1, true)

	assert.Equal(t, expectedTodoItemResponse, actualResponse)
}

func TestGivenRequestTitleWhenICallUpdateTodoItemRepositoryAndRepositoryHasErrorThenICanSeeToGettingErrors(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockTodoRepository := _mocks.NewMockTodoRepository(mockController)
	mockTodoRepository.EXPECT().
		UpdateTodoListItem(1, true).
		Return(model.TodoItem{}, errors.New("todo repository has a error")).Times(1)

	_, err := TodoService{mockTodoRepository}.UpdateItemCompleted(1, true)

	assert.Error(t, err)
}

func TestGivenRequestWhenICallDeleteTodoItemRepositoryThenICanSeeDeletedTodoItemResponse(t *testing.T) {
	expectedTodoItemResponse := model.TodoItem{ID: 1, Title: requestTitle, Completed: true}

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockTodoRepository := _mocks.NewMockTodoRepository(mockController)
	mockTodoRepository.EXPECT().
		DeleteTodoItemByID(1).
		Return(expectedTodoItemResponse, nil).Times(1)

	actualResponse, _ := TodoService{mockTodoRepository}.DeleteTodoItem(1)

	assert.Equal(t, expectedTodoItemResponse, actualResponse)
}

func TestGivenRequestTitleWhenICallDeleteTodoItemRepositoryAndRepositoryHasErrorThenICanSeeToGettingErrors(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockTodoRepository := _mocks.NewMockTodoRepository(mockController)
	mockTodoRepository.EXPECT().
		DeleteTodoItemByID(1).
		Return(model.TodoItem{}, errors.New("todo repository has a error")).Times(1)

	_, err := TodoService{mockTodoRepository}.DeleteTodoItem(1)

	assert.Error(t, err)
}
