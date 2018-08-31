package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/emmac1016/state-api/handlers"
	"github.com/emmac1016/state-api/internal"
	mgo "gopkg.in/mgo.v2"
)

// JSONData is used to extract data from state.json file
type JSONData struct {
	Name   string      `json:"state"`
	Border [][]float32 `json:"border"`
}

func init() {
	err := loadFixtureData()
	if err != nil {
		log.Fatal("Could not load fixture data: ", err)
	}
}

func loadFixtureData() error {
	err := dropFixtures()
	if err != nil {
		log.Print("Cannot drop fixtures: ", err)
		return err
	}

	err = loadStateFixtures()
	if err != nil {
		log.Print("Cannot load state fixtures: ", err)
		return err
	}

	return nil
}

func dropFixtures() error {
	log.Print("first dropping fixtures")
	conn, err := internal.ConnectDB()
	if err != nil {
		log.Print("Error getting db connection: ", err)
		return err
	}

	log.Print("connected to DB")
	session := conn.Copy()
	defer session.Close()

	db := session.DB(os.Getenv("MONGO_DB"))
	err = db.DropDatabase()
	if err != nil {
		log.Print("Error droping database: ", err)
		return err
	}

	return nil
}

func loadStateFixtures() error {
	log.Print("loading state fixtures")
	states, err := getStatesFromFile("./build/states.json")
	if err != nil {
		log.Print("Error in parsing file: ", err)
		return err
	}

	err = loadStates(states)
	if err != nil {
		log.Print("Error in inserting into db: ", err)
		return err
	}

	return nil
}

func getStatesFromFile(fileName string) ([]interface{}, error) {
	// Open file for reading
	file, err := os.Open(fileName)
	if err != nil {
		log.Print("Could not open file: ", err)
		return nil, err
	}
	defer file.Close()

	var state handlers.State
	data := &JSONData{}

	// Reading line by line
	scanner := bufio.NewScanner(file)
	states := make([]interface{}, 50)
	index := 0
	for scanner.Scan() {
		err := json.Unmarshal([]byte(scanner.Text()), data)
		if err != nil {
			log.Fatal("Could not unmarshal JSON: ", err)
			return nil, err
		}

		state = newState(data)
		states[index] = state
		index++
	}

	if err := scanner.Err(); err != nil {
		log.Print(err)
		return nil, err
	}

	return states, nil
}

func newState(data *JSONData) handlers.State {
	return handlers.State{
		Name: data.Name,
		Location: handlers.GeoJSON{
			Type:        "Polygon",
			Coordinates: data.Border,
		},
	}
}

func loadStates(states []interface{}) error {

	// TODO: add flag to pass in this data
	connInfo := &mgo.DialInfo{
		Addrs:    []string{os.Getenv("MONGO_HOST")},
		Timeout:  10 * time.Second,
		Database: os.Getenv("MONGO_DB"),
		Username: os.Getenv("MONGO_USER"),
		Password: os.Getenv("MONGO_PW"),
	}
	conn, err := internal.ConnectDB()
	if err != nil {
		log.Print("Error getting db connection: ", err)
		return err
	}

	session := conn.Copy()
	defer session.Close()

	log.Print(os.Getenv("MONGO_DB"))
	collection := session.DB(os.Getenv("MONGO_DB")).C("states")
	//err = collection.Insert(states...)
	bulkInsert := collection.Bulk()
	bulkInsert.Insert(states...)
	bulkInsert.Unordered()
	result, err := bulkInsert.Run()
	log.Print("result of bulk insert:")
	log.Print(result)
	return err
}
