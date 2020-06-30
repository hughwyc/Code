package main

import (
	"fmt"
)

func main() {
	// 使用反引号 `` 来复制一大串字符串（多行）
	str1 := `package main`
	fmt.Println(str1)

	var a bool
	fmt.Printf("a = %v\n", a) // %v 表示按照变量的值输出

	fmt.Println(str1[0:2])
}
