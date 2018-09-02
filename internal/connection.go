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

type Connection interface {
	Dial() error
}

type ConnectionHandler struct {
	conn string
	s    Session
}

func NewConnectionHandler(ci *ConnectionInfo) *ConnectionHandler {
	conn := ci.createConnectionString()
	return &ConnectionHandler{
		conn: conn,
	}
}

func (ch *ConnectionHandler) Dial() error {
	log.Print("calling dial here")
	session, err := mgo.Dial(ch.conn)
	if err != nil {
		log.Print("Error in creating db session: ", err)
		return err
	}

	ch.s = session

	return nil
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

func (ci *ConnectionInfo) createConnectionString() string {
	connString := fmt.Sprintf("mongodb://%s:%s@%s/%s",
		ci.Username,
		ci.Password,
		ci.Host,
		ci.Database,
	)

	return connString
}
