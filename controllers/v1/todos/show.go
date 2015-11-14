package todos

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	todoModels "simple-todo-api/models/todo"
	"simple-todo-api/modules/apperror"
	"strconv"
)

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var todoId int
	var err error
	if todoId, err = strconv.Atoi(vars["todoId"]); err != nil {
		panic(err)
	}
	todo := todoModels.RepoFindTodo(todoId)
	if todo.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(todo); err != nil {
			panic(err)
		}
		return
	}

	// If we did not find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	jsonErr := apperror.JsonErr{Code: http.StatusNotFound, Text: "Not Found"}
	jsonErr.ToJSON(w)
}
