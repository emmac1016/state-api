package internal

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDefaultConnection(t *testing.T) {
	tests := []struct {
		name    string
		expConn ConnectionInfo
	}{
		{
			name: "GetDefaultConnection create ConnectionInfo with env db values",
			expConn: ConnectionInfo{
				Host:     os.Getenv("APP_MONGO_HOST"),
				Database: os.Getenv("MONGO_DB"),
				Username: os.Getenv("MONGO_USER"),
				Password: os.Getenv("MONGO_PW"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualConn := GetDefaultConnection()
			assert.Equal(t, &tt.expConn, actualConn)
		})
	}
}

func TestCreateConnectionString(t *testing.T) {
	tests := []struct {
		name      string
		expConn   ConnectionInfo
		expString string
	}{
		{
			name: "createConnectionString creates expected connection string based on ConnectionInfo",
			expConn: ConnectionInfo{
				Host:     "localhost:27017",
				Database: "mydb",
				Username: "dbuser",
				Password: "dbpass",
			},
			expString: "mongodb://dbuser:dbpass@localhost:27017/mydb",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualString := tt.expConn.createConnectionString()
			assert.Equal(t, tt.expString, actualString)
		})
	}
}
