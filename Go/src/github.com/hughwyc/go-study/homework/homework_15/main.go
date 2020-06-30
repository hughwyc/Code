package main

import "fmt"

// 编写一个函数，可以接收一个数组，该数组有5个数，请找出最大的数和最小的数，并显示对应的数组下标

func myFunc(arr [5]int) (max, min, maxIndex, minIndex int) {

	max = arr[0]
	min = arr[0]
	maxIndex = 0
	minIndex = 0

	for i, v := range arr {
		if v > max {
			max = v
			maxIndex = i
		}
		if v < min {
			min = v
			minIndex = i
		}
	}
	return max, min, maxIndex, minIndex

}

func main() {

	arr := [5]int{1, 6, 3, 4, 5}

	max, min, maxIndex, minIndex := myFunc(arr)

	fmt.Printf("max = %v, min = %v, maxIndex = %v, minIndex = %v\n", max, min, maxIndex, minIndex)

}
