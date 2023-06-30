package world

import (
	"errors"
	"fmt"

	"github.com/hashicorp/errwrap"
)

type World struct {
	Message string
	Name    string
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
	return errwrap.Wrapf("world: %s", nil).Error(), nil
}

type ToolArgs struct {
	Limit      int
	Offset     int
	MethodName string
	// In the future add a field here that's required, but we don't supply it in ToolArgs in the outer layer
}

func Printer(args ToolArgs) {
	fmt.Println(args.Limit, args.MethodName)
}
