package main

import "fmt"

func main() {
	prices := []int{7, 6, 4, 3, 1}

	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	low := prices[0]
	max := 0
	// high := prices[0]

	for _, v := range prices {
		if v-low < 0 {
			low = v
		} else if v-low > max {
			max = v - low
		}
	}

	return max
}
