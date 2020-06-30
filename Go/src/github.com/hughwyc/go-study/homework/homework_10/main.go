package main

import (
	"fmt"
)

// 定义一个4行4列的二维数组，逐个从键盘输入值，
// 然后将第1行和第4行的数据进行交换，将第2行和第3行的数据进行交换

func main() {

	var arr [4][4]int

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

		var temp [len(arr[0])]int
		if i == 0 || i == 1 {
			temp = arr[i]
			arr[i] = arr[len(arr)-1-i]
			arr[len(arr)-1-i] = temp
		}

		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%v\t", arr[i][j])
		}

		fmt.Println()
	}

}
