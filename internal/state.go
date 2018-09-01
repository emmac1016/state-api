package internal

import (
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
	ID       bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	Name     string        `bson:"name" json:"state"`
	Location GeoJSON       `bson:"location" json:"border"`
}

//GeoJSON holds the longitude & latitude data to query from
type GeoJSON struct {
	Type        string        `bson:"type" json:"type"`
	Coordinates [][][]float64 `bson:"coordinates" json:"coordinates"`
}

// NewStateRepo returns an instance of StateRepo that will be used to execute queries
func NewStateRepo() (*StateRepo, error) {
	db, err := NewAppDB()
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
	}).All(&results)

	return results, err
}
