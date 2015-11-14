package todos

import (
	"encoding/json"
	"net/http"
	todoModels "simple-todo-api/models/todo"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todoModels.GetAllTodos()); err != nil {
		panic(err)
	}
}
