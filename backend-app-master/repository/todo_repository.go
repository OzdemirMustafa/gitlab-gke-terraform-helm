package repository

import (
	"backend-app/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type Repository struct {
	gorm *gorm.DB
}

func (r Repository) AddTodoListItemByTitle(title string) (model.TodoItem, error) {
	todoItem := model.TodoItem{Title: title, Completed: false}

	response := r.gorm.Create(&todoItem)

	return model.TodoItem{ID: todoItem.ID, Title: todoItem.Title, Completed: todoItem.Completed}, response.Error
}

func (r Repository) GetTodoListItems() ([]model.TodoList, error) {
	todoList := make([]model.TodoList, 0)
	response := r.gorm.Find(&todoList)

	return todoList, response.Error
}
func (r Repository) UpdateTodoListItem(id int, completed bool) (model.TodoItem, error) {
	todoItem := model.TodoItem{ID: id, Completed: completed}
	response := r.gorm.Model(&todoItem).Update("completed", completed)

	return todoItem, response.Error
}

func (r Repository) DeleteTodoItemByID(id int) (model.TodoItem, error) {
	todoItem := model.TodoItem{ID: id}
	response := r.gorm.Delete(todoItem)

	return todoItem, response.Error
}

func (r Repository) ConnectToDatabase() *Repository {
	dsn := os.Getenv("DB_USERNAME") + ":" +
		   os.Getenv("DB_PASSWORD") + "@tcp(" +
		   os.Getenv("DB_URL") + ")/" +
		   os.Getenv("DB_NAME")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		return nil
	}

	repository := &Repository{db}

	return repository
}
