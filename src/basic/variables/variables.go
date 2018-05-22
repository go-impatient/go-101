package main

import "fmt"

// 函数外部声明变量， 只能使用 `var` 关键字，不能使用简短声明的方式。
// Go语言中没有全局变量一说，只有包内变量和函数内变量。
// 此处定义的就是包内变量。包内变量的作用域是该当前包中所有文件中都可以访问。
//var aa = 3
//var ss = "Hello World!"
//var bb = true
//var 姓名 string
//var 性别  string

// 使用 `var()`集中定义变量，定义到一个组里面，可以少写很多的 `var` 关键字
var (
	aa = 3
	ss = "Hello World!"
	bb = true
	姓名 string
	性别 string
)

// Go支持用汉字等Unicode字符命名,但从编程习惯上来说，这并不是好选择
func variableUnicode()  {
	姓名 = "李明明"
	sex := true
	if sex {
		性别 = "男人"
	} else {
		性别 = "女人"
	}

	fmt.Printf("姓名: %s\n性别:%s\n", 姓名, 性别)
}

// 声明变量
func variableDeclaration()  {
	var a int  // 声明一个类型为int, 名称为 a 的变量，如果一个变量没有被赋予任何值的话，在Go语言自动为这个变量初始化一个默认值，例如 int 的 默认值是 0
	var b int
	var s string // string 的默认值是 空字符串。
	fmt.Printf("%d %d %q\n", a, b, s) // %q 会打印出引号
}

// 先定义变量再赋值
func variableAssignment()  {
	var a int

	fmt.Println("打印默认值： ", a)
	a = 4 // 赋予一个值
	fmt.Println("打印值：", a)
	a = 8 // 再次赋予一个值
	fmt.Println("打印一个新值：", a)
}

// 声明带初始值的变量
// 多个变量可以在一条语句中声明。
// 多变量声明的语法为：var name1, name2 type = initialvalue1, initialvalue2
func variableInitialValue()  {
	//var a, b int = 3, 4
	//var s string  = "abc"
	//fmt.Println(a,b,s)

	// 如果有多个变量同时声明，我们可以采用 var 加括号批量声明的方式:
	//var (
	//	a, b int = 3, 4  // 同时声明 a, b 的整数, 并且附上初始值
	//	s string  = "abc"
	//)

	// 如果指定了初始值，则 type 可以省略
	var (
		a, b = 3, 4
		s  = "abc"
	)
	fmt.Println(a,b,s)
}

// 直接赋值，Go语言编译器自动推断变量的类型
func variableTypeDeduction() {
	var x = "Hello World!"
	fmt.Println(x)

	// 各种类型的变量混在一起使用
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

// 简短声明方式， 关键字 `var` 就可以省掉
// 变量在声明的时候如果有初始值，我们可以使用 := 操作符来简短声明： 声明变量的简短语法是 name := initialvalue，一个短变量声明操作符在一次操作中完成两件事儿：声明变量，并初始化。
// 简短声明只能用在函数内部；简短声明不能提供数据类型。
func variableShortHand() {
	// a, b, c, s := 3, int64(4), true // 简短声明要求 := 操作符左边的所有变量都有初始值
	a, b, c, s := 3, int64(4), true, "def"
	b = 5
	// b = "hig" // 一个变量不能被赋予与其类型不同的值
	fmt.Println(a, b, c, s)


	// 简短声明的语法要求 :=操作符的左边至少有一个变量是尚未声明的
	// 退化赋值的前提条件是：最少有一个新变量被声明，且必须是同一作用域
	aa, bb := 20, 30 // 声明变量aa和bb
	fmt.Println("bb 的 内存地址：", &bb)
	fmt.Println("aa is", aa, "bb is", bb)
	bb, cc := 40, 50 // bb已经声明，但cc尚未声明;注意bb退化为赋值操作，仅有 cc 是变量定义
	fmt.Println("bb 的 内存地址：", &bb)
	fmt.Println("bb is", bb, "cc is", cc)
	bb, cc = 80, 90 // 给已经声明的变量bb和cc赋新值
	fmt.Println("changed bb is", bb, "cc is", cc)
	// 上面的程序中，在 bb, cc := 40, 50 这一行，虽然变量 bb 在之前已经被声明了，但是 cc 却是新声明的变量，因此这是合法的；我通过打印 &bb 的内存地址，可以确认属于同一变量。
}

// 空标识符
func variableBlankIdentifier() (int, string) {
	i := 0
	_ = i // 编译器会将未使的局部变量当做错误;(可使  "_ = i" 规避)。

	return 1, "a-b-c"
}

func main() {
	variableUnicode()
	variableDeclaration()
	variableAssignment()
	variableInitialValue()
	variableTypeDeduction()
	variableShortHand()

	// 特殊只写变量 "_", 于忽略值占位。
	_, ii := variableBlankIdentifier()
	fmt.Println(ii)

	// 包内全局变量打印
	fmt.Println(aa, bb, ss)
}