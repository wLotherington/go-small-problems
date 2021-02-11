// 7kyu
// https://www.codewars.com/kata/567501aec64b81e252000003/

package kata

import "math"

func WallPaper(l, w, h float64) string {
	if l*w*h == 0 {
		return "zero"
	}
	numbers := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen", "twenty"}
	area := (l*2 + w*2) * h * 1.15

	idx := int(math.Ceil(area / 5.2))
	return numbers[idx]
}
