package printer

import "fmt"

type ToolArgs struct {
	Limit      int
	Offset     int
	MethodName string
	// In the future add a field here that's required, but we don't supply it in ToolArgs in the outer layer
}

func Printer(args ToolArgs) {
	fmt.Println(args.Limit, args.MethodName)
}
