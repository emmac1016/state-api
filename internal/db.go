package internal

import (
	"log"

	"gopkg.in/mgo.v2"
)

// DBHandler holds the info needed to interact with the db
type DBHandler struct {
	DB      string
	conn    string
	session *mgo.Session
}

// Database defines the actions that a DBHandler can execute
type Database interface {
	BulkInsert(string, ...interface{}) (*mgo.BulkResult, error)
	Collection(string) *mgo.Collection
	Find(string, interface{}) *mgo.Query
	SetGeoSpatialIndex(string) error
}

var dialFunc = mgo.Dial

// NewDBHandler wraps connection handling
func NewDBHandler(ci *ConnectionInfo) (*DBHandler, error) {
	conn := ci.createConnectionString()
	dbh := &DBHandler{
		DB:   ci.Database,
		conn: conn,
	}
	err := dbh.Connect()
	if err != nil {
		log.Print("Failure to create DBHandler: ", err)
		return nil, err
	}

	return dbh, nil
}

// Connect wraps mgo.Dial function
func (dbh *DBHandler) Connect() error {
	session, err := dialFunc(dbh.conn)
	if err != nil {
		log.Print("Error in creating db session: ", err)
		return err
	}

	dbh.session = session

	return nil
}

// Collection is a wrapper around mgo's way of accessing the collection
func (dbh *DBHandler) Collection(name string) *mgo.Collection {
	return dbh.session.DB(dbh.DB).C(name)
}

// Find is a wrapper around mgo's Find function
func (dbh *DBHandler) Find(collectionName string, query interface{}) *mgo.Query {
	return dbh.Collection(collectionName).Find(query)
}

// BulkInsert is a wrapper around mgo's BulkInsert struct and behavior
func (dbh *DBHandler) BulkInsert(collectionName string, docs ...interface{}) (*mgo.BulkResult, error) {
	session := dbh.session.Copy()
	defer session.Close()

	collection := dbh.Collection(collectionName)
	bulkInsert := collection.Bulk()
	bulkInsert.Insert(docs...)
	bulkInsert.Unordered()
	return bulkInsert.Run()
}

// SetGeoSpatialIndex sets 2d geospatial index for a given collection,
// assumes GeoJSON format with location as the key
func (dbh *DBHandler) SetGeoSpatialIndex(collectionName string) error {
	session := dbh.session.Copy()
	defer session.Close()

	collection := dbh.Collection(collectionName)
	index := mgo.Index{
		Key: []string{"$2dsphere:location"},
	}

	err := collection.EnsureIndex(index)

	return err
}
