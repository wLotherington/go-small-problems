// 7kyu
// https://www.codewars.com/kata/5a092d9e46d843b9db000064/

package kata

func Solve(arr []int) int {
	nums := make(map[int]int)

	for _, n := range arr {
		if n < 0 {
			nums[-n]--
		} else {
			nums[n]++
		}
	}

	for k, v := range nums {
		if v > 0 {
			return k
		} else if v < 0 {
			return -k
		}
	}

	return 0
}
