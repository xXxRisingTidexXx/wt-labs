package jsonrpc

type stringID string

func NewString(id string) ID {
	return stringID(id)
}

func NewBytes(id []byte) ID {
	return stringID(id)
}

func (id stringID) toValue() (interface{}, bool) {
	return id, true
}

func (id stringID) toNumber() (int64, bool) {
	return 0, false
}

func (id stringID) toString() (string, bool) {
	return string(id), true
}

func (id stringID) isNull() bool {
	return false
}
