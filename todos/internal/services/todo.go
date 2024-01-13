package services

import (
	"errors"
	db "todos/internal/models"
)


type TodoService struct {}

func (ts *TodoService) GetTodo(id uint32) (*db.Todo, error) {
	var todo db.Todo
	if err := db.DB.First(&todo, id).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}

// TODO: add pagination
func (ts *TodoService) GetAllTodos() ([]db.Todo, error) {
	var todos []db.Todo
	if err :=	db.DB.Order("id ASC").Find(&todos).Error; err != nil {
		return nil, err
	}

	return todos, nil
}

func (ts *TodoService) CreateTodo(content string) (*db.Todo, error) {
	todo := &db.Todo{Content: content, IsComplete: false}
	if err := db.DB.Create(&todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

func (ts *TodoService) UpdateTodoContent(id uint32, content string) (*db.Todo, error) {
	todo := &db.Todo{ID: id}

	if err := db.DB.Model(todo).Updates(db.Todo{Content: content}).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

func (ts *TodoService) MarkTodoAsCompleted(id uint32) (*db.Todo, error) {
	return changeTodoCompleteness(id, true)
}

func (ts *TodoService) MarkTodoAsInCompleted(id uint32) (*db.Todo, error) {
	return changeTodoCompleteness(id, false)
}

func (ts *TodoService) DeleteTodo(id uint32) (*db.Todo, error) {
	todo, err := ts.GetTodo(id)
	if err != nil {
		return nil, err
	}

	if (db.DB.Delete(todo).RowsAffected == 0) {
		return nil, errors.New("")
	}

	return todo, nil
}

// private
func changeTodoCompleteness(id uint32, isComplete bool) (*db.Todo, error) {
	todo := &db.Todo{ID: id}

	if err := db.DB.Model(todo).Updates(db.Todo{IsComplete: isComplete}).Error; err != nil {
		return nil, err
	}

	return todo, nil
}
