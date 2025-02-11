package proto

type notConnectedError string

func (e *notConnectedError) Error() string {
	return "not connected"
}

var NotConnectedError error = new(notConnectedError)

type invalidDataError string

func (e *invalidDataError) Error() string {
	return "invalid data"
}

var InvalidDataError error = new(invalidDataError)

type protoNotFoundError string

func (e *protoNotFoundError) Error() string {
	return "proto not found"
}

var ProtoNotFoundError error = new(protoNotFoundError)
