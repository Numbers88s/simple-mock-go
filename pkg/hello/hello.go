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

type FrobinateArgs struct {
	After string
	Limit *int
}

func Frobinate(args FrobinateArgs) {
	limit := 100
	if args.Limit != nil {
		limit = *args.Limit
	}

	offset := decodeAfter(args.After) // just fake it

	world.Printer(world.ToolArgs{
		Limit:  limit,
		Offset: offset,
	})
}

func decodeAfter(after string) int {
	return 0
}
