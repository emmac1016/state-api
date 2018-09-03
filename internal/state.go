package internal

import (
	"encoding/json"
	"log"

	"gopkg.in/mgo.v2/bson"
)

// AbstractStateRepo defines the methods that StateRepo implements
type AbstractStateRepo interface {
	FindStateByCoordinates(float64, float64)
}

// StateRepo is a wrapper around DB struct to expand functionality
type StateRepo struct {
	dbh Database
}

// State defines the mongo document structure
type State struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Name     string        `bson:"name" json:"state"`
	Location GeoJSON       `bson:"location" json:"omitempty"`
}

//GeoJSON holds the longitude & latitude data to query from
type GeoJSON struct {
	Type        string        `bson:"type"`
	Coordinates [][][]float64 `bson:"coordinates"`
}

// NewStateRepo returns an instance of StateRepo that will be used to execute queries
func NewStateRepo() (*StateRepo, error) {
	connInfo := GetDefaultConnection()
	dbh, err := NewDBHandler(connInfo)
	if err != nil {
		log.Print("Failure to get StateRepo instance: ", err)
		return nil, err
	}

	return &StateRepo{dbh: dbh}, nil
}

// FindStateByCoordinates finds what state the given coordinates are in
func (sr *StateRepo) FindStateByCoordinates(longitude float64, latitude float64) ([]State, error) {
	var results []State

	err := sr.dbh.Find("states", bson.M{
		"location": bson.M{
			"$geoIntersects": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{longitude, latitude},
				},
			},
		},
	}).Select(bson.M{"name": 1, "_id": 0}).All(&results)

	if len(results) == 0 {
		return make([]State, 0), err
	}

	return results, err
}

// MarshalJSON returns only the name of the state from the struct
func (s *State) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Name)
}
