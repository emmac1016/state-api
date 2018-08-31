package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/emmac1016/state-api/handlers"
	"github.com/emmac1016/state-api/internal"
	"github.com/urfave/cli"
	mgo "gopkg.in/mgo.v2"
)

// JSONData is used to extract data from state.json file
type JSONData struct {
	Name   string      `json:"state"`
	Border [][]float32 `json:"border"`
}

var conn mgo.Session

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

		connInfo := &internal.DBConnectionInfo{
			Host:     c.String("host"),
			Database: c.String("db"),
			Username: c.String("user"),
			Password: c.String("pass"),
		}
		_, err := internal.ConnectDB(connInfo)

		if err != nil {
			fmt.Println("Error getting db connection: ", err)
			return nil
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	// err := cli.NewApp().Run(os.Args)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// host := flag.String("host", "localhost", "mongo host(s) to connect to")
	// db := flag.String("db", "test", "database to connect to")
	// user := flag.String("user", "admin", "mongo username to connect with")
	// pw := flag.String("pw", "pass", "password for given mongo username")
	// flag.Parse()

	// if *host == "" || *db == "" || *user == "" || *pw == "" {
	// 	flag.PrintDefaults()
	// 	os.Exit(1)
	// }

	//TODO: add flag to pass in this data

	// err := loadFixtureData()
	// if err != nil {
	// 	log.Fatal("Could not load fixture data: ", err)
	// }
}

func LoadFixtureData() error {
	// err := dropFixtures()
	// if err != nil {
	// 	log.Print("Cannot drop fixtures: ", err)
	// 	return err
	// }

	// err = loadStateFixtures()
	// if err != nil {
	// 	log.Print("Cannot load state fixtures: ", err)
	// 	return err
	// }

	log.Print("LoadFixtureData")
	return nil
}

// func dropFixtures() error {
// 	log.Print("first dropping fixtures")
// 	conn, err := internal.ConnectDB()
// 	if err != nil {
// 		log.Print("Error getting db connection: ", err)
// 		return err
// 	}

// 	log.Print("connected to DB")
// 	session := conn.Copy()
// 	defer session.Close()

// 	db := session.DB(os.Getenv("MONGO_DB"))
// 	err = db.DropDatabase()
// 	if err != nil {
// 		log.Print("Error droping database: ", err)
// 		return err
// 	}

// 	return nil
// }

// func loadStateFixtures() error {
// 	log.Print("loading state fixtures")
// 	states, err := getStatesFromFile("../build/states.json")
// 	if err != nil {
// 		log.Print("Error in parsing file: ", err)
// 		return err
// 	}

// 	err = loadStates(states)
// 	if err != nil {
// 		log.Print("Error in inserting into db: ", err)
// 		return err
// 	}

// 	return nil
// }

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

// func loadStates(states []interface{}) error {

// 	// TODO: add flag to pass in this data
// 	connInfo := &mgo.DialInfo{
// 		Addrs:    []string{os.Getenv("MONGO_HOST")},
// 		Timeout:  10 * time.Second,
// 		Database: os.Getenv("MONGO_DB"),
// 		Username: os.Getenv("MONGO_USER"),
// 		Password: os.Getenv("MONGO_PW"),
// 	}
// 	conn, err := internal.ConnectDB()
// 	if err != nil {
// 		log.Print("Error getting db connection: ", err)
// 		return err
// 	}

// 	session := conn.Copy()
// 	defer session.Close()

// 	log.Print(os.Getenv("MONGO_DB"))
// 	collection := session.DB(os.Getenv("MONGO_DB")).C("states")
// 	//err = collection.Insert(states...)
// 	bulkInsert := collection.Bulk()
// 	bulkInsert.Insert(states...)
// 	bulkInsert.Unordered()
// 	result, err := bulkInsert.Run()
// 	log.Print("result of bulk insert:")
// 	log.Print(result)
// 	return err
// }
