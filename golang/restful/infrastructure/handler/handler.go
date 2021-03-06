package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"restful/domain"
	"restful/usecase"

	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
)

type Handlers struct {
	uc       usecase.CustomerInputBoundary
	validate *validator.Validate
}

func NewHandlers(uc usecase.CustomerInputBoundary) *Handlers {
	return &Handlers{uc, validator.New()}
}

func (h *Handlers) Create(w http.ResponseWriter, r *http.Request) {
	var cid usecase.CustomerInputData
	if err := json.NewDecoder(r.Body).Decode(&cid); err != nil {
		log.Printf("error: %v %+v", r.URL, err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	if err := h.validate.Struct(&cid); err != nil {
		log.Printf("error: %v %+v", r.URL, err)
		http.Error(w, "validation error", http.StatusBadRequest)
		return
	}
	customer, err := h.uc.Create(cid)
	if err != nil {
		log.Printf("error: %v %+v", r.URL, err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	Response(w, http.StatusOK, *customer)
}

func (h *Handlers) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["id"])
	customer, err := h.uc.Read(vars["id"])
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	Response(w, http.StatusOK, *customer)
}

func (h *Handlers) Update(w http.ResponseWriter, r *http.Request) {
	var customer domain.Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		log.Printf("error: %v %+v", r.URL, err)
		http.Error(w, "fail to decode", http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)

	updatedCustomer, err := h.uc.Update(vars["id"], &customer)
	if err != nil {
		log.Printf("error: %v %+v", r.URL, err)
		http.Error(w, "fail to update", http.StatusInternalServerError)
		return
	}
	Response(w, http.StatusOK, updatedCustomer)
}

func (h *Handlers) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	err := h.uc.Delete(vars["id"])
	if err != nil {
		log.Println(err)
		http.Error(w, "fail to delete", http.StatusInternalServerError)
		return
	}
	Response(w, http.StatusOK, "deleted")
}

func Response(w http.ResponseWriter, statusCode int, message interface{}) {
	b, _ := json.Marshal(message)
	w.WriteHeader(statusCode)
	fmt.Fprintln(w, string(b))
}
