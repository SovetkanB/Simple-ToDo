package service

import (
	"errors"
	"log/slog"
	"simple-todo/internal/models"
	"simple-todo/internal/repo"
	"time"
)

type ToDoService interface {
	CreateToDo(todo *models.ToDo) error
	GetAllTodos() ([]models.ToDo, error)
	UpdateTodo(todo *models.ToDo) error
	DeleteTodo(id string) error
	//CompleteTodo(id string) error
}

type TodoServiceImpl struct {
	todoRepo repo.ToDoRepo
}

func NewToDoService(todoRepo repo.ToDoRepo) *TodoServiceImpl {
	return &TodoServiceImpl{
		todoRepo: todoRepo,
	}
}

func (s *TodoServiceImpl) CreateToDo(todo *models.ToDo) error {
	if todo.Name == "" {
		return errors.New("name cannot be empty")
	}

	todo.IsCompleted = false
	todo.CreatedAt = time.Now().Format(time.RFC3339)

	if err := s.todoRepo.Create(todo); err != nil {
		slog.Error("Failed to create todo", "error", err, "todoId", todo.ID)
		return err
	}

	slog.Info("Todo created", "todoId", todo.ID)
	return nil
}

func (s *TodoServiceImpl) GetAllTodos() ([]models.ToDo, error) {
	todos, err := s.todoRepo.GetAll()
	if err != nil {
		slog.Error("Failed to get all todos", "error", err.Error())
		return nil, err
	}

	return todos, nil
}

func (s *TodoServiceImpl) UpdateTodo(todo *models.ToDo) error {
	if err := s.todoRepo.Update(todo); err != nil {
		slog.Error("Failed to update todo", "error", err, "todoId", todo.ID)
		return err
	}

	slog.Info("Todo updated", "todoId", todo.ID)
	return nil
}

func (s *TodoServiceImpl) DeleteTodo(id string) error {
	if err := s.todoRepo.Delete(id); err != nil {
		slog.Error("Failed to delete todo", "error", err, "todoId", id)
		return err
	}

	slog.Info("Todo deleted", "todoId", id)
	return nil
}

// func (s *TodoServiceImpl) CompleteTodo(id string) error {
// 	todos, err := s.todoRepo.GetAll()
// 	if err != nil {
// 		return err
// 	}

// 	found := false
// 	for _, t := range todos {
// 		if t.ID == id {
// 			t.IsCompleted = true
// 			found = true
// 			break
// 		}
// 	}

// 	if !found {
// 		return errors.New("Todo not found")
// 	}

// 	return s.todoRepo.Update()
// }
