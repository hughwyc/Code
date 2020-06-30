package main

import "fmt"

// 定义一个3行4列的二维数组，逐个从键盘输入值，编写程序将四周的数清零

func main() {

	var arr [3][4]int

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("please input the value of arr[%v][%v]: ", i, j)
			fmt.Scanln(&arr[i][j])
		}
	}

	fmt.Println("\ninitial array: \n")

	for i := 0; i < len(arr); i++ {

		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%v\t", arr[i][j])
		}

		fmt.Println()
	}

	fmt.Println("\nfinal array: \n")

	for i := 0; i < len(arr); i++ {

		for j := 0; j < len(arr[i]); j++ {
			if i == 0 || i == len(arr)-1 || j == 0 || j == len(arr[i])-1 {
				arr[i][j] = 0
			}
			fmt.Printf("%v\t", arr[i][j])
		}

		fmt.Println()
	}

}
