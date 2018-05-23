package main

import "fmt"

// 普遍的枚举类型
func enumsDeclaration()  {
	const (
		Sunday = 0
		Monday = 1
		Tuesday = 2
		Wednesday = 3
		Thursday = 4
		Friday = 5
		Saturday = 6
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}

// 自增值枚举类型
func enumsIota()  {
	// 自定义一个Weekday命名类型
	type Weekday int

	// 关键字 iota 定义常量组中从 0 开始按 计数的 增枚举值。
	const (
		Sunday Weekday = iota    // 0, iota 是自增值的
		Monday           // 1 通常省略后续 表达式。
		Tuesday          // 2
		Wednesday        // 3
		Thursday         // 4
		Friday           // 5
		Saturday         // 6
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	const (
		c = iota
		cpp
		_   // 下划线，表示跳过
		java
		javascript
		ruby
		python
	)
	fmt.Println(c, cpp, java, javascript, ruby, python)

	const (
		_ = iota // 0
		KB = 1 << (10 * iota) // 1 << (10 * 1)
		MB // 1 << (10 * 2)
		GB // 1 << (10 * 3
		TB // 1 << (10 * 4)
		PB // 1 << (10 * 5)
	)
	fmt.Println(KB, MB, GB, TB, PB)
}

func main() {
	enumsDeclaration()
	enumsIota()
}