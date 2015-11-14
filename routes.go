package main

import (
	"net/http"
	v1todos "simple-todo-api/controllers/v1/todos"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		v1todos.Index,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		v1todos.Show,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		v1todos.Create,
	},
}
