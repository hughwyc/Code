package main // 把这个hello.go归属到main包中

import "fmt" // 引入一个包 fmt

// Ctrl + /  多行注释 1234567890
// /*     */  块注释

var a = 10
var b = "10"

var c = 1

func main() {
    // 调用包fmt里的Println函数，输出Hello world
    fmt.Println("Hello world")
    fmt.Println("Hello")
    fmt.Println(a) 
}