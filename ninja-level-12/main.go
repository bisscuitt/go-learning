package main

import (
	"fmt"
	"github.com/bisscuitt/go-learning/ninja-level-12/dog"
)

func main() {
	hy := 5
	dy := dog.Years(hy)
	fmt.Printf("%d human years is %d in dog years\n", hy, dy)
}
