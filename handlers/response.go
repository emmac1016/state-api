package handlers

import "net/http"

func fail(response []byte, w http.ResponseWriter, errorCode ...int) {
	var httpErrorCode int
	if errorCode == nil {
		httpErrorCode = http.StatusInternalServerError
	} else {
		httpErrorCode = errorCode[0]
	}

	buildResponse(response, httpErrorCode, w)
}

func buildResponse(response []byte, status int, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}
