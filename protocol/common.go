package protocol

//go:generate python doc2struct.py

import "reflect"

type CommonArg interface {
	APIName() string
	Method() string
	URI() string
	Document() string
	JPName() string
}

type CommonResponse struct {
	RequestId     string
	ErrorResponse struct {
		RequestId    string
		ErrorType    string
		ErrorMessage string
	}
}

var APIlist []CommonArg

var TypeMap = map[string]reflect.Type{}
