package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Args struct {
	longitude float64
	latitude  float64
}

func TestGetState(t *testing.T) {
	tests := []struct {
		name        string
		args        Args
		expResponse string
		expStatus   int
	}{
		{
			name: "GetState returns Pennsylvania for longitude -77.036133 and latitude 40.513799",
			args: Args{
				longitude: -77.036133,
				latitude:  40.513799,
			},
			expResponse: "[\"Pennsylvania\"]\n",
			expStatus:   http.StatusOK,
		},
		{
			name: "GetState returns empty array for a Canadian coordinate",
			args: Args{
				longitude: -102.65625,
				latitude:  52.696361078274485,
			},
			expResponse: "[]\n",
			expStatus:   http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			host := os.Getenv("APP_HOST")
			port := os.Getenv("APP_PORT")
			url := fmt.Sprintf("http://%s%s/?longitude=%v&latitude=%v", host, port, tt.args.longitude, tt.args.latitude)
			req := httptest.NewRequest("POST", url, nil)
			w := httptest.NewRecorder()
			GetState(w, req, nil)

			resp := w.Result()
			body, _ := ioutil.ReadAll(resp.Body)

			assert.Equal(t, tt.expResponse, string(body))
			assert.Equal(t, tt.expStatus, resp.StatusCode)
		})
	}
}

func TestGetState_BadRequests(t *testing.T) {
	tests := []struct {
		name        string
		request     string
		expResponse string
		expStatus   int
	}{
		{
			name:        "GetState returns 400 for missing latitude",
			request:     "?longitude=-77.036133",
			expResponse: "{\"message\": \"'longitude' and 'latitude' fields are required\"}",
			expStatus:   http.StatusBadRequest,
		},
		{
			name:        "GetState returns 400 for missing longitude",
			request:     "?latitude=40.513799",
			expResponse: "{\"message\": \"'longitude' and 'latitude' fields are required\"}",
			expStatus:   http.StatusBadRequest,
		},
		{
			name:        "GetState returns 400 for bad longitude value",
			request:     "?longitude=NotRight&latitude=40.513799",
			expResponse: "{\"message\": \"invalid longitude value\"}",
			expStatus:   http.StatusBadRequest,
		},
		{
			name:        "GetState returns 400 for bad latitude value",
			request:     "?longitude=-77.036133&latitude=NotRight",
			expResponse: "{\"message\": \"invalid latitude value\"}",
			expStatus:   http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			host := os.Getenv("APP_HOST")
			port := os.Getenv("APP_PORT")
			url := fmt.Sprintf("http://%s%s/%s", host, port, tt.request)
			req := httptest.NewRequest("POST", url, nil)
			w := httptest.NewRecorder()
			GetState(w, req, nil)

			resp := w.Result()
			body, _ := ioutil.ReadAll(resp.Body)

			assert.Equal(t, tt.expResponse, string(body))
			assert.Equal(t, tt.expStatus, resp.StatusCode)
		})
	}
}
