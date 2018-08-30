package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"gopkg.in/mgo.v2/bson"
)

// StateData is used to extract data from state.json file
type StateData struct {
	Name   string      `json:"state"`
	Border [][]float32 `json:"border"`
}

// State defines the mongo document structure
type State struct {
	ID          bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	Name        string        `bson:"name" json:"name"`
	Type        string        `bson:"type" json"type"`
	Coordinates [][]float32   `bson:"coordinates" json"coordinates"`
}

// GeoJSON holds the longitude & latitude data to query from
// type GeoJSON struct {
// 	Type        string    `bson:"type" json"type"`
// 	Coordinates []float32 `bson:"coordinates" json"coordinates"`
// }

func main() {
	// conn := fmt.Sprintf("mongodb://%s/", os.Getenv("MONGO_HOST"))
	// session, err := mgo.Dial(conn)
	// if err != nil {
	// 	log.Fatal("Could not connect to db: ", err)
	// 	os.Exit(1)
	// }
	// defer session.Close()

	file, err := os.Open("./build/states.json")
	if err != nil {
		log.Fatal("Could not open file: ", err)
		os.Exit(1)
	}
	defer file.Close()

	var state State
	stateData := &StateData{}
	// states := make([]interface{}, 50)
	// index := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())

		err := json.Unmarshal([]byte(scanner.Text()), stateData)
		if err != nil {
			log.Fatal(err)
		}

		state = State{
			Name:        stateData.Name,
			Type:        "Polygon",
			Coordinates: stateData.Border,
		}

		// fmt.Println(stateData.Name)
		// fmt.Println(stateData.Border)
		fmt.Println(state.Coordinates)
		fmt.Println()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
