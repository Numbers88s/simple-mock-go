package world

import (
	"errors"

	"github.com/hashicorp/go-multierror"
)

type World struct {
	Message string
}

type ClassicError error

var ErrMessageEmpty = errors.New("message is empty")

func New() *World {
	return &World{Message: "Hello World!"}
}

func (w *World) Say() (string, multierror.Error) {
	if w.Message == "" {
		return "", multierror.Error{}
	}

	return w.Message, multierror.Error{}
}
