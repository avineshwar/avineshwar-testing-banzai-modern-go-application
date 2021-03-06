/*
 * Todo API
 *
 * Manage TODOs
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		"GET",
		"/",
		Index,
	},

	{
		"CreateTodo",
		strings.ToUpper("Post"),
		"/todos",
		CreateTodo,
	},

	{
		"ListTodos",
		strings.ToUpper("Get"),
		"/todos",
		ListTodos,
	},

	{
		"MarkAsDone",
		strings.ToUpper("Patch"),
		"/todos/{id}/done",
		MarkAsDone,
	},
}
