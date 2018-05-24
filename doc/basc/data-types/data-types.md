## Go 语言类型系统

### 基本数据类型

我们可以将基本类型分为三大类：

- 布尔类型
- 数字类型
- 字符串类型


- 内建字符串类型: `string`，the set of string value (eg: "hi")
- 内建布尔类型: `bool`， the set of boolean (true, false)
- 内建数字类型:
```
int8, uint8 (byte), int16, uint16, int32 (rune), uint32, int64, uint64, int, uint, uinptr.
float32, float64.
complex64, complex128.

uint8(byte)     the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8      the set of all signed  8-bit integers (-128 to 127)
int16      the set of all signed 16-bit integers (-32768 to 32767)
int32(rune)      the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64      the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32      the set of all IEEE-754 32-bit floating-point numbers
float64      the set of all IEEE-754 64-bit floating-point numbers

complex64      the set of all complex numbers with float32 real and imaginary parts
complex128      the set of all complex numbers with float64 real and imaginary parts

byte      alias for uint8
rune      alias for int32
uint      either 32 or 64 bits
int      same size as uint
uintptr      an unsigned integer large enough to store the uninterpreted bits of a pointer value

```

#### 布尔类型

一般我们用于判断条件, 它的取值范围为 true, false, 声明如下：
```
var a bool
var a = true
a := true

const a = true
```

#### 数字类型：

数字类型主要分为有符号数和无符号数，有符号数可以用来表示负数，除此之外它们还有位数的区别，不同的位数代表它们实际存储占用空间，以及取值的范围。

例如：
```
var (
  a uint8 = 1
  b int8 = 1
)

var (
    c int = 64
)
```
注意：

    每种数字类型都有取值范围，超过了取值范围，出现 overflows 的错误。
    int，uint 的长度由操作系统的位数决定，在 32 位系统里面，它们的长度未 32 bit, 64 位系统，长度为 64 bit。

#### 字符串
```
var a = "hello" //单行字符串
var c = "\"" // 转义符

var b = `hello` //原样输出

//多行输出
var d = `
line3
line1
line2
`

var str = "hello, 世界"

b := str[0]  //b is a uint8 type, like byte
fmt.Println(b)  // 104
fmt.Println(string(b)) // h

fmt.Println(len(str)) //12, 查看字符串有多少个字节
fmt.Println(len([]rune(str))) // 8 查看有多少个字符
```
#### 特殊类型

byte，uint8 别名，用于表示二进制数据的 bytes
rune，int32 别名, 用于表示一个符号

```
var str = "hello, 世界"

for _, char := range str {
  fmt.Printf("%T", char)
}
```


### 复合数据类型

- pointer types - C pointer alike.
- struct types - C struct alike.
- function types - functions are first-class types in Go.
- container types:
- array types - fixed-length container types.
- slice type - dynamic-length and dynamic-capacity container types.
- map types - maps are associative arrays (or dictionaries). The standard Go compiler implements maps as hashtables.
- channel types - channels are used to synchronize data among goroutines (the green threads in Go).
- interface types - interfaces play a key role in reflection and polymorphism.