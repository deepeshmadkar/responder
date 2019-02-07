package responder

import (
	"encoding/json"
	"net/http"
)

const (
	successMessage = "Success."
	successCode    = 200
	errorMessage   = "Something went wrong."
	errorCode      = 400
)

// This is the format which will always used to respond
type jsonMessage struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

// Success method is using custom success to respond.
// This just a decoration method.
func Success(w http.ResponseWriter, data interface{}) {
	CustomSuccess(w, "", 0, data)
}

// Error method is using custom error to respond.
// This just a decoration method.
func Error(w http.ResponseWriter, data interface{}) {
	CustomError(w, "", 0, data)
}

// CustomSuccess method is using repond method to respond.
// If you want to use custom message and custom code then you this method to respond.
func CustomSuccess(w http.ResponseWriter, message string, code int, data interface{}) {
	if code == 0 {
		code = successCode
	}

	if message == "" {
		message = successMessage
	}
	respond(w, false, message, code, data)
}

// CustomError method is using repond method to respond.
// if you want to use custom message and custom code then you this method to respond.
func CustomError(w http.ResponseWriter, message string, code int, data interface{}) {

	if code == 0 {
		code = errorCode
	}

	if message == "" {
		message = errorMessage
	}
	respond(w, true, message, code, data)
}

// this method handles the respond.
func respond(w http.ResponseWriter, err bool, message string, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	jsonResponse := jsonMessage{
		Error:   err,
		Message: message,
		Code:    code,
		Data:    data,
	}
	response, _ := json.Marshal(jsonResponse)
	w.WriteHeader(code)
	w.Write(response)
}
