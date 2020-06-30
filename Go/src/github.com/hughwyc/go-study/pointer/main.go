package main

import (
	"fmt"
)

func main() {
	// 指针

	// 基本数据类型在内存中的布局
	var i int64 = 8
	fmt.Println("i = : ", i)
	fmt.Println("the adress of i: ", &i)

	// 下面的var ptr *int = &i
	// 1. ptr 是一个指针变量
	// 2. ptr 的类型是 *int
	// 3. ptr 本身的值是 &i
	var ptr *int
	ptr = &i
	fmt.Printf("ptr = &i = %v \n", ptr)

	// 使用*ptr可以获取ptr指向的值。
	fmt.Printf("*ptr = %v \n", *ptr)
}