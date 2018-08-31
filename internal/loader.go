package internal

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/emmac1016/state-api/handlers"
)

// JSONData is used to extract data from states.json file
type JSONData struct {
	Name   string      `json:"state"`
	Border [][]float32 `json:"border"`
}

// FixtureLoader is a wrapper around DB struct to expand functionality
type FixtureLoader struct {
	DB *DB
}

// LoadData drops and loads fixture data
func (fl *FixtureLoader) LoadData() error {
	err := fl.dropData()
	if err != nil {
		log.Print("Cannot drop fixtures: ", err)
		return err
	}

	err = fl.loadStateData()
	if err != nil {
		log.Print("Cannot load state fixtures: ", err)
		return err
	}
	return nil
}

func (fl *FixtureLoader) dropData() error {
	log.Print("Dropping ", fl.DB.Name)

	session := fl.DB.Connection.Copy()
	defer session.Close()

	db := session.DB(fl.DB.Name)
	err := db.DropDatabase()
	if err != nil {
		log.Print("Error droping database: ", err)
		return err
	}

	return nil
}

func (fl *FixtureLoader) loadStateData() error {
	log.Print("Loading state fixtures")

	absPath, _ := filepath.Abs("../state-api/build/states.json")
	states, err := getStatesFromFile(absPath)
	if err != nil {
		log.Print("Error in parsing file: ", err)
		return err
	}

	err = fl.loadStates(states)
	if err != nil {
		log.Print("Error in inserting into db: ", err)
		return err
	}

	return nil
}

func getStatesFromFile(fileName string) ([]interface{}, error) {
	log.Print("Parsing file: ", fileName)

	file, err := os.Open(fileName)
	if err != nil {
		log.Print("Could not open file: ", err)
		return nil, err
	}
	defer file.Close()

	var state handlers.State
	data := &JSONData{}

	scanner := bufio.NewScanner(file)
	states := make([]interface{}, 50)
	index := 0
	for scanner.Scan() {
		err := json.Unmarshal([]byte(scanner.Text()), data)
		if err != nil {
			log.Print("Could not unmarshal JSON: ", err)
			return nil, err
		}

		state = newState(data)
		states[index] = state
		index++
	}

	if err := scanner.Err(); err != nil {
		log.Print("Scanner error: ", err)
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

func (fl *FixtureLoader) loadStates(states []interface{}) error {
	log.Print("Loading state data into db")

	session := fl.DB.Connection.Copy()
	defer session.Close()

	collection := session.DB(fl.DB.Name).C("states")
	bulkInsert := collection.Bulk()
	bulkInsert.Insert(states...)
	bulkInsert.Unordered()
	_, err := bulkInsert.Run()

	return err
}
