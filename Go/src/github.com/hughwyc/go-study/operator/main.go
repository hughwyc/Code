package main

import (
	"fmt"
)

func main() {

	// 如果运算的数都是整数，那么除后，去掉小数部分，保留整数部分
	fmt.Println(10 / 4) // 结果为2
	fmt.Println(10 / 4.0) // 结果为2.5
	fmt.Println(10.0 / 4) // 结果为2.5

	// 取模的公式：a % b = a - (a / b) * b

	// 例题1 假如还有97天放假，问：xx个星期零xx天

	fmt.Println(97 / 7,"个星期")
	fmt.Println("零", 97 % 7,"天")

	// 例题2 华氏温度转摄氏温度
	var hua float32 = 134.2
	fmt.Println(5.0 / 9 * (hua - 100)) // 注意要是5.0


}