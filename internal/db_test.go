package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emmac1016/state-api/mocks"
)

func TestNewDBHandler(t *testing.T) {
	tests := []struct {
		name         string
		conn         ConnectionInfo
		want         string
		expectedConn string
		expectedDB   string
		err          error
	}{
		{
			name: "NewDBHandler returns valid DBHandler struct",
			conn: ConnectionInfo{
				Username: "dbuser",
				Password: "dbpass",
				Host:     "localhost:27017",
				Database: "mydb",
			},
			expectedConn: "mongodb://dbuser:dbpass@localhost:27017/mydb",
			expectedDB:   "mydb",
			err:          nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sess := &mocks.Session{}
			conn := &mocks.Connection{}
			conn.On("Dial", tt.expectedConn).Return(sess, tt.err).Once()
			conn.On("createConnectionString").Return(expectedConn)

			dbh, err := NewDBHandler(conn)

			if tt.err != nil {
				assert.NotNil(t, err)
			} else {
				assert.Equal(t, dbh.DB, tt.expected)
				assert.Nil(t, err)
				conn.AssertExpectations(t)
			}
		})
	}
}
