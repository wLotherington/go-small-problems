// 6kyu
// https://www.codewars.com/kata/56a5d994ac971f1ac500003e/

package kata

import "fmt"

func LongestConsec(strArr []string, k int) string {
  fmt.Println(strArr, k)
  subStr := ""
  longest := ""
  
  for i := 0; i < len(strArr) - k + 1; i++ {
    subStr = ""
    
    for j := 0; j < k; j++ {
      subStr += strArr[i + j]
    }
    
    if len(subStr) > len(longest) {
      longest = subStr
    }
  }
  
  return longest
}