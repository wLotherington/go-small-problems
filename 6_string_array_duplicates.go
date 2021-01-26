// 6kyu
// https://www.codewars.com/kata/59f08f89a5e129c543000069/

package kata

func Dup(arr []string) []string {
  tmp := ""
  
  for i, str := range arr {
    tmp = ""
    
    for _, char := range str {
      if len(tmp) == 0 || string(tmp[len(tmp) - 1]) != string(char) {
        tmp += string(char)
      }
    }
    
    arr[i] = tmp
  }
  
  return arr
}