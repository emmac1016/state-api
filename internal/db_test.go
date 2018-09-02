package internal

import (
	"testing"
)

func TestNewDB(t *testing.T) {
	tests := []struct {
		name      string
		want      string
		expectErr bool
	}{
		{
			name:      "NewDB returns valid StateRepo struct",
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock *mgo.Session, *mgo.Database, *mgo.Collection, and
			// *mgo.Query
			// session := &mgo.Session{}
			// database := &mgo.Database{}
			// collection := &mgo.Collection{}
			// query := &mgo.Query{}

			// We expect Dial against localhost
			// mgo.EXPECT().Dial("localhost").Return(session, nil)

			// // We expect the named database to be opened
			// session.EXPECT().DB("test").Return(database)

			// // We also expect the named collection to be opened
			// database.EXPECT().C("items").Return(collection)

			// // We then expect a query to be created against the collection
			// collection.EXPECT().Find(bson.M{"_id": bson.ObjectIdHex("52f6aef226f149b7048b4567")}).Return(query)
		})
	}
}
