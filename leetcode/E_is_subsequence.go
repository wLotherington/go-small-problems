// Easy
// https://leetcode.com/problems/is-subsequence/

package main

import "fmt"

func main() {
	s := "abcdfdw"
	t := "ahbgdcdfffdfffdddddfffdfffdzzzzw"

	fmt.Println(isSubsequence(s, t))
}

func isSubsequence(s string, t string) bool {
	sIdx := 0
	sLength := len(s)

	for i := 0; i < len(t); i++ {
		if sIdx == sLength {
			return true
		}

		if s[sIdx] == t[i] {
			sIdx++
		}
	}

	return sIdx == len(s)
}
