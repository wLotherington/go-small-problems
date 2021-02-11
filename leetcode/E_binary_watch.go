// Easy
// https://leetcode.com/problems/binary-watch/

package main

import "fmt"

func readBinaryWatch(num int) []string {
	watch := [10]int{0}

	for i := 0; i < num; i++ {
		watch[i] = 1
	}

	possibleTimes := permute(watch, 0, 9)

	fmt.Println(possibleTimes)

	return []string{"hi"}
}

func permute(arr [10]int, sIdx int, eIdx int) [][10]int {
	if sIdx == eIdx {
		return
	}
	return [][10]int{arr}
}

func main() {
	fmt.Println(readBinaryWatch(1))
}
