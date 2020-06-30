package main

import (
	"fmt"
)

func test() bool {
	fmt.Println("test...")
	return true
}

func main() {

	// 短路与 && 。 在if语句中，当&&前的语句为false时，就不再判断&&后面的语句是否为true，直接跳过
	var i int32 = 10
	if i < 9 && test() {
		fmt.Println("ok...")
	} else {
		fmt.Println("no...")
	}
	// 短路或 || 。 同理	


	// 位运算，需要考虑补码，补码计算完成之后，还要转回原码！！！
	fmt.Println("位运算")
	fmt.Println(-1 & -2)

	// 移位运算符， 不需要考虑补码
	fmt.Println(-1 << 2) // -4

	var a int32 = -1
	fmt.Println(a << 2)
}