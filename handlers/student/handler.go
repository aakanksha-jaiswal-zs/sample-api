package student

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"example.com/sample-api/errors"
	"example.com/sample-api/models"
	"example.com/sample-api/services"
)

type handler struct {
	// inject the service layer dependency
	service services.Student
}

func New(s services.Student) handler {
	return handler{service: s}
}

func (h handler) GetAll(w http.ResponseWriter, r *http.Request) {
	students, err := h.service.GetAll()
	if err != nil {
		// handle this error
		setStatus(w, err)
		return
	}

	resp, err := json.Marshal(students)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// todo set the content type--->w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	_, err = w.Write(resp)
	if err != nil {
		log.Println("error in writing response")
	}
}

func (h handler) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]
	if strings.TrimSpace(id) == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	studentID, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	student, err := h.service.Get(int64(studentID))
	if err != nil {
		// handle this error
		setStatus(w, err)
		return
	}

	resp, err := json.Marshal(student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK) // todo set status first, write response after that

	_, err = w.Write(resp)
	if err != nil {
		log.Println("error in writing response")
	}
}

func setStatus(w http.ResponseWriter, err error) {
	log.Println(err)

	switch err.(type) {
	case errors.EntityNotFound:
		w.WriteHeader(http.StatusNotFound)
	case errors.EntityAlreadyExists, errors.InvalidParam, errors.MissingParam: // todo statusOK
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h handler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var student models.Student

	err = json.Unmarshal(body, &student)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	student, err = h.service.Create(student)
	if err != nil {
		// handle this error
		setStatus(w, err)
		return
	}

	resp, err := json.Marshal(student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	_, err = w.Write(resp)
	if err != nil {
		log.Println("error in writing response")
	}
}

func (h handler) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]
	if strings.TrimSpace(id) == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	studentID, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var student models.Student

	err = json.Unmarshal(body, &student)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	student.ID = int64(studentID)

	student, err = h.service.Update(student)
	if err != nil {
		// handle this error
		setStatus(w, err)
		return
	}

	resp, err := json.Marshal(student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	_, err = w.Write(resp)
	if err != nil {
		log.Println("error in writing response")
	}
}

func (h handler) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]
	if strings.TrimSpace(id) == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	studentID, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.service.Delete(int64(studentID))
	if err != nil {
		// handle this error
		setStatus(w, err)
		return
	}

	w.WriteHeader(http.StatusOK) //todo status no content
}
