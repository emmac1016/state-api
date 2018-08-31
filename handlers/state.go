package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/emmac1016/state-api/internal"
	"github.com/julienschmidt/httprouter"
)

func GetState(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
	longitude, err := strconv.ParseFloat(r.FormValue("longitude"), 64)
	if err != nil {
		log.Print("invalid longitude value")
		return //return 4xx bad request
	}
	latitude, err := strconv.ParseFloat(r.FormValue("latitude"), 64)
	if err != nil {
		log.Print("invalid latitude value")
		return //return 4xx bad request
	}

	fmt.Fprintf(w, "Longitude: %s, Latitude: %s\n", longitude, latitude)

	sr, err := internal.NewStateRepo()
	state, err := sr.FindStateByCoordinates(longitude, latitude)

	fmt.Println(state)
}
