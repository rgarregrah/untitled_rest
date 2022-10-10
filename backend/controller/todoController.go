package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	todos "untitled/model"
	"untitled/util"
)

// Todo単体に関するハンドラ
func TodoHandler(w http.ResponseWriter, r *http.Request) {
	params := util.GetPathParams(r)

	// ["api", "todo", "1"]
	if len(params) != 3 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	id, err := strconv.Atoi(params[2])
	if err != nil || id < 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		getTodo(id, w, r)
	case "POST":
		upsertTodo(id, w, r)
	case "DELETE":
		deleteTodo(id, w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

// Todo全体に関するハンドラ
func TodosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getTodos(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := todos.FindAll()
	if err != nil {
		log.Print("[ERROR]", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func getTodo(id int, w http.ResponseWriter, r *http.Request) {
	todo, err := todos.Find(id)
	if err != nil {
		log.Print("[ERROR]", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusMethodNotAllowed)
		return
	}

	// なかった場合はゼロ値を返すので、404で返す
	if todo.Id == 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func upsertTodo(id int, w http.ResponseWriter, r *http.Request) {
	type reqStruct struct {
		Id     int    `json:"id"`
		Body   string `json:"body"`
		Status string `json:"status"`
	}
	req := reqStruct{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	err = todos.Upsert(req.Id, req.Body, req.Status)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func deleteTodo(id int, w http.ResponseWriter, r *http.Request) {
	err := todos.Delete(id)
	if err != nil {
		log.Print("[ERROR]", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
}
