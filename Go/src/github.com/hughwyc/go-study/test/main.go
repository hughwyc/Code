package main

import (
	"fmt"
	"time"
)

func main() {

	// 间隔时间打印输出
	var str string

	str = `正巧今年的情人节恰逢恋爱300天纪念。可惜我们各自被疫情困在家里，不能牵手度过。很幸运在校园里遇见你，认识你，爱上你。在异地的日子里有欢喜也有小吵闹，希望亲爱的小梅芝在余生多多包涵呀~Happy Valentine's Day~`

	str2 := []rune(str)

	for _, v := range str2 {
		fmt.Print(string(v))
		time.Sleep(time.Duration(250) * time.Millisecond)

	}

	//fmt.Println(str)

	//var temp byte
	//fmt.Println("Press any key to exit.")
	//_, _ = fmt.Scan(&temp)

	//var stone []int
	//var steel []int
	//
	//stone = append(stone, 9846)
	//steel = append(steel, 3916)
	//
	//stone = append(stone, 959)
	//steel = append(steel, 287)
	//
	//stone = append(stone, 16519)
	//steel = append(steel, 7194)
	//
	//stone = append(stone, 684)
	//steel = append(steel, 250)
	//
	//stone = append(stone, 12512)
	//steel = append(steel, 6373)
	//
	//stone = append(stone, 20105)
	//steel = append(steel, 1490)
	//
	//stone = append(stone, 702)
	//steel = append(steel, 257)
	//
	//stone = append(stone, 40776)
	//steel = append(steel, 2831)
	//
	//stone = append(stone, 21025)
	//steel = append(steel, 5503)
	//
	//stone = append(stone, 31602)
	//steel = append(steel, 15848)
	//
	//stone = append(stone, 618)
	//steel = append(steel, 551)
	//
	//stoneSum := 0
	//steelSum := 0
	//
	//for _, v := range stone {
	//	stoneSum += v
	//}
	//
	//for _, v := range steel {
	//	steelSum += v
	//}
	//
	//fmt.Printf("stoneSum = %.2f, steelSum = %.2f.", float64(stoneSum)*1.1, float64(steelSum)*1.1)

	//bear := make(map[string]string)
	//
	//bear["01"] = "brizz"
	//bear["02"] = "panda"
	//bear["03"] = "ice"
	//
	//fmt.Println(bear)
	//fmt.Println(len(bear))
	//
	//start := 2
	//end := 6
	//
	//mid := (end + start) >> 1
	//
	//fmt.Println("mid=", mid)

	//fmt.Println(time.Now().String())
	//
	//timeTemplate := "2006-01-02"
	//
	//endStamp, _ := time.ParseInLocation(timeTemplate, time.Now().String(), time.Local)
	//
	//endUnix := endStamp.Unix()
	//
	//fmt.Println(endUnix)
	//
	//fmt.Println(time.Now().Unix())

	//fmt.Println(strconv.Itoa(2019))

	//var arr = [...]int{1, 2, 3, 4, 5}
	//
	//slice := arr[1:4]
	//
	//for i, v := range slice {
	//	fmt.Printf("i=%v, v=%v \n", i, v)
	//}

}
