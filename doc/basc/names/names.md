## 命名

在Go语言中对函数名、变量名、常量名、类型名、自定义类型名、语句标号、包名等进行命
名都遵循一个简单的命名规则：
一个名字必须以字母（Unicode字母）或者 下划线开头，后面可以跟任意数量的字母、数字或者下划线组合而成；

命名建议：
 - 区分大小字母。
 - 推荐使用驼峰命名的拼写格式。
 - 局部变量推荐优先使用短命名。
 - 不要使用保留关键字。
 - 不建议使用与预定义常量 、类型、内置函数相同的命名。
 - 专有名称通常会全部大写，例如 escapeHTML。

尽管Go支持用汉字等Unicode字符命名，但从编程习惯上来说，这并不是好选择。

符号名字首字母大小写决定了其作用域。首字母大写的为导出成员，可被包外引用，而小写则仅能在包内使用。

 Go语言中类似`if`和`switch`的关键字有25个；关键字不能用于自定义名字，只能在特定语法结构中使用。

 ```
break      default       func     interface   select
case       defer         go       map         struct
chan       else          goto     package     switch
const      fallthrough   if       range       type
continue   for           import   return      var
 ```

 此外，还有大约30多个预定义的名字，比如`int`和`true`等，主要对对应内建的常量、类型和函数。

内建常量: 
```
true false iota nil
```

内建类型: 
```
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr
float32 float64 complex128 complex64
bool byte rune string error
```

内建函数: 
```
make len cap new append copy close delete
complex real imag
panic recover
```

这些内部预定于的名字并不是关键字，你可以再定义中重新使用它们。在一些特殊的场景中重新定义它们也是有意义的，但是也要注意避免过度而引起语义混乱。

如果一个名字是大写字母开头的（备注：必须是在函数外部定义的包级名字；包级函数名本身也是包级名字），那么它将是导出的，也就是说可以被外部的包访问，例如fmt包的Printf函数就是导出的，可以在fmt包外部访问。包本身的名字一般总是用小写字母。