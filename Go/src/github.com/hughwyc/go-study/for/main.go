package main

import (
	"fmt"
)

func main() {
	// 打印1-100之间所有是9的倍数的整数的个数及总和
	var count byte
	var sum int

	for i := 1; i <= 100; i ++ {
		if i % 9 == 0 {
			fmt.Println("i = ", i)
			count ++
			sum = sum + i
		}
	}

	fmt.Println("count = ", count)
	fmt.Println("sum = ", sum)

	var num int
	num = 9

	for i := 0; i <= num; i ++ {
		// fmt.Println(i, " + ", num - i, " = ", num )
		fmt.Printf("%v  +  %v  =  %v \n", i, num - i, num)
	}

	// // 打印金字塔
	// var n = 8
	// for i := 1; i <= n; i ++ {

	// 	for j := 1; j <= n - 1; j ++ {
	// 		if i <= n / 2 {
	// 			if i + j == n / 2 + 1 || j - i == n / 2 - 1 {
	// 				fmt.Printf("*")
	// 			} else {
	// 				fmt.Printf(" ")
	// 			}
	// 		} else {
	// 			if i - j == n / 2 || i + j == n + n / 2 {
	// 				fmt.Printf("*")
	// 			} else {
	// 				fmt.Printf(" ")
	// 			}
	// 		}

	// 	}
		
	// 	fmt.Println()
	// }

		// 打印金字塔
		var n = 11
		for i := 1; i <= n; i ++ {
	
			for j := 1; j <= n; j ++ {
				if i <= n / 2 + 1 {
					if i + j == n / 2 + 2 || j - i == n / 2  {
						fmt.Printf("*")
					} else {
						fmt.Printf(" ")
					}
				} else {
					if i - j == n / 2 || i + j == n + n / 2 + 1 {
						fmt.Printf("*")
					} else {
						fmt.Printf(" ")
					}
				}
	
			}
			
			fmt.Println()
		}
}