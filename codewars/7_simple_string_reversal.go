// 7kyu
// https://www.codewars.com/kata/5a71939d373c2e634200008e/

package kata

import "strings"

func Solve(s string) string {
  stripped := strings.Join(strings.Split(s, " "), "")
  result := ""
  idx := len(stripped) - 1

  for _, c := range s {
    if c == ' ' {
      result += " "
    } else {
      result += string(stripped[idx])
      idx--
    }
  }

  return result
}
