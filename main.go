package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/sample-api/driver"

	studentHandler "example.com/sample-api/handlers/student"
	studentService "example.com/sample-api/services/student"
	studentStore "example.com/sample-api/stores/student"

	"github.com/gorilla/mux"
)

func main() {
	const PORT = 8000
	cfg := driver.Configs{
		Host:     "localhost",
		Username: "root",
		Password: "password",
		Port:     3306,
		Database: "college",
	}

	db, err := driver.InitializeDB(cfg)
	if err != nil {
		log.Fatalf("cannot connect to database: %v\n", err)
	}

	defer db.Close()

	log.Printf("connected to mysql, %s:%d ; database: %s\n", cfg.Host, cfg.Port, cfg.Database)

	r := mux.NewRouter()

	store := studentStore.New(db)
	service := studentService.New(store)
	h := studentHandler.New(service)

	r.HandleFunc("/students", h.GetAll).Methods(http.MethodGet)         // get all student records
	r.HandleFunc("/students", h.Create).Methods(http.MethodPost)        // create a new record
	r.HandleFunc("/students/{id}", h.Get).Methods(http.MethodGet)       // get student by ID
	r.HandleFunc("/students/{id}", h.Update).Methods(http.MethodPut)    // update student by ID
	r.HandleFunc("/students/{id}", h.Delete).Methods(http.MethodDelete) // delete student by ID

	//r.HandleFunc("/employees").Methods(http.MethodGet)         // get all student records
	//r.HandleFunc("/employees").Methods(http.MethodPost)        // create a new record
	//r.HandleFunc("/employees/{id}").Methods(http.MethodGet)    // get student by ID
	//r.HandleFunc("/employees/{id}").Methods(http.MethodPut)    // update student by ID
	//r.HandleFunc("/employees/{id}").Methods(http.MethodDelete) // delete student by ID

	// todo: middleware

	srv := http.Server{
		Addr:    fmt.Sprintf("localhost:%d", PORT),
		Handler: r,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf("error in starting server: %v\n", err)
	}
}
