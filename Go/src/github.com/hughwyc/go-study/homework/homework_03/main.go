package main

import (
	"fmt"
)

// 编写一个函数，输出100以内的所有素数（只能被1和本身整除的数），
// 每行显示5个，并求和

func main() {

	var slice []int // 存储所有筛选出来的素数

	for i := 1; i < 101; i ++ {
		count := 0
		if i == 1 {
			count ++
		}
		for j := 1; j < i; j ++ {
			if i % j == 0 {
				count ++
			}
		}
		if count == 1 {
			//fmt.Println(i)
			slice = append(slice, i)
		}
	}

	fmt.Println(slice)

	sum := 0

	for i := 0; i < len(slice); i ++ {

		fmt.Printf("%4d",slice[i]) // 输出一次占4位

		if (i + 1) % 5 == 0 {
			fmt.Println() // 每5个一换行
		}

		sum += slice[i]
	}

	fmt.Println()
	fmt.Println("the sum is", sum)
}
