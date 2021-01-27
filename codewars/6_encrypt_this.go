// kyu6
// https://www.codewars.com/kata/5848565e273af816fb000449

package kata

import (
  s "strings"
  sc "strconv"
)

func encrypt(w string) string {
  if len(w) == 0 { return w }
  
  strArr := s.Split(w, "")
  
  for _, v := range strArr[0] {
    strArr[0] = sc.Itoa(int(v))
  }
  
  if len(w) <= 1 { return s.Join(strArr, "") }
  
  strArr[1], strArr[len(strArr) - 1] = strArr[len(strArr) - 1], strArr[1]
  
  return s.Join(strArr, "")
}

func EncryptThis(text string) string {
  strArr := s.Split(text, " ")
  
  for i, word := range strArr {
    strArr[i] = encrypt(word)
  }
  
  return s.Join(strArr, " ")
}