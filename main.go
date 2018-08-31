package main

import (
	"log"
	"net/http"

	"github.com/emmac1016/state-api/handlers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// if os.Getenv("ENV") == "DEV" {
	// 	log.Print("Loading fixture data")
	// 	err := internal.LoadFixtureData()
	// 	if err != nil {
	// 		log.Fatal("Failure to supply data")
	// 	}
	// }

	router := httprouter.New()
	router.POST("/", handlers.GetState)

	log.Fatal(http.ListenAndServe(":8080", router))
}
