package handlers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

// State defines the mongo document structure
type State struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	Name     string        `bson:"name" json:"state"`
	Location GeoJSON       `bson:"location" json:"border"`
}

//GeoJSON holds the longitude & latitude data to query from
type GeoJSON struct {
	Type        string      `bson:"type" json:"type"`
	Coordinates [][]float32 `bson:"coordinates" json:"coordinates"`
}

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
