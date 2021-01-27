// 6kyu
// https://www.codewars.com/kata/5b162ed4c8c47ea2f5000023/

package main 

import (
	"fmt"
	sc "strconv"
)

func main() {
	fmt.Println(Solve(1) == 1)
	fmt.Println(Solve(2) == 2)
	fmt.Println(Solve(18) == 18)
	fmt.Println(Solve(48) == 48)
	fmt.Println(Solve(100) == 99)
	fmt.Println(Solve(10) == 9)
	fmt.Println(Solve(110) == 99)
	fmt.Println(Solve(2090) == 1999)
}

func Solve(n int) int {
	var candidate int
	maxVal := n
	maxSum := digitSum(n)

	for i := (len(sc.Itoa(n)) - 1); i >= 0 ; i--{
		candidate = makeCandidate(n, i)

		if candidate < n && digitSum(candidate) > maxSum {
			maxVal = candidate
			maxSum = digitSum(candidate)
		}
	}

	return maxVal
}

func digitSum(n int) int {
	sum := 0
	nStr := sc.Itoa(n)

	for _, char := range nStr {
		tmp, _ := sc.Atoi(string(char))
		sum += tmp
	}

	return sum
}

func makeCandidate(n, idx int) int {
	nStr := sc.Itoa(n)
	candidateStr := ""

	for i, char := range nStr {
		if i == 0 {
			tmp, _ := sc.Atoi(string(char))

			if i == idx {
				tmp--
			}

			candidateStr += sc.Itoa(tmp)
		} else if idx == i {
			candidateStr += "8"
		} else {
			candidateStr += "9"
		}
	}

	candidate, _ := sc.Atoi(candidateStr)

	return candidate
}
