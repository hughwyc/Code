package main

import "fmt"

// 编写代码统计出字符串"Hello沙河小王子"中汉字的数量

func main() {

	str := "Hello沙河小王子"
	sli := []rune(str)
	count := 0

	for _, v := range sli {
		if v > 256 {
			count++
		}
		fmt.Printf("v = %c\n", v)
		fmt.Printf("v = %v\n", v)
	}

	fmt.Println("汉字的数量是：", count)
	fmt.Println("len(str)= ", len(str)) // 20 , len()函数求的是字节的数量

}
