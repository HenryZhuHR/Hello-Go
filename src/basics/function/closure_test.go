package go_function

import (
	"fmt"
	"testing"
)

func TestClosure(t *testing.T) {
	fplus := func(x, y int) int {
		return x + y
	}
	fmt.Println(fplus(3, 4))
}

const LIM = 100

var fibs [LIM]uint64

func Fibonacci(n int) (res uint64) {
	// 内存缓存与动态规划 https://learnku.com/go/t/38628
	// memoization: check if fibonacci(n) is already known in array:
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = Fibonacci(n-1) + Fibonacci(n-2)
	}
	fibs[n] = res
	return
}

func TestFibonacci(t *testing.T) {
	fmt.Println("Fibonacci Result = ", Fibonacci(10))
}
