package main

import (
	"fmt"
)

func main() {

	fmt.Println("----------小小计算器----------")
	fmt.Println("1.加法\n2.减法\n3.乘法\n4.除法\n0.退出\n请选择：")

	var operator int
	fmt.Scanln(&operator)

	if operator == 1 {
		fmt.Println("10 + 5 = 15")
	} else if operator == 2 {
		fmt.Println("10 - 5 = 5")
	} else if operator == 3 {
		fmt.Println("10 * 5 = 50")
	} else if operator == 4 {
		fmt.Println("10 / 5 = 2")
	} else if operator == 0 {
		fmt.Println("程序退出！")
	} else {
		fmt.Println("输入序号有误！")
	}
}