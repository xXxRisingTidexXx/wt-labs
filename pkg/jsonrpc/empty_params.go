package jsonrpc

type emptyParams struct{}

func NewEmpty() Params {
	return emptyParams{}
}

func (p emptyParams) len() int {
	return 0
}
