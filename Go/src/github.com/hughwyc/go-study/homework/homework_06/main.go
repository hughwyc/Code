package main

import (
	"fmt"
)

// 输出小写的a-z以及大写的Z-A

func main() {

	for i := 0; i < 26; i ++ {
		fmt.Printf(string(i + 97))
	}

	for i := 0; i < 26; i ++ {
		fmt.Printf(string(90 - i))
	}

}
