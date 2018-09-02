package internal

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/jinzhu/copier"
)

// JSONData is used to extract data from states.json file
type JSONData struct {
	Name   string      `json:"state"`
	Border [][]float64 `json:"border"`
}

// FixtureLoader is a wrapper around DB struct to expand functionality
type FixtureLoader struct {
	dbh *DBHandler
}

func NewFixtureLoader(dbh *DBHandler) *FixtureLoader {
	return &FixtureLoader{dbh: dbh}
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

	log.Print("Enforcing geospatial index on states collection")
	err = fl.dbh.SetGeoSpatialIndex("states")
	if err != nil {
		log.Print("Failure to set geo spatial index: ", err)
		return err
	}

	return nil
}

func (fl *FixtureLoader) dropData() error {
	log.Print("Dropping ", fl.dbh.DB)

	session := fl.dbh.Session.Copy()
	defer session.Close()

	db := session.DB(fl.dbh.DB)
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

func (fl *FixtureLoader) loadStates(states []interface{}) error {
	log.Print("Loading state data into db")
	_, err := fl.dbh.BulkInsert("states", states...)

	return err
}

func getStatesFromFile(fileName string) ([]interface{}, error) {
	log.Print("Parsing file: ", fileName)

	file, err := os.Open(fileName)
	if err != nil {
		log.Print("Could not open file: ", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var states []interface{}

	for scanner.Scan() {
		data := JSONData{}
		err := json.Unmarshal([]byte(scanner.Text()), &data)
		if err != nil {
			log.Print("Could not unmarshal JSON: ", err)
			return nil, err
		}

		coordinates := make([][][]float64, 1)
		coordinates[0] = data.Border

		state := State{
			Name: data.Name,
			Location: GeoJSON{
				Type:        "Polygon",
				Coordinates: coordinates,
			},
		}

		newState := State{}
		copier.Copy(&newState, &state)
		states = append(states, newState)
	}

	if err := scanner.Err(); err != nil {
		log.Print("Scanner error: ", err)
		return nil, err
	}

	return states, nil
}
