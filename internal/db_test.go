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
				assert.Nil(t, dbh)
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

func TestConnect(t *testing.T) {
	tests := []struct {
		name     string
		conn     string
		dialFunc func(string) (*mgo.Session, error)
		err      error
	}{
		{
			name: "Connect successfully gets db session",
			conn: "mongodb://dbuser:dbpass@localhost:27017/mydb",
			dialFunc: func(url string) (*mgo.Session, error) {
				return &mgo.Session{}, nil
			},
			err: nil,
		},
		{
			name: "DBHandler won't have a session if Connect fails to get db session",
			conn: "mongodb://dbuser:dbpass@localhost:27017/mydb",
			dialFunc: func(url string) (*mgo.Session, error) {
				return nil, errors.New("Fail")
			},
			err: errors.New("Fail"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// set oldDialFunc to old dialFunc
			oldDialFunc := dialFunc
			// as we are exiting, revert sqlOpen back to oldSqlOpen at end of function
			defer func() { dialFunc = oldDialFunc }()
			dialFunc = tt.dialFunc

			dbh := &DBHandler{
				conn: tt.conn,
			}
			err := dbh.Connect()

			if tt.err != nil {
				assert.NotNil(t, err)
				assert.Equal(t, err, tt.err)
				assert.Nil(t, dbh.session)
			} else {
				assert.NotNil(t, dbh.session)
				assert.Nil(t, err)
			}
		})
	}
}
