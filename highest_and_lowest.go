// 7kyu
// https://www.codewars.com/kata/554b4ac871d6813a03000035/

package kata

import (
  s "strings"
  sc "strconv"
)

func HighAndLow(numStr string) string {
  numArr := s.Split(numStr, " ")
  var max, min int
  
  for i, n := range numArr {
    tmp, _ := sc.Atoi(n)
    
    if i == 0 {
      min = tmp
      max = tmp
    }
    
    if tmp > max {
      max = tmp
    }
    
    if tmp < min {
      min = tmp
    }
  }
  
  return sc.Itoa(max) + " " + sc.Itoa(min)
}