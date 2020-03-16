package web

import "fmt"

//HTTPError error for handling resonse errors
type HTTPError interface {
	String() string
}

//ClientError used for display errors to the client
//should not have any verbose details/backtrace
type ClientError struct {
	Reason string `json:"reason"`
}

func (e ClientError) String() string {
	return e.Reason
}

//NewClientError creates a new client error
//for returning error responses client side
func NewClientError(format string, args ...interface{}) HTTPError {
	return ClientError{Reason: fmt.Sprintf(format, args...)}
}
