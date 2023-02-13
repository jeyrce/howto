package main

import (
	"fmt"
	"sort"
	"sync"
)

// 并发求每一行中位数并统计奇数数量
func main() {
	var oddCount int = 0
	var eveCount int = 0
	var odd = make(chan int, 1) // 如果是奇数则返回1,直接传计数指针也可以不用chan
	// go func(odd <-chan int) {
	// 	for {
	// 		select {
	// 		case num := <-odd:
	// 			if num == 0 {
	// 				eveCount++
	// 			} else {
	// 				oddCount++
	// 			}
	// 		default:
	// 			//
	// 		}
	// 	}
	// }(odd)
	arry := [][]int{
		{1, 3, 5, 7, 9},
		{0, 2, 4, 6, 8},
		{3, 5, 6, 0, 6},
	}
	var wg = sync.WaitGroup{}
	for _, item := range arry {
		wg.Add(1)
		go getMiddle(odd, &wg, item, &oddCount, &eveCount)
	}
	wg.Wait()
	fmt.Println("奇数", oddCount, "偶数", eveCount)
}

// 找出中位数，排序 + 取中间的索引
// 排序算法没自己写，偷懒了
func getMiddle(c chan int, wg *sync.WaitGroup, arry []int, oodCnt, eveCnt *int) int {
	defer wg.Done()
	sort.SliceStable(arry, func(i, j int) bool {
		return arry[i] <= arry[j]
	})
	// 直接取有序数组index比较暴力，还可以双指针，步长分别是1,2，当2步长走到末尾时1步长就在中间
	// var fast, slow = 0, 0
	// for {
	// 	fast+=2
	// 	slow+=1
	// 	if fast >= len(arry) - 1{
	// 		return arry[slow]
	// 	}
	// }
	index := len(arry) / 2
	target := arry[index]
	if target%2 != 0 {
		*oodCnt++
		// c <- 1
	} else {
		*eveCnt++
		// c <- 0
	}
	return target
}
