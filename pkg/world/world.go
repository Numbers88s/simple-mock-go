package world

import (
	"errors"

	"github.com/hashicorp/go-multierror"
)

type World struct {
	Message string
}

var ErrMessageEmpty = errors.New("message is empty")

func New() *World {
	return &World{Message: "Hello World!"}
}

func (w *World) Say() (string, error) {
	var result error

	if w.Message == "" {
		result = multierror.Append(result, ErrMessageEmpty)
		return "", result
	}

	return w.Message, nil
}
