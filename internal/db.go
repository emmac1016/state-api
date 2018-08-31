package internal

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
)

type DBConnectionInfo struct {
	Username string
	Password string
	Host     string
	Database string
}

// ConnectDB creates a mongodb session
func ConnectDB(connInfo *DBConnectionInfo) (*mgo.Session, error) {
	log.Print("dialing with info...")
	conn := createConnectionString(connInfo)
	session, err := mgo.Dial(conn)
	// db := session.DB(connInfo.Database)
	// err = db.Login(connInfo.Username, connInfo.Password)
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

	log.Print(connString)

	return connString
}
