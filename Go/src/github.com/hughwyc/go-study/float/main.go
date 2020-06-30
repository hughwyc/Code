package main

import "fmt"

func main() {

	var price float32 = 89.12
    fmt.Println("price=",price)

	// 浮点数=符号位+指数位+尾数位。其中尾数部分可能丢失。float64比float32精度高
	// 浮点数都是有符号位的
	var n1 float32 = -999999985333
	var n2 float64 = -180808255333
	fmt.Println("n1=", n1,"n2=",n2)
	// 调试过程中，一行语句前面的tab空位符只是作为1个位，用于定位错误位置。
}