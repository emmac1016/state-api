package main

import (
	"log"
	"net/http"

	"github.com/emmac1016/state-api/handlers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.POST("/", handlers.GetState)

	log.Fatal(http.ListenAndServe(":8080", router))
}
