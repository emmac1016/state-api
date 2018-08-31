package repositories

import "gopkg.in/mgo.v2/bson"

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

// FindStateByCoordinates finds what state the given coordinates are in
func FindStateByCoordinates(longitude float32, latitude float32) (*State, error) {
	return nil, nil
}
