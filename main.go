package main

import (
	"fmt"

	"github.com/numbers88s/simple-mock-go/pkg/hello"
)

func main() {
	w := &hello.World{Message: "Hello, World!"}
	fmt.Println(w.Message)
}
