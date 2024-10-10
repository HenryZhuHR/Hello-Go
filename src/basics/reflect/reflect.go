package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	author := "draven" // 每一个变量包含 pair 对，一个是类型，一个是值
	fmt.Println(" TypeOf author:", reflect.TypeOf(author))
	fmt.Println("ValueOf author:", reflect.ValueOf(author))

	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)

	if err != nil {
		fmt.Println("OpenFile err:", err)
		return
	}

	var r io.Reader = tty
	fmt.Println("TypeOf r:", reflect.TypeOf(r))

}
