package internal

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
)

// ConnectionInfo holds all data necessary for a db connection
type ConnectionInfo struct {
	Username string
	Password string
	Host     string
	Database string
}

// NewConnection connects and returns session
func NewConnection(ci *ConnectionInfo) (*mgo.Session, error) {
	conn := ci.createConnectionString()
	return connect(conn)
}

func connect(conn string) (*mgo.Session, error) {
	session, err := mgo.Dial(conn)
	if err != nil {
		log.Print("Error in creating db session: ", err)
		return nil, err
	}

	return session, nil
}

func (ci *ConnectionInfo) createConnectionString() string {
	connString := fmt.Sprintf("mongodb://%s:%s@%s/%s",
		ci.Username,
		ci.Password,
		ci.Host,
		ci.Database,
	)

	return connString
}
