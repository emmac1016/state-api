package main

import (
	"fmt"
	"log"
	"os"

	"github.com/emmac1016/state-api/internal"
	"github.com/urfave/cli"
)

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

		connInfo := internal.ConnectionInfo{
			Host:     c.String("host"),
			Database: c.String("db"),
			Username: c.String("user"),
			Password: c.String("pass"),
		}
		dbh, err := internal.NewDBHandler(&connInfo)
		if err != nil {
			fmt.Println("Error getting db connection: ", err)
			return nil
		}

		fl := internal.NewFixtureLoader(dbh)
		err = fl.LoadData()

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
