package main

import "fmt"

type World struct {
	Message string
}

func main() {
	w := &World{Message: "Hello, World!"}
	fmt.Println(w.Message)
}
