package main // 定义包名，必须非注释行的第一行

import "fmt"

var ( // 定义全局变量与一般程序语言差不多
	gval int = 100
)

func main() { // main函数，是程序执行的入口
	fmt.Println("Hello, World!", gval)

	/*
		定义变量
	*/

	var name string   //  var 变量名 变量类型
	short_var := 10   // 简短声明，只能在函数内部使用，不能用于全局变量
	const pi = 3.1415 // 常量
	const (
		n1 = iota // 0 , iota 是一个可以被编译器修改的常量，每次 const 出现时，iota 的值会重置为 0
		n2        // 1
		_         // 2 , 使用 _ 跳过某些值
		n4        // 3
	)
	const ( // 通过 iota 实现枚举
		_  = iota             // 0
		KB = 1 << (10 * iota) // 1 << (10 * 1)
		MB = 1 << (10 * iota) // 1 << (10 * 2)
		GB = 1 << (10 * iota) // 1 << (10 * 3)
		TB = 1 << (10 * iota) // 1 << (10 * 4)
		PB = 1 << (10 * iota) // 1 << (10 * 5)
	)

	fmt.Println(name, short_var, pi, n1, n2, n4, KB, MB, GB, TB, PB)
	/*
		数组 array
	*/
	array1 := [3]int{}             // 定义数组
	array2 := [3]int{1, 2, 3}      // 初始化数组
	array3 := [...]int{1, 2, 3}    // 初始化数组，不指定长度
	array4 := [5]int{1: 10, 3: 30} // 指定下标初始化

	fmt.Println(array1, array2, array3, array4)

	// 遍历数组
	for i := 0; i < len(array4); i++ {
		fmt.Print(array4[i], ", ")
	}
	for j, v := range array4 {
		fmt.Print(j, ":", v, ", ")
	}
	fmt.Println()

	/*
		切片 slice
	*/
	// 切片的长度是动态的，所以声明时只需要指定切片中的元素类型

	// 1. 使用下标创建切片，是最原始也最接近汇编语言的方式，它是所有方法中最为底层的一种
	// 编译器会将 arr[0:3] 或者 slice[0:3] 等语句转换成 OpSliceMake 操作
	arraytmp := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice1 := arraytmp[0:3]
	fmt.Println(slice1)

	// 2. 通过字面量创建切片
	slice2 := []int{1, 2, 3, 4, 5} // 这不同于数组的声明方式，数组的声明需要指定长度，或者 ...
	fmt.Println(slice2)

	// 3. 使用 make 创建切片，make 是一个内建函数，它的主要作用是创建切片、哈希表和 Channel 等内建的数据结构
	// 运行期间创建，返回的是类型的引用
	// 切片逃逸（过大）时，会在堆上分配内存，否则在栈上分配内存
	slice3 := make([]int, 5, 10) // make([]T, size, cap)  size 是切片的长度，cap 是切片的容量
	fmt.Println(slice3)

	// 使用 len() 和 cap() 获取切片的长度和容量
	fmt.Println(len(slice3), cap(slice3))

	// 4. 使用 append() 函数向切片中添加元素
	slice3 = append(slice3, 1, 2, 3, 4, 5)
	fmt.Println(slice3)

	/*
		哈希表 map
	*/
	// map 是一种无序的键值对的集合，也称为关联数组、字典或者散列表
	hash0 := make(map[string]int, 10) // make(map[K]V, cap)  cap 是 map 的容量
	hash1 := map[string]int{"a": 1, "b": 2}
	fmt.Println(hash0, hash1)

	// 遍历 map
	for k, v := range hash1 {
		fmt.Println(k, v)
	}

	// 通过 key 获取 value
	fmt.Println(hash1["a"])

	/*
		字符串 string
	*/
	// 字符串是一种值类型，且值不可变，即创建某个文本后你无法再次修改这个文本的内容，而是新建了一个文本

	// 1. 使用双引号创建字符串
	str1 := "hello world"
	fmt.Println(str1)

	// 2. 使用反引号创建多行字符串
	str2 := `hello
	world`
	fmt.Println(str2)

	// 类型转换
	// 当我们使用 Go 语言解析和序列化 JSON 等数据格式时，经常需要将数据在 string 和 []byte 之间来回转换
	// 当需要处理中文、日文或者其他复合字符时，则需要用到rune类型。rune类型实际是一个int32。
	// Go 使用了特殊的 rune 类型来处理 Unicode，让基于 Unicode的文本处理更为方便，也可以使用 byte 型进行默认字符串处理，性能和扩展性都有照顾
	// 1. string 转 []byte
	str3 := "hello world"
	// bytes := []byte(str3) // 和下面的等价
	bytes := []rune(str3)
	fmt.Println(bytes)

	// 2. []byte 转 string
	bytes[0] = 'H'
	str4 := string(bytes)
	fmt.Println(str4)

}
