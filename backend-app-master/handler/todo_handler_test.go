package handler

import (
	"backend-app/.mocks"
	"backend-app/model"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var invalidTitleTodoListItem = ""
var requestTitle = "First Todo Item"

var actualTodoListItemResponseBody model.TodoItem
var expectedTodoListItemResponseBody = model.TodoItem{ID: 1, Title: requestTitle, Completed: false}

var expectedUpdateItemCompletedResponseBody = model.TodoItem{ID: 1, Title: requestTitle, Completed: true}

var expectedTodoListResponseBody = []model.TodoList{
	{ID: 1, Title: "First Todo Item", Completed: false},
	{ID: 2, Title: "Second Todo Item", Completed: false},
}

var actualTodoListResponseBody []model.TodoList
var expectedTodoListEmptyResponse = []model.TodoList{}

func getAddTodoListItemContext(requestTitle string) (echo.Context, *httptest.ResponseRecorder) {
	request := httptest.NewRequest(http.MethodPost, "/addTodoListItem/:title", http.NoBody)
	recorder := httptest.NewRecorder()
	context := echo.New().NewContext(request, recorder)
	context.SetParamNames("title")
	context.SetParamValues(requestTitle)

	return context, recorder
}

func getUpdateTodoItemContext(requestID, requestTitle string) (echo.Context, *httptest.ResponseRecorder) {
	request := httptest.NewRequest(http.MethodPost, "/updateItemCompleted/:id/:completed", http.NoBody)
	recorder := httptest.NewRecorder()
	context := echo.New().NewContext(request, recorder)

	context.SetParamNames("id")
	context.SetParamValues(requestID)

	context.SetParamNames("completed")
	context.SetParamValues(requestTitle)

	return context, recorder
}

func getDeleteTodoItemContext(requestID string) (echo.Context, *httptest.ResponseRecorder) {
	request := httptest.NewRequest(http.MethodPost, "/updateItemCompleted/:id/:completed", http.NoBody)
	recorder := httptest.NewRecorder()
	context := echo.New().NewContext(request, recorder)

	context.SetParamNames("id")
	context.SetParamValues(requestID)

	return context, recorder
}

func getTodoListContext() (echo.Context, *httptest.ResponseRecorder) {
	request := httptest.NewRequest(http.MethodGet, "/getTodoList/", http.NoBody)
	recorder := httptest.NewRecorder()
	context := echo.New().NewContext(request, recorder)

	return context, recorder
}

func TestGivenInvalidRequestDataWhenIRequestToAddTodoListThenItShouldReturnStatusBadRequest(t *testing.T) {
	mockController := gomock.NewController(t)

	mockTodo := _mocks.NewMockTodoService(mockController)
	mockTodo.EXPECT().
		AddTodoListItem(invalidTitleTodoListItem).
		Return(model.TodoItem{}, nil).Times(1)

	context, recorder := getAddTodoListItemContext(invalidTitleTodoListItem)
	_ = TodoHandler{mockTodo}.AddTodoItem(context)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestGivenRequestDataWhenIRequestToAddTodoListThenAndIfResponseIsEmptyItShouldReturnStatusNotFound(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockTodo := _mocks.NewMockTodoService(mockController)
	mockTodo.EXPECT().
		AddTodoListItem(requestTitle).
		Return(model.TodoItem{}, nil).Times(1)

	context, recorder := getAddTodoListItemContext(requestTitle)
	_ = TodoHandler{mockTodo}.AddTodoItem(context)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
}

func TestGivenRequestDataWhenIRequestToAddTodoListThenItShouldReturnAddedData(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockTodo := _mocks.NewMockTodoService(mockController)
	mockTodo.EXPECT().
		AddTodoListItem(requestTitle).
		Return(expectedTodoListItemResponseBody, nil).Times(1)

	context, recorder := getAddTodoListItemContext(requestTitle)

	_ = TodoHandler{mockTodo}.AddTodoItem(context)
	_ = json.Unmarshal(recorder.Body.Bytes(), &actualTodoListItemResponseBody)

	assert.Equal(t, expectedTodoListItemResponseBody, actualTodoListItemResponseBody)
}

func TestGivenRequestWhenIRequestToGetTodoListThenItShouldReturnAddedData(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockTodo := _mocks.NewMockTodoService(mockController)
	mockTodo.EXPECT().GetTodoList().Return(expectedTodoListResponseBody, nil).Times(1)
	context, recorder := getTodoListContext()

	_ = TodoHandler{mockTodo}.GetTodoList(context)
	_ = json.Unmarshal(recorder.Body.Bytes(), &actualTodoListResponseBody)

	assert.Equal(t, expectedTodoListResponseBody, actualTodoListResponseBody)
}

func TestGivenRequestWhenResponseDataEmptyThenItShouldBeReturnEmptyModel(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockTodo := _mocks.NewMockTodoService(mockController)
	mockTodo.EXPECT().GetTodoList().Return([]model.TodoList{}, nil).Times(1)
	context, recorder := getTodoListContext()

	_ = TodoHandler{mockTodo}.GetTodoList(context)
	_ = json.Unmarshal(recorder.Body.Bytes(), &actualTodoListResponseBody)

	assert.Equal(t, expectedTodoListEmptyResponse, actualTodoListResponseBody)
}

func TestGivenRequestWhenGetTodoListHasErrorThenItShouldBeReturnError(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockTodo := _mocks.NewMockTodoService(mockController)
	mockTodo.EXPECT().
		GetTodoList().
		Return([]model.TodoList{}, errors.New("todo repository has a error")).Times(1)

	context, _ := getTodoListContext()

	err := TodoHandler{mockTodo}.GetTodoList(context)

	assert.Error(t, err)
}

func TestGivenUpdateDataWhenIRequestToUpdateItemCompletedThenICanSeeUpdatedCompleteValue(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockTodoService := _mocks.NewMockTodoService(mockController)
	mockTodoService.EXPECT().
		UpdateItemCompleted(gomock.Any(), gomock.Any()).
		Return(expectedUpdateItemCompletedResponseBody, nil).Times(1)

	context, recorder := getUpdateTodoItemContext("1", "true")

	_ = TodoHandler{mockTodoService}.UpdateTodoItemCompleted(context)
	_ = json.Unmarshal(recorder.Body.Bytes(), &actualTodoListItemResponseBody)

	assert.Equal(t, expectedUpdateItemCompletedResponseBody, actualTodoListItemResponseBody)
}

func TestGivenRequestWhenUpdateItemHasErrorThenItShouldBeReturnError(t *testing.T) {
	mockController := gomock.NewController(t)

	mockTodo := _mocks.NewMockTodoService(mockController)
	mockTodo.EXPECT().
		UpdateItemCompleted(gomock.Any(), gomock.Any()).
		Return(model.TodoItem{}, errors.New("todo repository has a error")).Times(1)

	context, _ := getUpdateTodoItemContext("", "")

	err := TodoHandler{mockTodo}.UpdateTodoItemCompleted(context)

	assert.Error(t, err)
}

func TestGivenDeleteDataWhenIRequestToDeleteTodoItemThenICanSeeDeletedObject(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockTodoService := _mocks.NewMockTodoService(mockController)
	mockTodoService.EXPECT().
		DeleteTodoItem(gomock.Any()).
		Return(expectedUpdateItemCompletedResponseBody, nil).Times(1)

	context, recorder := getDeleteTodoItemContext("1")

	_ = TodoHandler{mockTodoService}.DeleteTodoItem(context)
	_ = json.Unmarshal(recorder.Body.Bytes(), &actualTodoListItemResponseBody)

	assert.Equal(t, expectedUpdateItemCompletedResponseBody, actualTodoListItemResponseBody)
}

func TestGivenDeleteDataWhenIRequestToDeleteTodoItemAndHasErrorThenICanSeeError(t *testing.T) {
	mockController := gomock.NewController(t)

	mockTodo := _mocks.NewMockTodoService(mockController)
	mockTodo.EXPECT().
		UpdateItemCompleted(gomock.Any(), gomock.Any()).
		Return(model.TodoItem{}, errors.New("todo service has a error")).Times(1)

	context, _ := getDeleteTodoItemContext("")

	err := TodoHandler{mockTodo}.DeleteTodoItem(context)

	assert.Error(t, err)
}
