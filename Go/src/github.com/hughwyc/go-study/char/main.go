package main 

import "fmt" 

func main() {
    // 字符类型，注意字符常量是使用单引号的 ''
    var c1 byte = 'a'

    // 直接输出byte值时，输出的是对应字符的码值。
    fmt.Println("c1 =", c1)

    // 如果需要对应字符，需要使用格式化输出
    fmt.Printf("c1 = %c \n", c1)

    // var c2 byte = '北' // “北”字符对应的 UTF-8 码值为21271，overflow    
    var c2 int = '北'
    // 英文字母 1 个字节，汉字 3 个字节

    fmt.Printf("c2 = %c\nc2对应的码值 = %d\n", c2, c2)

    fmt.Printf("c3 = %c\n", 22269) // "国"对应的 UTF-8 的编码值

    // 字符类型是可以进行运算的，相当于一个整数，应为它都对应有Unicode编码

    var c4 int = '你' // 注意，一定要用单引号，如果是双引号，则代表字符串，而不是单个字符了
    var c5 int = '好'

    fmt.Println("c4 + c5 =", c4 + c5)
    fmt.Println("使用Println可以自动换行") // Println函数在输出后会自动换行，因为表面意思就是“输出一行”

    // 使用for循环输出a-Z
    for i := 0; i < 26; i++ {
        fmt.Printf("%c",'a' + i)
    }

    fmt.Printf("\n")

    for i := 0; i < 26; i++ {
        fmt.Printf("%c",'A' + i)
    }

}