---
outline: [3,6]
date: 2024-09-03
---

## 函数

### 函数定义

```go
func FuncName([parameter list]) [return type] {
    // function body
}
```

```go
func func(num int, str string) { }
func func(num1, num2 int) { }
```

Go 使用的是值传递，不管参数是基本类型，结构体还是指针，都会对传递的参数进行拷贝，区别无非是拷贝的目标对象还是拷贝指针。拷贝指针，也就是会同时出现两个指针指向原有的内存空间。

### 函数参数传递
函数的参数 值传递和引用传递
```go
func FuncParamsPassByReference(x int, y *int) {
	x += 1
	*y += 10
}
func TestFuncParamsPassByReference(t *testing.T) {
	x, y := 1, 2
	FuncParamsPassByReference(x, &y)
	fmt.Printf("x = %d, y = %d \n", x, y)
}
```

<!-- 传值还是传指针？

表面上看，指针参数性能会更好，但是要注意被复制的指针会延长目标对象的生命周期，还可能导致它被分配到堆上，其性能消耗要加上堆内存分配和垃圾回收的成本。

在栈上复制小对象，要比堆上分配内存要快的多。如果复制成本高，或者需要修改原对象，使用指针更好。 -->

### 函数返回值

- **函数的返回值**

```go
func FuncReturn() int {
	return 1
}
```

- **多返回值**

```go
func FuncMultiReturn() (int, int) {
	return 1, 2
}
```

- **命名（多）返回值**

```go
func FuncNamedReturn() (y int, z int) {
	y, z = 5, 6
	return // 不需要显式返回
	// return y, z	// 需要显式返回
}
```

### 结构体和结构体指针作为函数参数

```go
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
// [before func] ptr1=({30}, 0x140000960f8) ptr2=(&{40}, 0x140000a00c0)
// [   in  func] ptr1=({31}, 0x14000096110) ptr2=(&{41}, 0x140000a00c8)
// [after  func] ptr1=({30}, 0x140000960f8) ptr2=(&{41}, 0x140000a00c0)
```

### 可变参数函数参数

```go
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
```
### 闭包

```go
package function

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
```

```go
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
```


### 延迟调用 defer

defer关键字可以使得一个函数延迟一段时间调用，在函数返回之前这些defer描述的函数最后都会被逐个执行，看下面一个例子

#### defer 的执行顺序
```go
func TestDeferOrder(t *testing.T) {
	defer fmt.Println("Defer[1]")
	defer fmt.Println("Defer[2]")
}
// Defer[2]
// Defer[1]
```

可以看出
- defer 的调用类似于栈，先进后出的顺序执行。 因为 `Defer[n]` 是倒序打印的


#### defer 的参数传递


- 输出 `[3]num = 2` ，说明 defer 语句中的变量是在执行 defer 语句的时候就已经确定了


#### defer 的应用场景


##### 资源处理的成对操作

`defer` 通常用于 `open`/`close`, `connect`/`disconnect`, `lock`/`unlock` 等这些成对的操作, 来保证在任何情况下资源都被正确释放。在这个角度来说, `defer` 操作和 Java 中的 `try ... finally` 语句块是类似的

```go
var mutex sync.Mutex
var count = 0

func increment() {
    mutex.Lock()
    defer mutex.Unlock()
    count++
}
```
在 `increment()` 函数中, 为了避免竞态条件的出现, 而使用了 `Mutex` 进行加锁。而在进行并发编程时, 加锁了却忘记(或某种情况下 `unlock` 没有被执行), 往往会造成灾难性的后果。为了在任意情况下, 都要保证在加锁操作后, 都进行对应的解锁操作, 我们可以使用 `defer` 调用解锁操作。

```go
cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
if err != nil { log.Fatal(err) }
defer cli.Close()
```

### 参考资料

- defer-[1] [Go defer 深度剖析篇（1）— defer 全套使用姿势 - qiya的文章 - 知乎](https://zhuanlan.zhihu.com/p/351176808)

- defer-[2] [详解defer实现机制(附上三道面试题,我不信你们都能做对) - Golang梦工厂的文章 - 知乎](https://zhuanlan.zhihu.com/p/343223542)
