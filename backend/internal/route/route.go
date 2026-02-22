package route

import (
	"net/http"
	"simple-todo/internal/handler"
	"simple-todo/internal/repo"
	"simple-todo/internal/service"
)

func CreateRoute() http.Handler {
	mux := http.NewServeMux()

	todoRepo := repo.NewJSONTodoRepo("data/todos.json")

	todoService := service.NewToDoService(todoRepo)

	todoHandler := handler.NewToDoHandler(todoService)

	mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			todoHandler.CreateTodo(w, r)
		case http.MethodGet:
			todoHandler.GetAllTodos(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return mux
}
