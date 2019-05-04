package models

//Alert ...
type AlertError struct {
	Type string
	Code string
	Body interface{}
}

type AlertString struct {
	Type string
	Code string
	Body string
}
