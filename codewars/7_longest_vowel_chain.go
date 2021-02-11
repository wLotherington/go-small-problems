// 7kyu
// https://www.codewars.com/kata/59c5f4e9d751df43cf000035/

package kata

func Solve(s string) int {
	sub := ""
	max := 0

	for _, c := range s {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			sub += string(c)
		} else if max < len(sub) {
			max = len(sub)
			sub = ""
		} else {
			sub = ""
		}
	}

	return max
}
