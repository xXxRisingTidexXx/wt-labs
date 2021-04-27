package jsonrpc

type notificationID struct{}

func NewNotification() ID {
	return notificationID{}
}

func (id notificationID) toValue() (interface{}, bool) {
	return nil, false
}

func (id notificationID) toNumber() (int64, bool) {
	return 0, false
}

func (id notificationID) toString() (string, bool) {
	return "", false
}

func (id notificationID) isNull() bool {
	return false
}
