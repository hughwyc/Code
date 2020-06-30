package main

import "fmt"

// 已知有个排序好（升序）的数组，要求插入一个元素，最后打印该数组，顺序依然是升序

func main() {

	arr := [10]int{0, 10, 20, 30, 40, 60, 70, 80, 90, 100}
	slice := arr[:]

	num := 50

	for i := range slice {

		//fmt.Printf("i=%d, v=%d \n", i, v)

		if num < slice[1] {
			slice1 := []int{num}
			slice2 := append(slice1, slice...)
			fmt.Println("插入一个数值后的数组为：", slice2)
			break

		} else if i < len(slice) && num >= slice[i] && num < slice[i+1] {
			slice0 := []int{num} // 注意slice为引用类型，不可在原切片上进行append，可以新建一个切片，在其上进行append
			slice1 := append(slice0, slice[i+1:]...)
			slice2 := append(slice[:i+1], slice1...)
			fmt.Println("插入一个数值后的数组为：", slice2)
			break

		} else if num > slice[len(slice)-1] {
			slice1 := []int{num}
			slice2 := append(slice, slice1...)
			fmt.Println("插入一个数值后的数组为：", slice2)
			break
		}
	}

}
