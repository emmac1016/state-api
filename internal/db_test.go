package internal

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	mgo "gopkg.in/mgo.v2"
)

func TestNewDBHandler(t *testing.T) {
	tests := []struct {
		name         string
		conn         ConnectionInfo
		dialFunc     func(string) (*mgo.Session, error)
		expectedConn string
		expectedDB   string
		err          error
	}{
		{
			name: "NewDBHandler connects to the database returns valid DBHandler",
			conn: ConnectionInfo{
				Username: "dbuser",
				Password: "dbpass",
				Host:     "localhost:27017",
				Database: "mydb",
			},
			dialFunc: func(url string) (*mgo.Session, error) {
				return &mgo.Session{}, nil
			},
			expectedConn: "mongodb://dbuser:dbpass@localhost:27017/mydb",
			expectedDB:   "mydb",
			err:          nil,
		},
		{
			name: "NewDBHandler returns error if failure to connect to the database",
			conn: ConnectionInfo{
				Username: "dbuser",
				Password: "dbpass",
				Host:     "localhost:27017",
				Database: "mydb",
			},
			dialFunc: func(url string) (*mgo.Session, error) {
				return nil, errors.New("Fail")
			},
			expectedConn: "mongodb://dbuser:dbpass@localhost:27017/mydb",
			expectedDB:   "mydb",
			err:          errors.New("Fail"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// set oldDialFunc to old dialFunc
			oldDialFunc := dialFunc
			// as we are exiting, revert sqlOpen back to oldSqlOpen at end of function
			defer func() { dialFunc = oldDialFunc }()
			dialFunc = tt.dialFunc

			dbh, err := NewDBHandler(&tt.conn)

			if tt.err != nil {
				assert.NotNil(t, err)
				assert.Equal(t, err, tt.err)
			} else {
				assert.Equal(t, dbh.DB, tt.expectedDB)
				assert.Equal(t, dbh.conn, tt.expectedConn)
				assert.Nil(t, err)
			}
		})
	}
}
