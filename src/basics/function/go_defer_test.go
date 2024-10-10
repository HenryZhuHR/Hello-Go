package go_function

import (
	"fmt"
	"testing"
)

func TestDeferOrder(t *testing.T) {
	defer fmt.Println("Defer[1]")
	defer fmt.Println("Defer[2]")
	defer fmt.Println("Defer[3]")
}

func TestDeferReturn(t *testing.T) {
	num := 0
	fmt.Println("[1]num =", num)
	num += 1
	defer fmt.Println("[2]num =", num)
	num += 1
	fmt.Println("[3]num =", num)
}
