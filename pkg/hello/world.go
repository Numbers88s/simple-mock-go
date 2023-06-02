package hello

type World struct {
	Message string
}

func New() *World {
	return &World{Message: "Hello World!"}
}

func (w *World) Say() string {
	return w.Message
}
