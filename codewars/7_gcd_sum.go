// 7kyu
// https://www.codewars.com/kata/5dd259444228280032b1ed2a/

package kata

func Solve(s int, g int) []int {
	cand := s - g
	result := []int{-1, -1}

	if cand%g == 0 {
		result = []int{g, cand}
	}

	return result
}
