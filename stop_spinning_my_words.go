// 6kyu
// https://www.codewars.com/kata/5264d2b162488dc400000001/

package kata

import s "strings"

func SpinWords(str string) string {
  arr := s.Split(str, " ")
  
  for i, w := range arr {
    if (len(w) >= 5) {
      tmp := ""
      
      for j := len(w) - 1; j >= 0; j-- {
        tmp += string(w[j])
      }
      
      arr[i] = tmp
    }
  }

  return s.Join(arr, " ")
}