
// 获取指定目录下的子文件列表

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fileloc := `D:\Documents\Downloads\BaiduNetdiskDownload\We are bears\Season1\`
    dirList, e := ioutil.ReadDir(fileloc)
    if e != nil {
        fmt.Println("read dir error")
        return
	}
	
    for i, v := range dirList {
		fmt.Println(i, "=", v.Name())
		// 在这里更改需要截取的字符，[0:2]代表截取前两个
		err1 := os.Rename(fileloc+ v.Name(), fileloc+v.Name()[0:2]+".mp4") 
		if err1 != nil {
			panic(err1)
		} else {
			fmt.Println("第",i+1,`个文件重命名成功`)
		}
    }
}