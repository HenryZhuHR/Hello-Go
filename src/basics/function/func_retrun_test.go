package go_function

import (
	"fmt"
	"testing"
)

// 函数的返回值
func FuncReturn(input int) int {
	return input + 1
}

// 多返回值
func FuncMultiReturn(input int) (int, int) {
	return input + 1, input + 2
}

// 命名（多）返回值
func FuncNamedReturn() (y int, z int) {
	y, z = 5, 6
	return // 不需要显式返回
	// return y, z	// 需要显式返回
}

func TestReturn(t *testing.T) {
	x := 1
	fmt.Println("FuncReturn:", FuncReturn(x))
	y, z := FuncMultiReturn(x)
	fmt.Println("FuncMultiReturn:", y, z)
	y, z = FuncNamedReturn()
	fmt.Println("FuncNamedReturn:", y, z)
}
