package main

import (
	"log"
	"net/http"
	"untitled/controller"
	"untitled/util"

	"github.com/rs/cors"
)

func main() {
	defer util.CloseDB()
	mux := http.NewServeMux()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "DELETE"},
	})

	// mux.HandleFunc("/", controller.TestHandler)
	mux.HandleFunc("/api/todos", controller.TodosHandler)
	mux.HandleFunc("/api/todo/", controller.TodoHandler)

	handler := c.Handler(mux)

	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatal(err)
	}
}
