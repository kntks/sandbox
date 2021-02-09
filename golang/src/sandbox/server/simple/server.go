package simple

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// このサーバはexponential backoffを検証するためだけのシンプルなサーバ
// 基本的にはhttp status200, 300, 400, 500を返す

type SimpleServer struct{}

func ok(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{result: ok}`))
}

func badRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(`{result: badrequest}`))
}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(`{result: internal server error}`))
}

func random(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	x := rand.Intn(10)
	w.Write([]byte(strconv.Itoa(x)))
}

func GetHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/200", ok)
	mux.HandleFunc("/4xx", badRequest)
	mux.HandleFunc("/5xx", internalServerError)
	mux.HandleFunc("/random", random)
	return mux
}

func (SimpleServer) Start() {
	server := &http.Server{
		Addr:              ":8081",
		Handler:           GetHandler(),
		ReadTimeout:       time.Minute,
		ReadHeaderTimeout: time.Minute,
	}
	fmt.Println("Server Listen port: 8081")
	log.Fatal(server.ListenAndServe())
}
