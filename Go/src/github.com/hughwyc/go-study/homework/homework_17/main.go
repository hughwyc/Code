package main

import (
	"fmt"
	"math"
)

// 跳水比赛，8个评委打分。运动员的成绩是，去掉一个最高分，去掉一个最低分，剩余6个分数的平均分。
//使用一维数组实现：
//（1）请把打最高分和最低分的评委找出来
//（2）找出最佳评委（打分与最后得分最接近的评委）和最差评委（打分与最后得分相差最大的评委）

func main() {
	arr := [8]int{4, 2, 3, 1, 5, 6, 8, 7}
	fmt.Println("the scores are: ", arr)

	max := arr[0]
	min := arr[0]
	maxIndex := 0
	minIndex := 0

	for i, v := range arr {
		if v > max {
			max = v
			maxIndex = i
		}
		if v < min {
			min = v
			minIndex = i
		}
	}

	fmt.Printf("the maxIndex = %v, the minIndex = %v \n", maxIndex, minIndex)

	// 去掉最高分和最低分，求剩余平均分
	sum := 0
	for i, v := range arr {
		if i == maxIndex || i == minIndex {
			continue
		}
		sum += v
	}

	average := float64(sum) / float64(len(arr)-2)
	fmt.Println("最终得分为：", average)

	minDiff := math.Abs(float64(arr[0]) - average)
	maxDiff := minDiff
	for _, v := range arr {
		diff := math.Abs(float64(v) - average)
		if minDiff > diff {
			minDiff = diff
		}
		if maxDiff < diff {
			maxDiff = diff
		}
	}

	fmt.Printf("maxDiff = %v, minDiff = %v.\n", maxDiff, minDiff)
	var bestJudgeIndex []int
	var worstJudgeIndex []int

	for i, v := range arr {
		diff := math.Abs(float64(v) - average)
		if diff == minDiff {
			bestJudgeIndex = append(bestJudgeIndex, i)
		} else if diff == maxDiff {
			worstJudgeIndex = append(worstJudgeIndex, i)
		}
	}

	fmt.Printf("the worst judge's Index = %v, with maxDiff = %v.\n", worstJudgeIndex, maxDiff)
	fmt.Printf("the best judge's Index = %v, with minDiff = %v.\n", bestJudgeIndex, minDiff)
}
