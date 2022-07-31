package controller

import "strings"

// Define shape of json response
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct{}

// MakeSuccessResponse creates a response with success message
func MakeResponse(status bool, message string, data interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Data:    data,
		Errors:  nil,
	}
}

// MakeErrorResponse creates a response with error message
func MakeErrorResponse(message string, err string, data interface{}) Response {
	errors := strings.Split(err, "\n")
	return Response{
		Status:  false,
		Message: message,
		Data:    data,
		Errors:  errors,
	}
}
