package jsonrpc

type positionalParams []interface{}

func NewPositional(params ...interface{}) Params {
	return positionalParams(params)
}

func (p positionalParams) len() int {
	return len(p)
}
