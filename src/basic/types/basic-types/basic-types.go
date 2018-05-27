package main

import (
	"fmt"
	"unsafe"
	"math/cmplx"
	"math"
	"strconv"
)

// 数字类型
// int 和 uint 根据不同的底层平台（Underlying Platform）来决定位数，在 32 位系统里面，它们的长度未 32 bit, 64 位系统，长度为 64 bit。。
func numericInt() {
	var (
		a int = 98
		b uint = 98
		// b uint = -2 // 每种数字类型都有取值范围，超过了取值范围，出现 overflows 的错误。
	)
	fmt.Println(a,b)

	// 格式说明符 %T 用于打印类型，而 %d 用于打印字节大小
	// Go 的 unsafe 包提供了一个 Sizeof 函数，该函数接收变量并返回它的字节大小。
	// unsafe 包应该小心使用，因为使用 unsafe 包可能会带来可移植性问题。出于演示的目的，我们在这里是可以使用的。
	fmt.Printf("a的类型是：%T, 大小是：%d\n",a, unsafe.Sizeof(a))
	fmt.Printf("b的类型是%T, 大小是：%d\n",b, unsafe.Sizeof(b))
	// 从上面的输出，我们可以推断出 a 和 b 为 int 类型，且大小都是 32 位（4 字节）。
	// 如果你在 64 位系统上运行上面的代码，会有不同的输出。在 64 位系统下，a 和 b 会占用 64 位（8 字节）的大小。
}

func numericInt2() {
	type Minute int
	type Hour int
	minutes := Minute(80)
	hour := Hour(8)

	if minutes > 60 {
		fmt.Println("Minutes is greater than 60")
	}
	if hour < 9 {
		fmt.Println("Hhours is greater than 9")
	}

	// invalid operation: minutes > hour (mismatched types Minute and Hour)
	//if minutes > hour {
	//	fmt.Println("This will never be executed")
	//}
}

func numericInt3() {
	type Duration int64
	var dur Duration

	// 本质上，他们并不是同一种类型，所以对于Go这种强类型语言，他们是不能相互赋值的。
	// cannot use int64(100) (type int64) as type Duration in assignment
	// dur = int64(100)

	dur = Duration(100)
	fmt.Println(dur)
}

func numericFloat() {
	a, b := 3.1415, 5.97
	fmt.Printf("a的类型是 %T, b的类型是 %T\n", a, b)
	sum := a + b
	diff := a - b
	fmt.Println("和：", sum, "差：", diff)

	no1, no2 := 56, 89
	fmt.Println("和：", no1+no2, "差：", no1-no2)


}

func numericFloat2()  {

	// 两个浮点数相加减，精度丢失
	// 出现浮点数不精确的原因是，浮点数储存至内存中时，2的-1、-2……-n次方不能精确的表示小数部分，所以再把这个数从地址中取出来进行计算就出现了偏差
	// float32和float64直接互转会精度丢失, 四舍五入后错误
	// int64转float64在数值很大的时候出现偏差
	// 两位小数乘100强转int, 比期望值少了1

	f1, f2 := 1.0, 0.9
	fmt.Println("精度丢失：", f1-f2);
	fmt.Println("恢复精度：",Round2(f1-f2, 2))
	fmt.Println("恢复精度：",Round(f1-f2, 2))

	// case: float32==>float64
	// 从数据库中取出80.45, 历史代码用float32接收
	var a float32 = 80.45
	var b float64
	// 有些函数只能接收float64, 只能强转
	b = float64(a)
	// 打印出值, 强转后出现偏差
	fmt.Println(a) //output:80.45
	fmt.Println(b) //output:80.44999694824219
	// ... 四舍五入保留小数点后1位, 期望80.5, 结果是80.4

	// case: int64==>float64
	var c int64 = 987654321098765432
	fmt.Printf("%.f\n", float64(c)) //output:987654321098765440

	// case: int(float64(xx.xx*100))
	var d float64 = 1129.6
	var e int64 = int64(d * 100)
	fmt.Println(e) //output:112959
}

// 利用fmt.Sprintf()
func Round2(f float64, n int) float64 {
	floatStr := fmt.Sprintf("%."+strconv.Itoa(n)+"f", f)
	inst, _ := strconv.ParseFloat(floatStr, 64)
	return inst
}
// 利用math.Trunc()
func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}

func numericComplex()  {
	c1 := complex(3, 4)
	c2 := 3 + 4i
	fmt.Println(cmplx.Abs(c1))
	fmt.Println(cmplx.Abs(c2))
	cadd := c1 + c2
	fmt.Println("sum:", cadd)
	cmul := c1 * c2
	fmt.Println("product:", cmul)
}

// 欧拉公式
func euler()  {
	a := cmplx.Pow(math.E, 1i * math.Pi) + 1
	b := cmplx.Exp(1i * math.Pi) + 1
	fmt.Printf("a:%.3f\nb:%.3f\n", a, b)
}

// 勾股定理
// 如果直角三角形两直角边分别为A，B，斜边为C，那么 A^2+B^2=C^2;即直角三角形两直角边长的平方和等于斜边长的平方。
func triangle()  {
	var a, b int = 3,4
	var c int
	// 类型转换是强制的
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}


// 布尔类型
func boolean()  {
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)
	c := a && b
	fmt.Println("c:", c)
	d := a || b
	fmt.Println("d:", d)

	var e bool // e 默认值是false
	e = true
	fmt.Println("e:", e)
}


func main()  {
	numericInt()
	numericInt2()
	numericInt3()
	numericFloat()
	numericFloat2()
	numericComplex()
	euler()
	boolean()
}