package main

import (
	"fmt"
	"strconv"
	"time"
)

// 如果从1990年1月1日开始执行“三天打鱼，两天晒网”，
// 如何判断在以后的某一天中是“打鱼”还是“晒网”？

func dateDifference(start, end string) int64 {

	timeTemplate := "2006-1-2"

	// 转换成时间戳
	startStamp, _ := time.ParseInLocation(timeTemplate, start, time.Local)
	endStamp, _ := time.ParseInLocation(timeTemplate, end, time.Local)

	startUnix := startStamp.Unix()
	endUnix := endStamp.Unix()

	date := (endUnix - startUnix) / 86400

	return date
}

func main() {

	var year int
	var month int
	var day int

	fmt.Println("请输入年月日，以-隔开(如：2006-1-2)：")
	_, err := fmt.Scanf("%d-%d-%d", &year, &month, &day)

	if err != nil {
		fmt.Println("Input wrong ! : ",err)
	}

	endTime := strconv.Itoa(year) + "-" + strconv.Itoa(month) + "-" + strconv.Itoa(day)

	days := dateDifference("1900-1-1", endTime)
	fmt.Println("距1900-1-1经过", days, "天")

	if days % 5 < 3 {
		fmt.Println("这一天在打鱼")
	} else {
		fmt.Println("这一天在晒网")
	}
}