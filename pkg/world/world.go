package world

import (
	"errors"

	"github.com/hashicorp/errwrap"
)

type World struct {
	Message string
}

type AppError struct {
	Err error
}

func (e *AppError) WrappedErrors() []error {
	return []error{e.Err}
}

type ClassicError error

var ErrMessageEmpty = errors.New("message is empty")

func New() *World {
	return &World{Message: "Hello World!"}
}

func (w *World) Say() (string, error) {
	err := errwrap.Wrapf("world: %s", ErrMessageEmpty)
	if w.Message == "" {
		return "", err
	}

	return w.Message, nil
}
