package main

import (
	"fmt"
)

func leapYear(year int32) bool {
	if (year % 4 == 0 && year %100 != 0) || year % 400 == 0 {
		return true
	}else{
		return false
	}
}

func main() {

	var year int32
	var month int32
	fmt.Println("请输入年份：")
	fmt.Scanln(&year)
	for {
		if year < 0 {
			fmt.Println("输入的年份不正确，请重新输入")
			_, _ = fmt.Scanln(&year)
		} else {
			break
		}
	}

	fmt.Println("请输入月份：")
	fmt.Scanln(&month)
	for {
		if month < 1 || month > 12 {
			fmt.Println("输入的月份不正确，请重新输入")
			fmt.Scanln(&month)
		} else {
			break
		}
	}

	//fmt.Println("year=", year)
	//fmt.Println("month=", month)

	// 判断是否为闰年
	if month == 1 || month == 3 || month == 5 ||
		month == 7 || month == 8 || month == 10 || month == 12 {
		fmt.Println("这个月份有31天")
	} else if month == 4 || month == 6 || month == 9 || month == 11 {
		fmt.Println("这个月份有30天")
	} else {
		if leapYear(year) {
			fmt.Println("这个月份有29天")
		} else {
			fmt.Println("这个月份有28天")
		}
	}

}