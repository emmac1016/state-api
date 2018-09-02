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
	db *DB
}

// State defines the mongo document structure
type State struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"omitempty"`
	Name     string        `bson:"name" json:"state"`
	Location GeoJSON       `bson:"location" json:"omitempty"`
}

//GeoJSON holds the longitude & latitude data to query from
type GeoJSON struct {
	Type        string        `bson:"type" json:"omitempty"`
	Coordinates [][][]float64 `bson:"coordinates" json:"omitempty"`
}

// NewStateRepo returns an instance of StateRepo that will be used to execute queries
func NewStateRepo() (*StateRepo, error) {
	connInfo := GetDefaultConnection()
	db, err := NewDB(connInfo)
	if err != nil {
		log.Print("Failure to get StateRepo instance: ", err)
		return nil, err
	}

	return &StateRepo{db: db}, nil
}

// FindStateByCoordinates finds what state the given coordinates are in
func (sr *StateRepo) FindStateByCoordinates(longitude float64, latitude float64) ([]State, error) {
	var results []State
	session := sr.db.Connection.Copy()
	defer session.Close()

	collection := session.DB(sr.db.Name).C("states")
	err := collection.Find(bson.M{
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
