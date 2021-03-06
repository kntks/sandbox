package infrastructure

import (
	"fmt"
	"log"
	"net/http"
	"restful/infrastructure/handler"
	"restful/infrastructure/repository"
	"restful/usecase"
	"time"

	"github.com/gorilla/mux"
)

func getCustomerUsecases() *usecase.CustomerInteractor {
	repo, err := repository.NewCustomerRepository("redis:6379")
	if err != nil {
		log.Fatal(err)
	}

	return usecase.NewCustomerInteractor(repo)
}

func router() *mux.Router {
	handlers := handler.NewHandlers(getCustomerUsecases())
	r := mux.NewRouter()

	r.HandleFunc("/create", handlers.Create).Methods(http.MethodPost)
	r.HandleFunc("/customer/{id}", handlers.Read).Methods(http.MethodGet)
	r.HandleFunc("/customer/{id}", handlers.Update).Methods(http.MethodPatch)
	r.HandleFunc("/customer/{id}", handlers.Delete).Methods(http.MethodDelete)

	return r
}

func Run() {
	srv := &http.Server{
		Handler:           router(),
		Addr:              ":8000",
		ReadHeaderTimeout: 15 * time.Second,
		WriteTimeout:      15 * time.Second,
	}

	fmt.Println("Server Listen port :8000")
	log.Fatal(srv.ListenAndServe())
}
