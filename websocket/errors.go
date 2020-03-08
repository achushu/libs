package websocket

import "errors"

var (
	ErrClientNotFound = errors.New("could not find client")
)

//go:generate stringer -type=ConnectionErrorType -linecomment
type ConnectionErrorType int

const (
	ClientNotFound ConnectionErrorType = iota // could not find client
)

type ConnectionError struct {
	error
	ClientID string
}

func NewConnectionError(t ConnectionErrorType, id string) ConnectionError {
	return ConnectionError{
		errors.New(ClientNotFound.String() + " " + id),
		id,
	}
}
