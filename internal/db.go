package internal

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
)

// DBConnectionInfo holds all data necessary for a db connection
type DBConnectionInfo struct {
	Username string
	Password string
	Host     string
	Database string
}

// DB holds database name and connection used to run queries
type DB struct {
	Name       string
	Connection *mgo.Session
}

// NewDB returns DB struct with connection and database name
func NewDB(connInfo *DBConnectionInfo) (*DB, error) {
	connString := createConnectionString(connInfo)
	conn, err := connect(connString)
	if err != nil {
		log.Print("Error in creating DB instance: ", err)
		return nil, err
	}

	return &DB{
		Name:       connInfo.Database,
		Connection: conn,
	}, nil
}

func connect(conn string) (*mgo.Session, error) {
	session, err := mgo.Dial(conn)
	if err != nil {
		log.Print("Error in creating db session: ", err)
		return nil, err
	}

	return session, nil
}

func createConnectionString(connInfo *DBConnectionInfo) string {
	connString := fmt.Sprintf("mongodb://%s:%s@%s/%s",
		connInfo.Username,
		connInfo.Password,
		connInfo.Host,
		connInfo.Database,
	)

	return connString
}
