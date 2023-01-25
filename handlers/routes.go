package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/api/todos", app.getAllTodos)
	router.HandlerFunc(http.MethodPost, "/api/todos", app.editTodo)
	router.HandlerFunc(http.MethodGet, "/api/todos/:id", app.getOneTodo)
	router.HandlerFunc(http.MethodDelete, "/api/todos/:id", app.deleteTodo)

	return router
}
