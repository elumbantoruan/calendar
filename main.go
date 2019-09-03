package main

import (
	"calendar/handlers"
	"calendar/service"
	"calendar/storage"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("main")
	m, err := registerHandlers()
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", m)

	err = http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func registerHandlers() (*mux.Router, error) {
	m := mux.NewRouter()

	mem := storage.NewMemoryStorage()
	cal := service.NewCalendar(mem)
	handler := handlers.NewHandleCalendar(cal)

	m.HandleFunc("/calendar", handler.AddEvent).Methods("POST")

	return m, nil
}
