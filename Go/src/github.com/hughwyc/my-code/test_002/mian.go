package main

import (
    "fmt"
	"io/ioutil"
	"strings"
)

func main() {
    b, err := ioutil.ReadFile(`D:\Documents\Downloads\Bilibili_Download\62217176\12\62217176.info`)
    if err != nil {
        fmt.Print(err)
    }
    fmt.Println(b)
    str := string(b)
	fmt.Println(str)
	
	// 找到 "PartName":" 的位置
	startPosition := strings.LastIndex(str, `"PartName":"`)
	endPosition := strings.LastIndex(str, `","Format"`)
	filename := string([]byte(str)[startPosition + 12 : endPosition])
	fmt.Println(filename)
}
