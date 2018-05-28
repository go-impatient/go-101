package main

import "fmt"

// 值类型的初始化
// 在Go语言中，如果你声明了一个变量但是没有对它进行赋值操作，那么这个变量就会有一个类型的默认零值。
// 这是每种类型对应的零值：
// ```
// bool    -> false                              
// numbers -> 0                                 
// string  -> ""
// pointers -> nil
// slices -> nil
// maps -> nil
// channels -> nil
// functions -> nil
// interfaces -> nil
// ```
func valueTypes() {
	var a string
	fmt.Printf("a: %#v\n", a) // 输出为 a: ""

	var b *string
	fmt.Printf("b: %#v\n", b) // 输出为 b: (*string)(nil)
	// fmt.Printf("b: %#v\n", *b) // panic: runtime error: invalid memory address or nil pointer dereference

	c := new(string)
	fmt.Printf("c: %#v\n", c) // 输出为 c: (*string)(0xc42007c1c0)
	fmt.Printf("c: %#v\n", *c) // 输出为 c: ""

	d := "hello world!"
	fmt.Printf("d: %#v\n", d) // 输出为 d: "hello world!"

	e := string("hello world!")
	fmt.Printf("e: %#v\n", e) // 输出为 e: "hello world!"

	f := &e
	fmt.Printf("f: %#v\n", f) // 输出为 f: (*string)(0xc42007c1e0)

	// 报错，cannot take the address of string("hello world!")
	// g := &string("hello world!")
}

// 引用类型的初始化
// 不管使用nil切片还是空切片，对其调用内置函数`append`、`len`和`cap`的效果都是一样的
func referenceTypes() {
	var s []string // 声明为nil字符串切片
	fmt.Printf("s: %#v\n", s) // 输出为 s: []string(nil)

	s = append(s, "hello world!")
	fmt.Printf("s: %#v\n", s) // 输出为 s: []string{"hello world!"}

	s1 := []string{"a", "b", "c"}
	fmt.Printf("s1: %#v\n", s1) // 输出为 s1: []string{"a", "b", "c"}

	s2 := []string{} // 声明为空字符串切片
	fmt.Printf("s2: %#v\n", s2) // 输出为 []string{}

	s3 := &s2
	fmt.Printf("s3: %#v\n", s3) // 输出为 s3: &[]string{}

	s4 := &s1
	fmt.Printf("s4: %#v\n", s4) // 输出为 s4: &[]string{"a", "b", "c"}

	// 不推荐的方式
	s5 := new([]string) // 声明为nil字符串切片
	fmt.Printf("s5: %#v\n", s5) // 输出为 s5: &[]string(nil)

	s6 := make([]string, 0) // 声明为空字符串切片
	fmt.Printf("s6: %#v\n", s6) // 输出为 s6: []string{}

	// 不推荐的方式
	s7 := new([]string)
	*s7 = make([]string, 0)
	fmt.Printf("s7: %#v\n", s7) // 输出为 s7: &[]string{}

	slice := [5]string{"a", "b", "c"}
	s8 := slice[:]
	fmt.Printf("s8: %#v\n", s8) // 输出为 s8: []string{"a", "b", "c", "", ""}
}

func main() {
	valueTypes()
	referenceTypes()
}