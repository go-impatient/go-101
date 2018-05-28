## Go 语言类型系统

### 基本数据类型

我们可以将基本类型分为三大类：布尔类型、数字类型、字符串类型 。


- 内建字符串类型: `string`，the set of string value (eg: "hi")
- 内建布尔类型: `bool`， the set of boolean (true, false)
- 内建数字类型:

int8, uint8 (byte), int16, uint16, int32 (rune), uint32, int64, uint64, int, uint, uinptr.
float32, float64.
complex64, complex128.

Go语言中共17种内建基本类型都属于不同的类型，内建类型不用导入任何包，这些类型的所有名称都是非导出标识符。

17种内建基本类型中有15种是数字类型。数字类型包括整数类型、浮点类型和复数类型。

Go还支持两种内置别名类型：
 - `byte`是一个内置的别名`uint8`。所以和`uint8`是相同的类型。
 - `rune`是一个内置的别名`int32`。所以和`int32`是相同的类型。

名称以u开头的整数类型是无符号类型。无符号类型的值始终为非负。类型名称中的数字表示该类型的值将在内存中占用多少二进制位。
内存中int和uint值的大小取决于操作系统，可以是32位，也可以是64位。uintptr值的大小必须大到足以存储指针值无解释位的无符号整数。

```
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
| 类型 | 长度 | 默认值 | 说明 |
|---- | -----| ------ |---------------------------------------------------------|
|bool | 1    | fasle  |       |
|byte | 1    | 0      | uint8  |
|int,uint |  4, 8 |  0 |  默认整数类型，依据目标平台，32 或者 64位   |
|int8, uint8   |  1    |   0     |  -128 ~ 127， 0 ~ 255   |
|int16, uint16   |  2    |   0     |  -32,768 ~ 32,767 ， 0 ~ 65,535 |
|int32,uint32     |   4   |   0     |  -21亿 ~ 21亿，0 ~ 42 亿  |
|int64,uint64 | 8 | 0 | |
|float32     |   4   |   0.0     |     |
|float64     |  8    |   0.0     |  默认浮点数类型   |
|complex64     |   8   |        |     |
|complex128     |  16    |        |     |
|rune  |   4   |     0   |   Unicode Code Point，int32  |  
|uintptr     |  4,8    |   0     | 足以存储指针的uint    |
|string |  | "" | 字符串，默认值为空字符串，而非NULL |

#### 布尔类型

布尔型的值只可以是常量 true 或者 false。一般我们用于判断条件, 它的取值范围为 true, false。

#### 数字类型：

数字类型主要分为有符号数和无符号数，有符号数可以用来表示负数，除此之外它们还有位数的区别，不同的位数代表它们实际存储占用空间，以及取值的范围。


注意：

    每种数字类型都有取值范围，超过了取值范围，出现 overflows 的错误。
    int，uint 的长度由操作系统的位数决定，在 32 位系统里面，它们的长度未 32 bit, 64 位系统，长度为 64 bit。

##### 复数
complex64：实部和虚部都是 float32 类型的的复数。
complex128：实部和虚部都是 float64 类型的的复数。

内建函数 `complex` 用于创建一个包含实部和虚部的复数。`complex` 函数的定义如下：
```
func complex(r, i FloatType) ComplexType
```
该函数的参数分别是实部和虚部，并返回一个复数类型。
实部和虚部应该是相同类型，也就是 float32 或 float64。
如果实部和虚部都是 float32 类型，则函数会返回一个 complex64 类型的复数。
如果实部和虚部都是 float64 类型，则函数会返回一个 complex128 类型的复数。


#### 字符串

字符串就是一串固定长度的字符连接起来的字符序列。Go的字符串是由单个字节连接起来的。Go语言的字符串的字节使用UTF-8编码标识Unicode文本。

#### 特殊类型

byte，uint8 别名，用于表示二进制数据的 bytes
rune，int32 别名, 用于表示一个符号


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

| 类型 | 长度 | 默认值 | 说明 |
|---- | -----| ------ |---------------------------------------------------------|
| pointer |     | nil  | 指针 |
| struct  |     |      | 结构体 |
| function |    | nil     | 函数  |
| array   |     |      | 数组 ，作为容器使用  |
| slice   |     | nil  | 切片，引用类型 ， 作为容器使用 |
| map   |       | nil   | 字典，引用类型 ，作为容器使用 | 
| interface  |    | nil | 接口 ，作为容器使用 |
| channel  |    | nil  | 通道，引用类型，作为容器使用 |

数组和结构体是聚合类型；它们的值由许多元素或成员字段的值组成。
数组是由同构的元素组成——每个数组元素都是完全相同的类型——结构体则是由异构的元素组成的。
数组和结构体都是有固定内存大小的数据结构。
相比之下，slice和map则是动态的数据结构，它们将根据需要动态增长。

引用类型(reference type),特指 `slice`,`map`,`channel`这三种预定义类型。
相比数字、数组等类型，引用类型拥有更为复杂的储存结构。除分配内存外，它们必须初始化一系列属性，诸如指针、长度，甚至包括哈希分布、数据队列等。
内置函数`new` 按指定类型长度分配零值内存，返回指针，并不关心类型内部构造和初始化方式。
而引用类型则必须使用 `make` 函数创建，编译器会将make转换为目标类型专用的创建函数（或指令），以确保完成全部内存分配和相关属性初始化。
