package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRangeNum(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max-min) + min
	return randNum
}

func main() {

	random := GenerateRangeNum(0, 100) // return a random integer for (0,100)
	fmt.Println("random=:", random)

	var num int

	fmt.Println("please guess a number in range of 0-100:")

	for i := 1; i < 11; i++ {

		for {
			//stdin := bufio.NewReader(os.Stdin) // 关于go fmt.Scan Scanf Scanln的一个小问题/https://blog.csdn.net/xychidy520_java/article/details/81948432
			_, err := fmt.Scanln(&num)
			//_, _ = stdin.ReadString('\n')

			if err != nil {
				fmt.Printf("input wrong: %v, try again.\n", err)
			} else {
				break
			}
		}

		if num == random {

			//使用switch
			switch i {
			case 1:
				fmt.Println("you are a genius.")
			case 2, 3:
				fmt.Println("you are pretty smart.")
			case 4, 5, 6, 7, 8, 9:
				fmt.Println("just fine.")
			case 10:
				fmt.Println("finally, you get it.")
			}
			break
		} else if i == 10 && num != random {
			fmt.Println("10 times out, what can I say to you ?")
		} else {
			fmt.Println("oops, try again:")
		}

		//fmt.Println("i=", i)
	}
}
