package go_interface

import (
	"fmt"
	"testing"
)

type User interface { // 接口不能成员变量
	GetName() string
	GetAge() int
}

// struct 实现接口
// Go 语言中接口的实现都是隐式的，我们只需要实现 getName() string 和 getAge() int 方法就实现了 person 接口。
type vipUser struct {
	name string
	age  int
}

func (v vipUser) GetName() string { return v.name }
func (v vipUser) GetAge() int     { return v.age }

func TestInterface(t *testing.T) {
	// var person Pers
	var u User = vipUser{"Thompson", 18}
	fmt.Println("vipUser:", u.GetName(), u.GetAge())
}
