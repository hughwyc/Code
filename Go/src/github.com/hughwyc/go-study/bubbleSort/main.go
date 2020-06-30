package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 使用冒泡排序法将随机的5个数进行排序

func bubbleSort(arr []int) []int {

	temp := 0

	for i := 0; i < len(arr)-1; i++ {

		for j := 0; j < len(arr)-1-i; j++ {
			if (arr)[j] > (arr)[j+1] {
				temp = (arr)[j]
				(arr)[j] = (arr)[j+1]
				(arr)[j+1] = temp
			}
		}

	}

	return arr

}

func main() {

	//产生5个随机数
	rand.Seed(time.Now().UnixNano())

	var arr []int
	var num int

	fmt.Println("请输入随机数的个数：")
	fmt.Scanln(&num)

	for i := 0; i < num; i++ {

		arr = append(arr, rand.Intn(100))
	}

	fmt.Println("原始切片：", arr)

	fmt.Println("排序后：", bubbleSort(arr))
}
