package go_function

import (
	"fmt"
	"testing"
)

// 函数的参数 值传递
func FuncParamsPassByValue(num int, str string) (int, string) {
	num += 10
	str += " world"
	return num, str
}

func TestFuncParamsPassByValue(t *testing.T) {
	num, str := FuncParamsPassByValue(10, "hello")
	fmt.Printf("Return = (%d, %s) \n", num, str)
}

// 函数的参数 引用传递
func FuncParamsPassByReference(x int, y *int) {
	x += 1
	*y += 10
}
func TestFuncParamsPassByReference(t *testing.T) {
	x, y := 1, 2
	FuncParamsPassByReference(x, &y)
	fmt.Printf("x = %d, y = %d \n", x, y)
}

type MyStruct struct {
	i int
}

func FuncParamsStruct(
	ptr1 MyStruct, // 传递 结构体 时	   会拷贝结构体中的全部内容；
	ptr2 *MyStruct, // 传递 结构体指针 时	会拷贝结构体指针；
) {
	ptr1.i++
	ptr2.i++
	fmt.Printf("[   in  func] ptr1=(%v, %p) ptr2=(%v, %p)\n", ptr1, &ptr1, ptr2, &ptr2)
}

func TestFuncParamsStruct(t *testing.T) {
	ptr1 := MyStruct{i: 30}
	ptr2 := &MyStruct{i: 40}
	fmt.Printf("[before func] ptr1=(%v, %p) ptr2=(%v, %p)\n", ptr1, &ptr1, ptr2, &ptr2)
	FuncParamsStruct(ptr1, ptr2)
	fmt.Printf("[after  func] ptr1=(%v, %p) ptr2=(%v, %p)\n", ptr1, &ptr1, ptr2, &ptr2)
}

func FuncVariableParams(args ...int) {
	for _, arg := range args {
		fmt.Printf("%d, ", arg)
	}
	fmt.Println()
}

func TestFuncVariableParams(t *testing.T) {
	fmt.Printf("Variable Params: ")
	FuncVariableParams(1, 2, 3, 4, 5) // 传递任意个数的 int 值
}
