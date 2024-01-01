package service

import (
	"fmt"
	"testing"
	"time"
	"todocible_api/dto"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var todoService TodoService

func TestMain(m *testing.M) {
	todoService = NewTodoService()

	m.Run()
}

func TestSuccessCreate(t *testing.T) {
	todo, err := todoService.Create(dto.TodoRequest{
		Title:       "Coba",
		Description: "Hello World",
	})

	fmt.Println(todo.Id)
	fmt.Println(todo.Title)
	fmt.Println(todo.Description)
	fmt.Println(todo.DueDate)
	fmt.Println(todo.Completed)

	assert.Nil(t, err)
}

func TestCreateWithoutTitle(t *testing.T) {
	_, err := todoService.Create(dto.TodoRequest{})

	assert.NotNil(t, err)
}

func TestGetAll(t *testing.T) {
	todos, err := todoService.GetAll()

	for _, todo := range todos {
		fmt.Println("===================")
		fmt.Println(todo.Id)
		fmt.Println(todo.Title)
		fmt.Println(todo.Description)
		fmt.Println(todo.DueDate)
		fmt.Println(todo.Completed)
	}

	assert.Nil(t, err)
}

func TestGetWithValidId(t *testing.T) {
	newTodo, err := todoService.Create(dto.TodoRequest{
		Title:       "Coba",
		Description: "Hello World",
		DueDate:     time.Now(),
	})

	assert.Nil(t, err)

	todo, err := todoService.Get(newTodo.Id)

	assert.Nil(t, err)

	fmt.Println("===================")
	fmt.Println(todo.Id)
	fmt.Println(todo.Title)
	fmt.Println(todo.Description)
	fmt.Println(todo.DueDate)
	fmt.Println(todo.Completed)
}

func TestGetWithInvalidId(t *testing.T) {
	_, err := todoService.Get(uuid.New().String())

	assert.NotNil(t, err)

	fmt.Println(err.Error())
}

func TestUpdate(t *testing.T) {
	newTodo, err := todoService.Create(dto.TodoRequest{
		Title:       "Coba",
		Description: "Hello World",
		DueDate:     time.Now(),
	})

	assert.Nil(t, err)
	assert.Equal(t, "Coba", newTodo.Title)

	todo, err := todoService.Update(newTodo.Id, dto.TodoRequest{
		Title: "YO",
	})

	assert.Nil(t, err)
	assert.Equal(t, "YO", todo.Title)
}

// func TestCompleted(t *testing.T) {
// 	todos := todoService.GetAll()
// 	todo, _ := todoService.Get(todos[0].Id)
// 	assert.Equal(t, false, todo.Completed)

// 	todoService.SetCompleted(todo.Id, true)

// 	todo, _ = todoService.Get(todo.Id)
// 	assert.Equal(t, true, todo.Completed)
// }

// func TestDelete(t *testing.T) {
// 	todos := todoService.GetAll()
// 	todo, _ := todoService.Get(todos[0].Id)
// 	assert.NotNil(t, todo)

// 	todoService.Delete(todo.Id)

// 	todo, _ = todoService.Get(todo.Id)
// 	assert.Nil(t, todo)
// }
