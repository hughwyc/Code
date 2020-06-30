package main

import (
	"fmt"
	_ "os"
)

func addElement(flag bool, sliceIn []string) ([]string) {

	sliceIn = append(sliceIn, "a")
	if flag == true {
		sliceIn = addElement(false, sliceIn)
	}

	return sliceIn
}

func main() {

	// 从外部新建一个数组，经过一个函数之后数组内元素增加，再返回这个新的数组(切片)。
	var slice []string

	sliceOut := addElement(true, slice)

	fmt.Println(sliceOut)



	// Rename 可以实现文件的移动 ！
	// err := os.Rename(`D:\Test_Coding\file02.txt`, `D:\file02.txt`)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// var times [5][0]int
    // for range times {
    //     fmt.Println("hello")
	// }
	
	// /* 创建切片 */
	// var strings []string
	// // {"abc", "123", "xyz","789"}  
	// // printSlice(strings)

	// /* 向切片添加一个元素 */
	// strings = append(strings, "1a1")
 
	// /* 打印原始切片 */
	// fmt.Println("strings ==", strings)
 
	

}