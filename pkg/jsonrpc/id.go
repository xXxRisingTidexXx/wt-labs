package jsonrpc

type ID interface {
	toValue() (interface{}, bool)
	toNumber() (int64, bool)
	toString() (string, bool)
	isNull() bool
}
