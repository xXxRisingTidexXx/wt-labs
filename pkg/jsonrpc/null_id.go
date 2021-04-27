package jsonrpc

type nullID struct{}

func (id nullID) toValue() (interface{}, bool) {
	return nil, true
}

func (id nullID) toNumber() (int64, bool) {
	return 0, false
}

func (id nullID) toString() (string, bool) {
	return "", false
}

func (id nullID) isNull() bool {
	return true
}
