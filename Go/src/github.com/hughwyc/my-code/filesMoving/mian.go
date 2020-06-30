package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
)

// 本案例实现：将文件夹下满足要求的文件移动到指定文件夹下

// GetAllFile Golang循环遍历文件夹（包括子文件夹）下的所有满足要求的文件目录，返回到一个切片中
func GetAllFile(pathname string, sliceIn []string) ([]string, error) {
	rd, err := ioutil.ReadDir(pathname)
	
    for _, fi := range rd {
		if fi.IsDir() { // 判断是否为文件夹，如果是文件夹，再找里面的文件			
			sliceIn, _ = GetAllFile(pathname + "\\" + fi.Name(), sliceIn)			
		} else { // 是文件			
			suffixPosition := strings.LastIndex(fi.Name(), ".")
			suffix := string([]byte(fi.Name())[suffixPosition:])
			if suffix == ".flv" { // 找到目标文件，返回到一个切片中
				// fmt.Println(pathname + "\\" + fi.Name())
				sliceIn = append(sliceIn, pathname + "\\" + fi.Name()) 
			}
        }
    }
    return sliceIn, err
}

func main() {

	var pathname = `D:\Documents\Downloads\Bilibili_Download\62217176`
	var slice []string
	sliceOut, _ := GetAllFile(pathname, slice)

	for _, sli := range sliceOut {
		// fmt.Println(sli)
		// 提取最后的文件名
		filenamePosition := strings.LastIndex(sli, "\\")
		// filename := string([]byte(sli)[filenamePosition + 1:])
		// fmt.Println(filename)

		// 提取出名字中的编号
		// 读取info文件中的字符串，作为命名。62217176.info
		infoPath := string([]byte(sli)[:filenamePosition + 1])
		infoName := infoPath + "62217176.info"

		b, err := ioutil.ReadFile(infoName) // 读取文件
		if err != nil {
			fmt.Print(err)
		}
		str := string(b) // 转换为字符串输出

		startPosition := strings.LastIndex(str, `"PartName":"`)
		endPosition := strings.LastIndex(str, `","Format"`)
		title := string([]byte(str)[startPosition + 12 : endPosition])
		// fmt.Println(title)

		// 实现移动文件（同时实现文件的移动和重命名）
		err2 := os.Rename(sli, `D:\Videos\Golang_Study\` + title + `.flv`)
		if err2 != nil {
			fmt.Println(err2)
			return
		}

	}
	
}