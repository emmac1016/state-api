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
