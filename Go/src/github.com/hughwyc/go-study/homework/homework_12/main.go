package main

import (
	"fmt"
)

// 试写出实现查找的核心代码，比如已知数组arr[10]string，里面保存了10个元素，
// 现在要查找“AA”在其中是否存在，打印提示，如果有多个“AA”，也要输出对应的下标

func main() {
	arr := [10]string{"aa", "AA", "bb", "cc", "dd", "AA", "ff", "gg", "hh", "ii"}

	// 使用顺序查找
	for i, v := range arr {
		if v == "AA" {
			fmt.Printf("\"AA\" was found in arr, the index is arr[%v]\n", i)
		}
	}

}
