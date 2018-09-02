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

type Query interface {
	All(interface{}) error
	Select(interface{}) *mgo.Query
}

type Collection interface {
	EnsureIndex(mgo.Index) error
	Find(interface{}) *mgo.Query
}

type Session interface {
	Close()
	Copy() *mgo.Session
	DB(name string) *mgo.Database
}

type Database interface {
	C(name string) *Collection
}

type DBHandler struct {
	DB      string
	Session Session
}

type DBH interface {
	Collection(string) Collection
	Find(collection string, query interface{}) Query
	SetGeoSpatialIndex(string) error
}

func NewDBHandler(ci *ConnectionInfo) (*DBHandler, error) {
	session, err := ci.Dial()
	if err != nil {
		log.Print("Error in creating DBHandler: ", err)
		return nil, err
	}

	return &DBHandler{
		DB:      ci.Database,
		Session: session,
	}, nil
}

func (dbh *DBHandler) Collection(name string) Collection {
	return dbh.Session.DB(dbh.DB).C(name)
}

func (dbh *DBHandler) Find(collection string, query interface{}) Query {
	return dbh.Collection(collection).Find(query)
}

// SetGeoSpatialIndex sets 2d geospatial index for a given collection,
// assumes GeoJSON format with location as the key
func (dbh *DBHandler) SetGeoSpatialIndex(collectionName string) error {
	session := dbh.Session.Copy()
	defer session.Close()

	collection := dbh.Collection(collectionName)
	index := mgo.Index{
		Key: []string{"$2dsphere:location"},
	}

	err := collection.EnsureIndex(index)

	return err
}
