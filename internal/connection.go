package internal

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/mgo.v2"
)

// ConnectionInfo holds all data necessary for a db connection
type ConnectionInfo struct {
	Username string
	Password string
	Host     string
	Database string
}

// NewConnection connects and returns session given connection info
func NewConnection(ci *ConnectionInfo) (*mgo.Session, error) {
	conn := ci.createConnectionString()
	return connect(conn)
}

// GetDefaultConnection returns connection info for the App based on environment
func GetDefaultConnection() *ConnectionInfo {
	return &ConnectionInfo{
		Host:     os.Getenv("APP_MONGO_HOST"),
		Database: os.Getenv("MONGO_DB"),
		Username: os.Getenv("MONGO_USER"),
		Password: os.Getenv("MONGO_PW"),
	}
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
