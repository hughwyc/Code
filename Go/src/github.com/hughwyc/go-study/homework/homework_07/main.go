package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 随机生成10个整数（范围1-100）保存到数组，并
// 倒序打印，求平均值，最大值的下标，并查找里面是否有55

func main() {

	rand.Seed(time.Now().UnixNano())
	var arr [10]int

	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(100)
	}

	fmt.Println(arr)

	sum := 0
	maxVal := arr[0]
	maxIndex := 0
	flag := false

	for i := 0; i < len(arr); i++ {
		fmt.Printf("%v\t", arr[len(arr)-1-i])
		sum += arr[i]
		if maxVal < arr[i] {
			maxVal = arr[i]
			maxIndex = i
		}
		if arr[i] == 55 {
			flag = true
		}
	}

	fmt.Println()
	fmt.Printf("average=%.2f\n", float64(sum)/float64(len(arr)))
	fmt.Println("maxIndex=", maxIndex)

	if flag == true {
		fmt.Println("55 is in the array.")
	} else {
		fmt.Println("55 is not in the array.")
	}

}
