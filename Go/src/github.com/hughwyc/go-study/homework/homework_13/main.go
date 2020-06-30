package main

import (
	"fmt"
	"math/rand"
	"time"
)

//随机生成10个整数（1-100之间），使用冒泡排序法进行排序，
// 然后使用二分法查找，查找是否有90这个数，并显示下标，
// 如果没有则提示“找不到该数”

func bubbleSort(arr *[10]int) {
	temp := 0

	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}
}

func binarySearch(arr [10]int, startIndex, endIndex, num int) {

	middleIndex := (endIndex + startIndex) >> 1

	if num == (arr)[startIndex] {
		fmt.Printf("the num %v was found in arr, the index is %v.\n", num, startIndex)
	} else if num > (arr)[startIndex] && num < (arr)[middleIndex] {
		binarySearch(arr, startIndex+1, middleIndex-1, num)
	} else if num == (arr)[middleIndex] {
		fmt.Printf("the num %v was found in arr, the index is %v.\n", num, middleIndex)
	} else if num > (arr)[middleIndex] && num < (arr)[endIndex] {
		binarySearch(arr, middleIndex+1, endIndex-1, num)
	} else if num == (arr)[endIndex] {
		fmt.Printf("the num %v was found in arr, the index is %v.\n", num, endIndex)
	} else if endIndex-startIndex < 0 || num < (arr)[startIndex] || num > (arr)[endIndex] {
		fmt.Printf("the num %v was not found in arr.\n", num)
	}

}

func main() {

	rand.Seed(time.Now().UnixNano())
	var arr [10]int

	for i := range arr {
		arr[i] = rand.Intn(100)
	}

	fmt.Println("initial arr = \n", arr)
	bubbleSort(&arr)
	fmt.Println("bubbleSorted arr = \n", arr)

	arr2 := [10]int{1, 2, 3, 4, 5, 6, 7, 80, 90, 100}
	fmt.Println("arr2 = \n", arr2)
	binarySearch(arr2, 0, len(arr)-1, 90)
}
