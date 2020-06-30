package main

import (
	"fmt"
)

// 试保存13579五个奇数到数组，并倒序打印

func main() {

	arr := [5]int{1, 3, 5, 7, 9}
	for i := range arr {
		fmt.Printf("%v\t", arr[len(arr)-1-i])
	}

}
