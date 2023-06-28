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

func (w *World) Say() (string, ClassicError) {
	var result ClassicError

	if w.Message == "" {
		result = multierror.Append(result, ErrMessageEmpty)
		return "", result
	}

	return w.Message, nil
}
