package jsonrpc

type structuredError struct {
	c int
	m string
	d interface{}
}

func NewStructuredError(code int, message string, data interface{}) Error {
	return structuredError{code, message, data}
}

func (e structuredError) code() int {
	return e.c
}

func (e structuredError) message() string {
	return e.m
}

func (e structuredError) data() interface{} {
	return e.d
}
