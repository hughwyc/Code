package main

import (
	"fmt"
)

//定义一个数组，并给出8个整数，求该数组中大于平均值数的个数，和小于平均数的个数。

func main() {
	arr := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	sum := 0
	bigNum := 0
	smallNum := 0

	for _, v := range arr {
		sum += v
	}

	for _, v := range arr {
		if v > sum/len(arr) {
			bigNum++
		} else if v < sum/len(arr) {
			smallNum++
		}
	}

	fmt.Printf("bigNum = %v, smallNum = %v", bigNum, smallNum)
}
