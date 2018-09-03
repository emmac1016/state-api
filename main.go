package main

import (
	"log"
	"net/http"
	"os"

	"github.com/emmac1016/state-api/handlers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.POST("/", handlers.GetState)

	log.Fatal(http.ListenAndServe(os.Getenv("APP_PORT"), router))
}
