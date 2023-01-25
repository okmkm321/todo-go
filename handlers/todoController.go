package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"
	"todo-lesson/models"

	"github.com/julienschmidt/httprouter"
)

type TodoPayload struct {
	ID          int    `json:"id"`
	Title       string `json:"name"`
	Description string `json:"description"`
}

func (app *Application) getAllTodos(w http.ResponseWriter, r *http.Request) {
	ctg, err := app.Models.DB.TodoGetAll()
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}

	err = app.WriteJSON(w, http.StatusOK, ctg, "todos")
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}
}

func (app *Application) getOneTodo(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.Logger.Print(errors.New("invalid id parameter"))
		app.ErrorJSON(w, err)
		return
	}

	ctg, err := app.Models.DB.GetTodo(id)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}
	err = app.WriteJSON(w, http.StatusOK, ctg, "todo")
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}
}

func (app *Application) editTodo(w http.ResponseWriter, r *http.Request) {
	var tp TodoPayload
	err := json.NewDecoder(r.Body).Decode(&tp)
	if err != nil {
		log.Println(err)
		app.ErrorJSON(w, err)
		return
	}

	var tm models.Todo

	if tp.ID != 0 {
		id := tp.ID
		t, _ := app.Models.DB.GetTodo(id)
		tm = *t
		tm.UpdatedAt = time.Now()
	}

	tm.ID = tp.ID
	tm.Title = tp.Title
	tm.Description = tp.Description
	tm.UpdatedAt = time.Now()

	if tp.ID == 0 {
		err = app.Models.DB.TodoCreate(tm)
		if err != nil {
			app.ErrorJSON(w, err)
			return
		}
	} else {
		err = app.Models.DB.TodoUpdate(tm)
		if err != nil {
			app.ErrorJSON(w, err)
			return
		}
	}

	ok := JsonResp{
		OK: true,
	}

	err = app.WriteJSON(w, http.StatusOK, ok, "response")
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}
}

func (app *Application) deleteTodo(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}

	err = app.Models.DB.TodoDelete(id)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}

	ok := JsonResp{
		OK: true,
	}

	err = app.WriteJSON(w, http.StatusOK, ok, "response")
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}
}
