// 获取指定目录下的子文件列表

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fileloc := `D:\Documents\科研数据\M1表面粗糙度\TC4-进口粉-2020-06-12打印\`
	dirList, e := ioutil.ReadDir(fileloc)
	count := 0
	if e != nil {
		fmt.Println("read dir error")
		return
	}

	for i, v := range dirList {
		fmt.Printf("%v = %v, %T\n", i, v.Name(), v.Name())
		// 在这里更改需要截取的字符，[0:2]代表截取前两个

		if strings.HasPrefix(v.Name(), "20") {
			count++
			fmt.Println(v.Name()[1:])
			err1 := os.Rename(fileloc+ v.Name(), fileloc+ v.Name()[1:])
			if err1 != nil {
				panic(err1)
			} else {
				fmt.Println("第",count,`个文件重命名成功`)
			}
		}

	}
}