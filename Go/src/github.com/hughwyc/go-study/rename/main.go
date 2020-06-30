package main

import (
	"fmt"
	"os"
)

func main(){
	err1 := os.Rename(`C:\Users\Hugh\Desktop\01test.txt`, `C:\Users\Hugh\Desktop\01.txt`)
	if err1 != nil {
		panic(err1)
	} else {
		fmt.Println(`文件重命名成功`)
	}
}