package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"simple-todo/helpers"
	"simple-todo/internal/models"
	"simple-todo/internal/service"
)

type TodoHandler struct {
	todoService service.ToDoService
}

func NewToDoHandler(todoService service.ToDoService) *TodoHandler {
	return &TodoHandler{
		todoService: todoService,
	}
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.ToDo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		slog.Error("Failed to decode todo request", "error", err.Error())
		helpers.ResponseError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.todoService.CreateToDo(&todo); err != nil {
		helpers.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.ResponseJson(w, http.StatusCreated, todo)
}

func (h *TodoHandler) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.todoService.GetAllTodos()
	if err != nil {
		helpers.ResponseError(w, http.StatusInternalServerError, "Failed to retrieve todos")
		return
	}

	helpers.ResponseJson(w, http.StatusOK, todos)
}

func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
}
