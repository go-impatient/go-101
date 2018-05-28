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

参考资料：
- https://www.jianshu.com/p/18d3bde7d835
- https://www.tapirgames.com/blog/golang-has-no-reference-values

### 值类型的初始化
Go中值类型（以 string 为例）的初始化方式：
```
var a1 string
fmt.Printf("a1: %#v \n", a1) // a1: ""

var a2 *string
fmt.Printf("a2: %#v \n", a2) // a2: (*string)(nil)
// panic: runtime error: invalid memory address or nil pointer dereference
//fmt.Printf(**"a2: %#v \n"**, *a2)

a3 := new(string)
fmt.Printf("a3: %#v \n", a3) // a3: (*string)(0xc42000e200)
fmt.Printf(**"a3: %#v \n"**, *a3) // a3: ""

a4 := "hello"
fmt.Printf("a4: %#v \n", a4) // a4: "hello"

a5 := string("hello")
fmt.Printf("a5: %#v \n", a5) // a5: "hello"

a6 := &a5
fmt.Printf("a6: %#v \n", a6) // a6: (*string)(0xc42000e1e0)

// 报错，cannot make type string
//a7 := make(string, 1)

// 报错，cannot take the address of string("hello")
//a8 := &string("hello")
```
注释部分是输出的信息，可以看到有的结果打出来是个值，有的是一个指针。这一部分重点我只想说明两个点：

Go会自动将申明变量初始化为0值，所谓的0值是：int就是0，string就是空字符，bool就是false等
对于通过new创建的变量，是一个指针，它与var声明的变量是不同的，var声明的变量仅是一个nil。而new(string) 为string分配一片内存，初始化为 0。通过上面注释的报错信息大家可以理解。

### 引用类型的初始化
Go中的引用类型仅有三种：`map`、 `slice`、 `channel`, 这里举例就用slice来进行。
```
var s1 []string
fmt.Printf("s1: %#v\n", s1) // s1: []string(nil)
s1 = append(s1, "hello")
fmt.Printf("s1: %#v\n", s1) // s1: []string{"hello"}

var s2 *[]string
fmt.Printf("s2: %#v\n", s2) // s2: (*[]string)(nil)

s3 := []string{"a", "b", "c"}
fmt.Printf("s3: %#v\n", s3) // s3: []string{"a", "b", "c"}

s4 := &[]string{}
fmt.Printf("s4: %#v\n", s4) // s4: &[]string{}

s5 := &s3
fmt.Printf("s5: %#v\n", s5) // s5: &[]string{"a", "b", "c"}

s6 := new([]string)
fmt.Printf("s6: %#v\n", s6) // s6: &[]string(nil)
// first argument to append must be slice; have *[]string
//s6 = append(s6, "hello") // 这是一个空引用的指针，所以报错

s7 := make([]string, 0)
fmt.Printf("s7: %#v\n", s7) // s7: []string{}

// 有毛病才用这种方式
s8 := new([]string)
*s8 = make([]string, 0)
fmt.Printf("s8: %#v\n", s8) // s8: &[]string{}

arr := [5]string{"a", "b", "c"}
s9 := arr[:]
fmt.Printf("s9: %#v\n", s9) // s9: []string{"a", "b", "c", "", ""}
```
这里我重点分析一下s6、s7、s8这三种初始化方式。
先说s6，使用的是new。
new(T) 为每个新的类型T分配一片内存，初始化为 0 并且返回类型为*T的内存地址：这种方法 返回一个指向类型为 T，值为 0 的地址的指针
这里所谓的值为0，并不是数值0，而是go的默认0值，对应 `slice` 就是nil。

在Go中绝对不会采用这种方式来初始化 `slice` 的，原因是为什么呢？我这里先简单说一下，是因为Go中`slice`是如下定义的：
```
type slice struct{
	array unsafe.Pointer
	len   int
	cap   int
}
```
如果用 new 因为返回的是 *T 的内存地址，无法完成对 `slice`的初始化，无法让slice正常使用，想要让他可以正常使用，就得像s8的处理方式一样，再用make对应  *T 进行一次初始化。如果这么干，你说是不是有毛病才用的方式？

slice 的初始化需要初始化 len、cap的值，让 array 指向一个数组的指针。完成这些初始化后，slice才能正常使用。

总结
关于自定义结构的初始化与上面string的初始化一样，map、channel与slice一样。大家可以自己写点代码试试

- make只用于map，slice和channel，并且不返回指针。
- 要获得一个显式的指针，使用new进行分配，或者显式地使用一个变量的地址。
