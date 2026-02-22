package repo

import (
	"encoding/json"
	"errors"
	"os"
	"simple-todo/internal/models"
	"sync"
)

type ToDoRepo interface {
	Create(todo *models.ToDo) error
	GetAll() ([]models.ToDo, error)
	Update(todo *models.ToDo) error
	Delete(id string) error
}

type JsonToDoRepo struct {
	todoFile string
	mu       sync.RWMutex
}

func NewJSONTodoRepo(todoFile string) *JsonToDoRepo {
	return &JsonToDoRepo{
		todoFile: todoFile,
	}
}

func (r *JsonToDoRepo) readTodos() ([]models.ToDo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	data, err := os.ReadFile(r.todoFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.ToDo{}, nil
		}

		return nil, err
	}

	if len(data) == 0 {
		return []models.ToDo{}, nil
	}

	var todos []models.ToDo
	if err := json.Unmarshal(data, &todos); err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *JsonToDoRepo) writeTodos(todos []models.ToDo) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	data, err := json.MarshalIndent(todos, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.todoFile, data, 0644)
}

func (r *JsonToDoRepo) Create(todo *models.ToDo) error {
	todos, err := r.readTodos()
	if err != nil {
		return err
	}

	for _, t := range todos {
		if t.ID == todo.ID {
			return errors.New("todo with this Id already exists")
		}
	}

	todos = append(todos, *todo)
	return r.writeTodos(todos)
}

func (r *JsonToDoRepo) GetAll() ([]models.ToDo, error) {
	return r.readTodos()
}

func (r *JsonToDoRepo) Update(todo *models.ToDo) error {
	todos, err := r.readTodos()
	if err != nil {
		return err
	}

	found := false
	for i, t := range todos {
		if t.ID == todo.ID {
			todos[i] = *todo
			found = true
			break
		}
	}

	if !found {
		return errors.New("todo not found")
	}

	return r.writeTodos(todos)
}

func (r *JsonToDoRepo) Delete(id string) error {
	todos, err := r.readTodos()
	if err != nil {
		return err
	}

	found := false
	newTodos := []models.ToDo{}
	for _, t := range todos {
		if t.ID != id {
			newTodos = append(newTodos, t)
		} else {
			found = true
		}
	}

	if !found {
		return errors.New("todo not found")
	}

	return r.writeTodos(newTodos)
}
