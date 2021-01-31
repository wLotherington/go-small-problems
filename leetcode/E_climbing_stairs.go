package main

import "fmt"

func main() {
	fmt.Println(climbStairs(10))
}

func climbStairs(n int) int {
	return climbStairsWithMemo(n, map[int]int{})
}

func climbStairsWithMemo(n int, memo map[int]int) int {
	if n <= 3 {
		return n
	}

	if memo[n] == 0 {
		memo[n] = climbStairsWithMemo(n-1, memo) + climbStairsWithMemo(n-2, memo)
	}

	return memo[n]
}
