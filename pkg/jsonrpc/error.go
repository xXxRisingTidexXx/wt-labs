package jsonrpc

type Error interface {
	code() int
	message() string
	data() interface{}
}
