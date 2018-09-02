package internal

import (
	"log"

	"gopkg.in/mgo.v2"
)

// DB holds database name and connection used to run queries
type DB struct {
	Name       string
	Connection *mgo.Session
}

// NewDB returns DB struct with connection and database name
func NewDB(ci *ConnectionInfo) (*DB, error) {
	conn, err := NewConnection(ci)

	if err != nil {
		log.Print("Error in creating DB instance: ", err)
		return nil, err
	}

	return &DB{
		Name:       ci.Database,
		Connection: conn,
	}, nil
}

// SetGeoSpatialIndex sets 2d geospatial index for a given collection,
// assumes GeoJSON format with location as the key
func (db *DB) SetGeoSpatialIndex(collectionName string) error {
	session := db.Connection.Copy()
	defer session.Close()

	collection := session.DB(db.Name).C(collectionName)
	index := mgo.Index{
		Key: []string{"$2dsphere:location"},
	}

	err := collection.EnsureIndex(index)

	return err
}
