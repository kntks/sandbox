package server

import (
	"log"
	"net/http"
	"time"

	"sandbox/auth"
	myauth "sandbox/auth"

	"github.com/gin-gonic/gin"
)

type Name struct {
	First     string `json:"first"`
	Last      string `json:"last"`
	FirstFuri string `json:"firstFuri"`
	LastFuri  string `json:"lastFuri"`
}

type User struct {
	Name    Name   `json:"name"`
	Age     uint   `json:"age"`
	Address string `json:"address"`
}

func Public(c *gin.Context) {
	c.JSON(http.StatusOK, User{
		Name:    Name{},
		Age:     1,
		Address: "Japan",
	})
}

func GetToken(c *gin.Context) {
	token, err := auth.GenerateToken()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func GetUser(c *gin.Context) {
	auth := myauth.NewAuth(c.Request)

	if err := auth.Valid(); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "No"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func GetHandler() http.Handler {
	r := gin.Default()
	r.GET("/ping", Public)
	r.GET("/token", GetToken)
	r.GET("/user", GetUser)
	return r
}

func Start() {
	server := &http.Server{
		Addr:              ":8080",
		Handler:           GetHandler(),
		ReadTimeout:       time.Minute,
		ReadHeaderTimeout: time.Minute,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
