package main

import (
	"fmt"

	"github.com/mikhailche/grpc-test/fractalizer"
)

func main() {
	fmt.Println("Hello, playground")
	var point fractalizer.Point
	point.X = 2
	fmt.Println(point)
}
