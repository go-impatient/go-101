package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

func eval(a, b int, op string) (int, error) {
	switch op {
		case "+":
			return a + b, nil
		case "-":
			return a-b, nil
		case "*":
			return a*b, nil
		case "/":
			q, _ := div(a, b)
			return q, nil
		default:
			return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

func div(a, b int) (q, r int)  {
	q = a/b
	r = a%b
	return
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Println("Calling function %s with args " + "(%d, %d)", opName, a, b)
	return op(a,b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sumArags(values ...int) int {
	sum := 0
	for i := range values  {
		sum += values[i]
	}

	return sum
}

func swap(a, b int) {
	a, b = b, a
}

func swap2(a, b *int)  {
	*b, *a = *a, *b
}

func swap3(a, b int) (int, int) {
	return b, a
}

func chpointer(a *int)  {
	// 指针不能运算
	// *a = *a + 1
	// *a += 1
	*a++
	return
}

func main() {
	if result, err :=  eval(1, 2, "x"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	q, r := div(13,3)
	fmt.Println(q, r)

	// fmt.Println(apply(pow, 3,4))
	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3,4))

	fmt.Println(sumArags(1,2,3))

	var a, b = 3, 4

	//swap(a, b)
	//
	//fmt.Println(a, b)
	//
	//swap2(&a, &b)
	//
	//fmt.Println(a, b)

	a, b = swap3(a, b)

	fmt.Println(a, b)

	c := 10
	chpointer(&c)
	fmt.Println(c)
}
