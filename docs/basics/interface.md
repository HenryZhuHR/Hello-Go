---
outline: [3,6]
date: 2024-09-03
---

## 接口

### 接口概念

官方文档 [Interfaces](https://gobyexample.com/interfaces)

> Interfaces are named collections of method signatures.

Go关于接口的发展历史有一个分水岭，在Go1.17及以前，官方在参考手册中对于接口的定义为：一组方法的集合。

> An interface type specifies a method set called its interface.

接口实现的定义为

> A variable of interface type can store a value of any type with a method set that is any superset of the interface. Such a type is said to implement the interface

翻译过来就是，当一个类型的方法集是一个接口的方法集的超集时，且该类型的值可以由该接口类型的变量存储，那么称该类型实现了该接口。

不过在Go1.18时，关于接口的定义发生了变化，接口定义为：一组类型的集合。

An interface type defines a type set.
接口实现的定义为

A variable of interface type can store a value of any type that is in the type set of the interface. Such a type is said to implement the interface
翻译过来就是，当一个类型位于一个接口的类型集内，且该类型的值可以由该接口类型的变量存储，那么称该类型实现了该接口。并且还给出了如下的额外定义。

当如下情况时，可以称类型T实现了接口I
T不是一个接口，并且是接口I类型集中的一个元素
T是一个接口，并且T的类型集是接口I类型集的一个子集
如果T实现了一个接口，那么T的值也实现了该接口。
Go在1.18最大的变化就是加入了泛型，新接口定义就是为了泛型而服务的，不过一点也不影响之前接口的使用，同时接口也分为了两类，

### 接口基本使用

#### 接口的定义与实现
接口定义了一组方法的集合，接口类型是一种抽象类型，不会暴露出实现细节，只会暴露出方法的签名。

> 接口不能成员变量，只能定义方法。

```go
type User interface {
    GetName() string
    GetAge() int
}
```

Go 语言中接口的实现都是隐式的，我们只需要实现 `GetName() string` 和 `GetAge() int` 方法就实现了 `User` 接口，不需要显式的声明实现了某个接口。struct 实现接口的方法如下：

```go
type vipUser struct {
	name string
	age  int
}
func (v vipUser) GetName() string { return v.name}
func (v vipUser)  GetAge() int    { return v.age }
```

而在使用接口的时候，我们只需要将「实现了接口的对象」(`vipUser`)赋值给「接口类型的变量」(`User`)即可，如下：
```go
func TestInterface(t *testing.T) {
	var u User = vipUser{"Thompson", 18}
	fmt.Println("vipUser:", u.GetName(), u.GetAge())
}
// vipUser: Thompson 18
```
