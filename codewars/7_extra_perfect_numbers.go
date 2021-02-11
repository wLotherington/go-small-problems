// 7kyu
// https://www.codewars.com/kata/5a662a02e626c54e87000123/

package kata

import "strconv"

func ExtraPerfect(n int) (result []int) {
	for i := 1; i <= n; i++ {
		cand := strconv.FormatInt(int64(i), 2)
		if cand[0] == cand[len(cand)-1] {
			result = append(result, i)
		}
	}

	return
}
