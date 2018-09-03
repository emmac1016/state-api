package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/emmac1016/state-api/internal"
	"github.com/julienschmidt/httprouter"
)

type StatesResponse struct {
	States []internal.State `json:"states"`
}

// GetState returns state where a given longitude and latitude reside in
func GetState(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var message string
	if err := r.ParseForm(); err != nil {
		log.Print("ParseForm() err: ", err)
		fail(nil, w)
		return
	}

	long := r.FormValue("longitude")
	lat := r.FormValue("latitude")
	if len(long) == 0 || len(lat) == 0 {
		log.Print("latitude or longitude is missing")
		message = `{message: "'longitude' and 'latitude' fields are required"}`
		fail([]byte(message), w, http.StatusBadRequest)
		return
	}

	fltLong, err := strconv.ParseFloat(long, 64)
	if err != nil {
		log.Print("invalid longitude value")
		message = `{message: "invalid longitude value"}`
		fail([]byte(message), w, http.StatusBadRequest)
		return
	}
	fltLat, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		log.Print("invalid latitude value")
		message = `{message: "invalid latitude value"}`
		fail([]byte(message), w, http.StatusBadRequest)
		return
	}

	sr, err := internal.NewStateRepo()
	states, err := sr.FindStateByCoordinates(fltLong, fltLat)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(states); err != nil {
		log.Print("failed to encode json")
		message = `{message: "failure to render response"}`
		fail([]byte(message), w, http.StatusBadRequest)
		return
	}
}
