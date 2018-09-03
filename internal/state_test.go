package internal

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Args struct {
	longitude float64
	latitude  float64
	states    []State
}

type ExpectedValues struct {
	state   string
	results []State
}

func TestNewStateRepo(t *testing.T) {
	tests := []struct {
		name string
		exp  ExpectedValues
	}{
		{
			name: "NewStateRepo successfully obtains db handler and return valid StateRepo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr, err := NewStateRepo()
			assert.NotNil(t, sr.dbh)
			assert.Nil(t, err)
		})
	}
}

func TestFindStateByCoordinates(t *testing.T) {
	tests := []struct {
		name string
		args Args
		exp  ExpectedValues
	}{
		{
			name: "FindStateByCoordinates successfully finds and returns state",
			args: Args{
				longitude: -77.036133,
				latitude:  40.513799,
			},
			exp: ExpectedValues{
				state:   "Pennsylvania",
				results: []State{State{Name: "Pennsylvania"}},
			},
		},
		{
			name: "FindStateByCoordinates returns ",
			args: Args{
				longitude: -102.65625,
				latitude:  52.696361078274485,
			},
			exp: ExpectedValues{
				state:   "",
				results: make([]State, 0),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr, _ := NewStateRepo()

			results, err := sr.FindStateByCoordinates(tt.args.longitude, tt.args.latitude)

			assert.Equal(t, results, tt.exp.results)
			assert.Nil(t, err)
		})
	}
}

func TestMarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		args Args
		exp  string
	}{
		{
			name: "MarshalJSON should only return state name",
			args: Args{
				states: []State{State{Name: "Pennsylvania", Location: GeoJSON{Type: "Polygon"}}},
			},
			exp: "[\"Pennsylvania\"]\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var b bytes.Buffer
			enc := json.NewEncoder(&b)
			enc.SetEscapeHTML(false)
			if err := enc.Encode(tt.args.states); err != nil {
				t.FailNow()
			}

			assert.Equal(t, tt.exp, b.String())

		})
	}
}
