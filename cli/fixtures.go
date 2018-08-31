package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/emmac1016/state-api/handlers"
	"github.com/emmac1016/state-api/internal"
	"github.com/urfave/cli"
)

// JSONData is used to extract data from state.json file
type JSONData struct {
	Name   string      `json:"state"`
	Border [][]float32 `json:"border"`
}

type FixtureLoader struct {
	DB *internal.DB
}

func main() {
	app := cli.NewApp()
	app.Name = "Fixture Loader"
	app.Usage = "load test data for local development"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "host",
			Value:  "localhost:27017",
			Usage:  "mongo host(s) to connect to",
			EnvVar: "MONGO_HOST",
		},
		cli.StringFlag{
			Name:   "db",
			Value:  "test",
			Usage:  "database to connect to",
			EnvVar: "MONGO_DB",
		},
		cli.StringFlag{
			Name:   "user",
			Value:  "admin",
			Usage:  "mongo username to connect with",
			EnvVar: "MONGO_USER",
		},
		cli.StringFlag{
			Name:   "pass",
			Value:  "",
			Usage:  "password for given mongo username",
			EnvVar: "MONGO_PW",
		},
	}

	app.Action = func(c *cli.Context) error {
		// All arguments are required
		for _, flagName := range c.FlagNames() {
			if !c.IsSet(flagName) {
				fmt.Printf("%s was not given", flagName)
				return nil
			}
		}

		connInfo := internal.DBConnectionInfo{
			Host:     c.String("host"),
			Database: c.String("db"),
			Username: c.String("user"),
			Password: c.String("pass"),
		}
		db, err := internal.NewDB(&connInfo)
		if err != nil {
			fmt.Println("Error getting db connection: ", err)
			return nil
		}

		fl := &FixtureLoader{DB: db}
		err = fl.loadData()

		if err != nil {
			fmt.Println("Loading fixtures failed: ", err)
			return nil
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func (fl *FixtureLoader) loadData() error {
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
