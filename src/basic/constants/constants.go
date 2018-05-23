package main

import (
	"fmt"
	"math"
	"unsafe"
)

const x, y int = 1, 2
const s = "hello world!" // 类型推断

// 声明常量
// Go语言的常量命名不推荐全部大写，Go语言中的命名首字母大小写是有含义的，（即，大写为公有的，小写为私有的）。
func constantDeclaration() {
	const xx, yy  = 3, 4 // 多个常量初始化，这里未规定类型的常量，const数值可作为各种类型使用
	const s = "hi" // 在不同作用域中定义同名常量

	// 定义一组常量
	const (
		a, b = 100, 200 //类型推断
		t bool = false
	)

	fmt.Println(x,y,s,a,t)
	var c = int(math.Sqrt(float64(x*x + y*y))) // 如果xx,yy 声明为int类型,所以要强制类型转换
	var cc = int(math.Sqrt(xx*xx + yy*yy)) // 如果x,y 未规定类型，const数值可作为各种类型使用， 在这里它们又可以是int, float64,这里就是float64，所以不用类型转换

	fmt.Println(c, cc)
}


func constantFunction() {
	// 常量值还可以是 len、cap、unsafe.Sizeof 等编译期可确定结果的函数返回值。
	const (
		a = "abc"
		b = len(a)
		// 每种类型都有它的大小和对齐值，可以通过 unsafe.Sizeof 获取其大小
		c = unsafe.Sizeof(b)
	)

	println(a, b, c)
}

func main() {
	// 包内全局常量打印
	fmt.Println(x,y,s)
	constantDeclaration()
	constantFunction()
}