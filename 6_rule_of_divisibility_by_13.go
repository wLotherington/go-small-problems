// kyu6
// https://www.codewars.com/kata/564057bc348c7200bd0000ff

package kata

import sc "strconv"

func Thirt(n int) int {
  var seq []int = []int{1, 10, 9, 12, 3, 4, 1, 10, 9, 12, 3, 4, 1, 10, 9, 12, 3, 4, 1, 10, 9, 12, 3, 4, 1, 10, 9, 12, 3, 4, 1, 10, 9, 12, 3, 4, 1, 10, 9, 12, 3, 4, 1, 10, 9, 12, 3}
  var digits []string
  
  for _, d := range sc.Itoa(n) {
    digits = append(digits, string(d))
  }
  
  total := 0
  idx := 0
  for i := len(digits) - 1; i >= 0; i-- {
    temp, _ := sc.Atoi(digits[i])
    total += temp * seq[idx]
    idx++
  }
  
  if total == n {
    return total
  } else {
    return Thirt(total)
  }
}