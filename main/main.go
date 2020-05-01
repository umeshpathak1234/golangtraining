package main

import (
	"fmt"

	"github.com/umeshpathak1234/golangtraining/testing"
)

func main() {
	fmt.Println("Hi There")
	fmt.Println("The output is", testing.Newfunc(2, 3))
	fmt.Println("The new structures created are", testing.Teststruct())
	fmt.Println("NEw function", testing.Teststructone("Rock", "wa", 27))
}
