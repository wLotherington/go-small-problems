// 7kyu
// https://www.codewars.com/kata/5bd776533a7e2720c40000e5/

package kata

import "sort"

func Pendulum(values []int) []int {
	sort.Ints(values)
	result := []int{}

	for i, n := range values {
		if i%2 == 0 {
			result = append([]int{n}, result...)
		} else {
			result = append(result, n)
		}
	}

	return result
}
