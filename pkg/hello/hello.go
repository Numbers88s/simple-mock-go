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
	say, err := h.World.Say()
	if err != nil {
		return ""
	}

	return say
}
