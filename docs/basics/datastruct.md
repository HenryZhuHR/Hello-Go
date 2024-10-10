---
outline: [3,6]
date: 2024-09-03
---
## 数据结构

### 数组 array


```go
array1 := [3]int{}             // 定义数组
array2 := [3]int{1, 2, 3}      // 初始化数组
array3 := [...]int{1, 2, 3}    // 初始化数组，不指定长度
array4 := [5]int{1: 10, 3: 30} // 指定下标初始化

// 遍历数组
for i := 0; i < len(array4); i++ {
    fmt.Print(array4[i], ", ")
}
for j, v := range array4 {
    fmt.Print(j, ":", v, ", ")
}
```

### 切片 slice


### 映射表 map


