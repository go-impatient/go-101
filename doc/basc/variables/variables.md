## 什么是变量
    
变量（Variable）就是给某个内存地址起的一个名字，拥有指定 `名称` 和 `类型` 的 数据储存位置，所以说是一段用来储存数据的内存。
我们用变量来存储某个特定类型的值。
作为静态类型语言，Go变量总是有固定的数据类型，类型决定了变量内存的长度和存储格式。我们只能修改变量值，无法改变类型。
通过类型转换或指针操作，我们可用不同方式修改变量值，但这并不意味着改变了变量类型。
因为内存分配发生在运行期，所以在编码阶段我们用一个易于阅读的名字来表示这段内存。实际上，编译后的机器码从不使用变量名，而是直接通过内存地址来访问目标数据。保存在符号表中的变量名等信息可被删除，或用于输出更详细的错误信息。
在 Go 中有多种声明变量的语法。Go语言采用的是后置类型的声明方式:
` 
  <命名> <类型>
`
   
即，变量声明的时候，把名称写在前，类型写在后。这样不是为了凸显与众不同，而是为了让声明更加清晰易懂。


## 变量声明
在Go语言中我们通常使用关键字 `var` 声明变量。  
变量不能在运行期间改变变量类型。使用`var`定义变量，自动初始化值，可以省略变量类型，由编译器自动推断类型。 
全局变量不可以使用 `:=` 的声明方式，只能使用 `var` 声明变量方式。
    

## 变量的生命周期

变量的生命周期指的是在程序运行期间变量有效存在的时间间隔。对于在包一级声明的变量来说，它们的生命周期和整个程序的运行周期是一致的。而相比之下，在局部变量的声明周期则是动态的：从每次创建一个新变量的声明语句开始，直到该变量不再被引用为止，然后变量的存储空间可能被回收。函数的参数变量和返回值变量都是局部变量。它们在函数每次被调用的时候创建。

## 变量初始化

Golang的数据类型可以分为：值类型 与 引用类型。

### 值类型的初始化
Go中值类型（以 string 为例）的初始化方式：
```
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
```

通过`new()`创建的变量, 是一个指针。而new(string) 为string分配一片内存，初始化为 0。

在Go语言中，如果你声明了一个变量但是没有对它进行赋值操作，那么这个变量就会有一个类型的默认零值。
这是每种类型对应的零值：
```
bool    -> false                              
numbers -> 0                                 
string  -> ""

pointers -> nil
slices -> nil
maps -> nil
channels -> nil
functions -> nil
interfaces -> nil
```
`new(T)`是一个分配内存的内建函数，它的第一个参数是一个类型，不是一个值, 但不同于其他语言中同名的new所作的工作，它只是将内存清零，而不是初始化内存。new(T)为一个类型为T的新项目分配了值为零的存储空间并返回其地址，也就是一个类型为*T的值。用Go的术语来说，就是它返回了一个指向新分配的类型为T的零值的指针。
在Go语言的代码定义如下：
```
func new(t Type) *Type 
```

### 引用类型的初始化
Go中的引用类型仅有三种：`map`、 `slice`、 `channel`, 这里举例就用slice来进行。
```
	var s []string // 声明为nil字符串切片
	fmt.Printf("s: %#v\n", s) // 输出为 s: []string(nil)

	s = append(s, "hello world!")
	fmt.Printf("s: %#v\n",s) // 输出为 s: []string{"hello world!"}

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
	// var p *[]string = new([]string) // 为切片结构分配内存；*p == nil；很少这样使用
	s7 := new([]string)
	*s7 = make([]string, 0)
	fmt.Printf("s7: %#v\n", s7) // 输出为 s7: &[]string{}
	
	slice := [5]string{"a", "b", "c"}
	s8 := slice[:]
	fmt.Printf("s8: %#v\n", s8) // 输出为 s8: []string{"a", "b", "c", "", ""}
```

`new(T)` 为每个新的类型T分配一片内存，初始化为 0 并且返回类型为*T的内存地址：这种方法 返回一个指向类型为 T，值为 0 的地址的指针
这里所谓的值为0，并不是数值0，而是go的默认0值，对应 `slice` 就是nil。

在Go中绝对不会采用`new`这种方式来初始化 `slice` 的，原因是为什么呢？我这里先简单说一下，是因为Go中`slice`是如下定义的：
```
type slice struct{
	array unsafe.Pointer
	len   int
	cap   int
}
```
如果用 new 因为返回的是 *T 的内存地址，无法完成对 `slice`的初始化，无法让slice正常使用，想要让他可以正常使用，就得像s7的处理方式一样，再用make对应  *T 进行一次初始化。如果这么干，你说是不是有毛病才用的方式？

slice 的初始化需要初始化 len、cap的值，让 array 指向一个数组的指针。完成这些初始化后，slice才能正常使用。

总结:
- make只用于map，slice和channel，并且不返回指针。
- 要获得一个显式的指针，使用new进行分配，或者显式地使用一个变量的地址。
