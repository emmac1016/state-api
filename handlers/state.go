package handlers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetState(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
	longitude := r.FormValue("longitude")
	latitude := r.FormValue("latitude")

	fmt.Fprintf(w, "Longitude: %s, Latitude: %s\n", longitude, latitude)
}
