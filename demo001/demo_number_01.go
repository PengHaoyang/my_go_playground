package main

import (
	"fmt"
	"reflect"
	"strconv"
)

/* 测试 int 运算的输出 */
func trySomeInt() {
	var a int = 2 // can be infer, so remove
	var b = 3
	fmt.Println("a: ", a, ", ", reflect.TypeOf(a))
	fmt.Println("b: ", b, ", ", reflect.TypeOf(b))
	fmt.Println("a+b: ", a+b) /* 不同类型的数值不能相加（比如int 和 int8） */
	fmt.Println("a-b: ", a-b)
	fmt.Println("a*b: ", a*b)
	fmt.Println("a/b: ", a/b)
	fmt.Println("~~~~~~~")
}

/* 测试 float 运算的输出 */
func trySomeFloat() {
	var a = 2.7
	var b = 3.1
	fmt.Println("a: ", a, ", ", reflect.TypeOf(a))
	fmt.Println("b: ", b, ", ", reflect.TypeOf(b))
	fmt.Println("a+b: ", a+b)
	fmt.Println("a-b: ", a-b)
	fmt.Println("a*b: ", a*b)
	fmt.Println("a/b: ", a/b)
	fmt.Println("~~~~~~~")
}

func main() {
	trySomeInt()
	trySomeFloat()
	fmt.Println(strconv.Atoi("3234"))
}
