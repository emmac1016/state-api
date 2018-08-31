package internal

import (
	"log"

	"gopkg.in/mgo.v2"
)

// ConnectDB creates a mongodb session
func ConnectDB(connInfo *mgo.DialInfo) (*mgo.Session, error) {
	log.Print("dialing with info...")
	session, err := mgo.DialWithInfo(connInfo)
	log.Print("session recieved")
	if err != nil {
		log.Print("Error in creating db session: ", err)
		return nil, err
	}

	return session, nil
}
