package hello

import "github.com/numbers88s/simple-mock-go/pkg/world"

type Hello struct {
	World *world.World
}

func New() *Hello {
	m := world.New()
	return &Hello{World: m}
}

func (h *Hello) Say() string {
	return h.World.Say()
}
